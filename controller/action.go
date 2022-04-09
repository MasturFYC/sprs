package controller

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"fyc.com/sprs/models"
	"github.com/gin-gonic/gin"
)

func ActionUploadFile(c *gin.Context) {
	b := make([]byte, 4)
	rand.Read(b)
	uid := hex.EncodeToString(b)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.Request.ParseMultipartForm(5 << 20) // maxMemory 5MB
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Access the photo key - First Approach
	file, h, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ext := path.Ext(h.Filename)
	sfile := fmt.Sprintf("%s%s", uid, ext)
	targetPath := path.Join(os.Getenv("UPLOADFILE_LOCATION"), sfile)

	tmpfile, err := os.Create(targetPath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer tmpfile.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = io.Copy(tmpfile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	rowsAffected, err := update_file_name(db, &id, &sfile)

	if err != nil {
		c.JSON(http.StatusRequestURITooLong, gin.H{"error": err.Error()})
		return
	}
	res := Response{
		ID:      rowsAffected,
		Message: sfile,
	}

	c.JSON(http.StatusOK, res)
}

func ActionGetByOrder(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	actions, err := action_getByOrder(db, &id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &actions)
}

func ActionGetItem(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	act, err := getAction(db, &id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &act)
}

func ActionDelete(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	deletedRows, err := deleteAction(db, &id)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Action deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)

}

func ActionCreate(c *gin.Context) {

	var act models.Action

	err := c.BindJSON(&act)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	id, err := createAction(db, &act)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	act.ID = id

	c.JSON(http.StatusCreated, &act)
}

func ActionUpdate(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var act models.Action

	err := c.BindJSON(&act)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateAction(db, &id, &act)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Action updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)

}

func action_getByOrder(db *sql.DB, OrderID *int64) ([]models.Action, error) {
	// defer db.Close()
	var actions = make([]models.Action, 0)

	sqlStatement := `SELECT
    id, action_at, pic, descriptions, order_id, file_name
  FROM actions
  WHERE order_id=$1`

	rows, err := db.Query(sqlStatement, OrderID)

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

func getAction(db *sql.DB, id *int64) (models.Action, error) {
	var act models.Action

	sqlStatement := `SELECT
    id, action_at, pic, descriptions, order_id, file_name
  FROM actions
  WHERE id=$1`
	//stmt, _ := db.Prepare(sqlStatement)

	//defer stmt.Close()
	row := db.QueryRow(sqlStatement, id)

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

func deleteAction(db *sql.DB, id *int64) (int64, error) {
	// create the delete sql query
	sqlStatement := `DELETE FROM actions WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete action. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func createAction(db *sql.DB, act *models.Action) (int64, error) {

	sqlStatement := `INSERT INTO actions
    (action_at, pic, descriptions, order_id)
  VALUES
    ($1, $2, $3, $4)
  RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement,
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

func updateAction(db *sql.DB, id *int64, act *models.Action) (int64, error) {

	sqlStatement := `UPDATE actions SET
    action_at=$2, pic=$3, descriptions=$4, order_id=$5
  WHERE id=$1`

	res, err := db.Exec(sqlStatement,
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

	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func update_file_name(db *sql.DB, id *int64, file_name *string) (int64, error) {

	sqlStatement := `UPDATE actions SET file_name=$2 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, file_name)

	if err != nil {
		log.Printf("Unable to update action. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

/* referrences
/// https://github.com/axetroy/gin-uploader/blob/master/handler.go?msclkid=727d479fb62f11ecb05502e0899ec2f3
*/
func ActionGetFile(c *gin.Context) {
	txt := c.Param("txt")

	targetPath := path.Join(os.Getenv("UPLOADFILE_LOCATION"), txt)

	if !exists(targetPath) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	http.ServeFile(c.Writer, c.Request, targetPath)
}

func ActionGetPreview(c *gin.Context) {

	txt := c.Param("txt")

	targetPath := path.Join(os.Getenv("UPLOADFILE_LOCATION"), txt)

	if strings.ToLower(path.Ext(targetPath)) == ".pdf" {
		targetPath = path.Join(os.Getenv("UPLOADFILE_LOCATION"), "default.jpg")
	}

	if !exists(targetPath) {

		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return

	}
	//			fileBytes, err := ioutil.ReadFile(targetPath)
	f, err := os.Open(targetPath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	//encoding message is discarded, because OP wanted only jpg, else use encoding in resize function
	img, _, err := image.Decode(f)
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, gin.H{"error": err.Error()})
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

	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Data(http.StatusOK, "application/jpg", imgBytes)

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
