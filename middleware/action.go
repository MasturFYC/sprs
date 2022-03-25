package middleware

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"fyc.com/sprs/models"
	"github.com/gorilla/mux"
)

func Action_UploadFile(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	b := make([]byte, 4)
	rand.Read(b) // Doesn’t actually fail
	uid := hex.EncodeToString(b)

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = r.ParseMultipartForm(5 << 20) // maxMemory 32MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Access the photo key - First Approach
	file, h, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ext := filepath.Ext(h.Filename)
	sfile := fmt.Sprintf("%s%s", uid, ext)
	targetPath := filepath.Join(os.Getenv("UPLOADFILE_LOCATION"), sfile)

	tmpfile, err := os.Create(targetPath)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer tmpfile.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rowsAffected, err := update_file_name(&id, &sfile)

	if err != nil {
		w.WriteHeader(http.StatusRequestURITooLong)
		return
	}

	//msg := fmt.Sprintf("File successfully uploaded. Total rows/record affected %v", rowsAffected)

	// format the response message
	res := Response{
		ID:      rowsAffected,
		Message: sfile,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetActions(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	actions, err := getAllActions(&id)

	if err != nil {
		log.Printf("Unable to get all actions. %v", err)
	}

	json.NewEncoder(w).Encode(&actions)
}

func GetAction(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	act, err := getAction(&id)

	if err != nil {
		log.Fatalf("Unable to get category. %v", err)
	}

	json.NewEncoder(w).Encode(&act)
}

func DeleteAction(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteAction(&id)

	msg := fmt.Sprintf("Action deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

func CreateAction(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var act models.Action

	err := json.NewDecoder(r.Body).Decode(&act)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := createAction(&act)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	act.ID = id

	json.NewEncoder(w).Encode(&act)
}

func UpdateAction(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var act models.Action

	err := json.NewDecoder(r.Body).Decode(&act)

	//log.Print(act)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := updateAction(&id, &act)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Action updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

func getAllActions(OrderID *int64) ([]models.Action, error) {
	// defer Sql().Close()
	var actions []models.Action

	sqlStatement := `SELECT
    id, action_at, pic, descriptions, order_id, file_name
  FROM actions
  WHERE order_id=$1`

	rows, err := Sql().Query(sqlStatement, OrderID)

	if err != nil {
		log.Fatalf("Unable to execute actions query %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var act models.Action
		err := rows.Scan(
			&act.ID,
			&act.ActionAt,
			//	&act.Code,
			&act.Pic,
			&act.Descriptions,
			&act.OrderId,
			&act.FileName,
		)

		if err != nil {
			log.Fatalf("ACTION unable to scan the row. %v", err)
		}

		actions = append(actions, act)
	}

	return actions, err
}

func getAction(id *int64) (models.Action, error) {
	var act models.Action

	sqlStatement := `SELECT
    id, action_at, pic, descriptions, order_id, file_name
  FROM actions
  WHERE id=$1`
	//stmt, _ := Sql().Prepare(sqlStatement)

	//defer stmt.Close()
	row := Sql().QueryRow(sqlStatement, id)

	err := row.Scan(
		&act.ID,
		&act.ActionAt,
		//	&act.Code,
		&act.Pic,
		&act.Descriptions,
		&act.OrderId,
		&act.FileName,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return act, nil
	case nil:
		return act, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return act, err
}

// func getProductsByCategory(id *int) ([]models.Product, error) {
// 	// defer Sql().Close()

// 	var products []models.Product

// 	sqlStatement := `SELECT
// 		p.id, p.name, p.spec, p.base_unit,
// 		p.base_weight, p.base_price, p.first_stock,
// 		p.stock, p.is_active, p.is_sale, p.category_id
// 	FROM products AS p
// 	WHERE p.category_id=$1
// 	ORDER BY p.name`

// 	rows, err := Sql().Query(sqlStatement, id)

// 	if err != nil {
// 		log.Fatalf("Unable to execute product query %v", err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {

// 		var product models.Product

// 		err := rows.Scan(
// 			&product.ID,
// 			&product.Name,
// 			&product.Spec,
// 			&product.BaseUnit,
// 			&product.BaseWeight,
// 			&product.BasePrice,
// 			&product.FirstStock,
// 			&product.Stock,
// 			&product.IsActive,
// 			&product.IsSale,
// 			&product.CategoryID,
// 		)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		products = append(products, product)
// 	}

// 	return products, err
// }

func deleteAction(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM actions WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete action. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createAction(act *models.Action) (int64, error) {

	sqlStatement := `INSERT INTO actions
    (action_at, pic, descriptions, order_id)
  VALUES
    ($1, $2, $3, $4)
  RETURNING id`

	var id int64

	err := Sql().QueryRow(sqlStatement,
		act.ActionAt,
		//	act.Code,
		act.Pic,
		act.Descriptions,
		act.OrderId).Scan(&id)

	if err != nil {
		log.Printf("Unable to create action. %v", err)
		return 0, err
	}

	return id, err
}

func updateAction(id *int64, act *models.Action) (int64, error) {

	sqlStatement := `UPDATE actions SET
    action_at=$2, pic=$3, descriptions=$4, order_id=$5
  WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		act.ActionAt,
		// act.Code,
		act.Pic,
		act.Descriptions,
		act.OrderId,
	)

	if err != nil {
		log.Printf("Unable to update action. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Printf("Error while updating action. %v", err)
	// }

	return rowsAffected, err
}

func update_file_name(id *int64, file_name *string) (int64, error) {

	sqlStatement := `UPDATE actions SET file_name=$2 WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id, file_name)

	if err != nil {
		log.Printf("Unable to update action. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Printf("Error while updating action. %v", err)
	// }

	return rowsAffected, err
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func Action_GetFile(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	txt := params["txt"]

	targetPath := filepath.Join(os.Getenv("UPLOADFILE_LOCATION"), txt)

	if exists(targetPath) {
		fileBytes, err := ioutil.ReadFile(targetPath)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(fileBytes)

		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func Action_GetPreview(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	txt := params["txt"]

	targetPath := filepath.Join(os.Getenv("UPLOADFILE_LOCATION"), txt)
	log.Printf("%s", strings.ToLower(filepath.Ext(targetPath)))

	if strings.ToLower(filepath.Ext(targetPath)) == ".pdf" {
		targetPath = filepath.Join(os.Getenv("UPLOADFILE_LOCATION"), "default.jpg")
	}

	if exists(targetPath) {

		//			fileBytes, err := ioutil.ReadFile(targetPath)
		f, err := os.Open(targetPath)
		if err != nil {
			//log.Fatal(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		//encoding message is discarded, because OP wanted only jpg, else use encoding in resize function
		img, _, err := image.Decode(f)
		if err != nil {
			w.WriteHeader(http.StatusInsufficientStorage)
			return
			//				log.Fatal(err)
		}

		//this is the resized image
		resImg := resize(img, 64, 80)

		//this is the resized image []bytes
		imgBytes := imgToBytes(resImg)

		// if err != nil {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	return
		// }

		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(imgBytes)

		return
	}

	w.WriteHeader(http.StatusNotFound)

}

func resize(img image.Image, length int, width int) image.Image {
	//truncate pixel size
	minX := img.Bounds().Min.X
	minY := img.Bounds().Min.Y
	maxX := img.Bounds().Max.X
	maxY := img.Bounds().Max.Y
	for (maxX-minX)%length != 0 {
		maxX--
	}
	for (maxY-minY)%width != 0 {
		maxY--
	}
	scaleX := (maxX - minX) / length
	scaleY := (maxY - minY) / width

	imgRect := image.Rect(0, 0, length, width)
	resImg := image.NewRGBA(imgRect)
	draw.Draw(resImg,
		resImg.Bounds(),
		&image.Uniform{C: color.White},
		image.Point{}, draw.Over)
	for y := 0; y < width; y += 1 {
		for x := 0; x < length; x += 1 {
			averageColor := getAverageColor(img, minX+x*scaleX, minX+(x+1)*scaleX, minY+y*scaleY, minY+(y+1)*scaleY)
			resImg.Set(x, y, averageColor)
		}
	}
	return resImg
}

func getAverageColor(img image.Image, minX int, maxX int, minY int, maxY int) color.Color {
	var averageRed float64
	var averageGreen float64
	var averageBlue float64
	var averageAlpha float64
	scale := 1.0 / float64((maxX-minX)*(maxY-minY))

	for i := minX; i < maxX; i++ {
		for k := minY; k < maxY; k++ {
			r, g, b, a := img.At(i, k).RGBA()
			averageRed += float64(r) * scale
			averageGreen += float64(g) * scale
			averageBlue += float64(b) * scale
			averageAlpha += float64(a) * scale
		}
	}

	averageRed = math.Sqrt(averageRed)
	averageGreen = math.Sqrt(averageGreen)
	averageBlue = math.Sqrt(averageBlue)
	averageAlpha = math.Sqrt(averageAlpha)

	averageColor := color.RGBA{
		R: uint8(averageRed),
		G: uint8(averageGreen),
		B: uint8(averageBlue),
		A: uint8(averageAlpha)}

	return averageColor
}

func imgToBytes(img image.Image) []byte {
	var opt jpeg.Options
	opt.Quality = 80

	buff := bytes.NewBuffer(nil)
	err := jpeg.Encode(buff, img, &opt)
	if err != nil {
		log.Fatal(err)
	}

	return buff.Bytes()
}
