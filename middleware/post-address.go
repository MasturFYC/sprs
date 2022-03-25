package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

// func GetPostAddresses(w http.ResponseWriter, r *http.Request) {
// 	EnableCors(&w)

// 	addresses, err := getAllPostAddresses()

// 	if err != nil {
// 		log.Fatalf("Unable to get all post addresses. %v", err)
// 	}

// 	json.NewEncoder(w).Encode(&addresses)
// }

func GetPostAddress(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ha, err := getPostAddress(&id)

	if err != nil {
		//log.Fatalf("Unable to get post address. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&ha)
}

func DeletePostAddress(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows := deletePostAddress(&id)

	msg := fmt.Sprintf("Post address deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      id,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreatePostAddress(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var ha models.PostAddress

	err := json.NewDecoder(r.Body).Decode(&ha)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rowAffected, err := createPostAddress(&ha)

	if err != nil {
		//log.Fatalf("Nama post address tidak boleh sama.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Post address created successfully. Total rows/record affected %v", rowAffected)

	// format the reponse message
	res := Response{
		ID:      rowAffected,
		Message: msg,
	}

	json.NewEncoder(w).Encode(&res)

}

func UpdatePostAddress(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var ha models.PostAddress

	err := json.NewDecoder(r.Body).Decode(&ha)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	updatedRows := updatePostAddress(&id, &ha)

	msg := fmt.Sprintf("Post address updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getPostAddress(id *int64) (models.PostAddress, error) {

	var ha models.PostAddress

	var sqlStatement = `SELECT 
		order_id, street, region, city, phone, zip
	FROM post_addresses
	WHERE order_id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&ha.OrderID, &ha.Street, &ha.Region, &ha.City, &ha.Phone, &ha.Zip)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return ha, nil
	case nil:
		return ha, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return ha, err
}

// func getAllPostAddresses() ([]models.PostAddress, error) {

// 	var addresses []models.PostAddress

// 	var sqlStatement = `SELECT
// 		order_id, street, region, city, phone, zip
// 	FROM post_addresses
// 	ORDER BY name`

// 	rs, err := Sql().Query(sqlStatement)

// 	if err != nil {
// 		log.Fatalf("Unable to execute post addresses query %v", err)
// 	}

// 	defer rs.Close()

// 	for rs.Next() {
// 		var ha models.PostAddress

// 		err := rs.Scan(&ha.OrderID, &ha.Street, &ha.Region, &ha.City, &ha.Phone, &ha.Zip)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		addresses = append(addresses, ha)
// 	}

// 	return addresses, err
// }

func deletePostAddress(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM post_addresses WHERE order_id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete post address. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createPostAddress(ha *models.PostAddress) (int64, error) {

	sqlStatement := `INSERT INTO post_addresses
	(order_id, street, region, city, phone, zip) 
	VALUES 
	($1, $2, $3, $4, $5, $6)`

	res, err := Sql().Exec(sqlStatement,
		ha.OrderID,
		ha.Street,
		ha.Region,
		ha.City,
		ha.Phone,
		ha.Zip,
	)

	if err != nil {
		log.Fatalf("Unable to create post address. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Unable to create post address. %v", err)
	}

	return rowsAffected, err
}

func updatePostAddress(id *int64, ha *models.PostAddress) int64 {

	sqlStatement := `UPDATE post_addresses SET
		street=$2, region=$3, city=$4, phone=$5, zip=$6
	WHERE order_id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		ha.Street,
		ha.Region,
		ha.City,
		ha.Phone,
		ha.Zip,
	)

	if err != nil {
		log.Fatalf("Unable to update post address. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating post address. %v", err)
	}

	return rowsAffected
}
