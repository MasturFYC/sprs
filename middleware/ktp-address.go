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

// func GetKTPAddresses(w http.ResponseWriter, r *http.Request) {
// 	EnableCors(&w)

// 	addresses, err := getAllKTPAddresses()

// 	if err != nil {
// 		log.Fatalf("Unable to get all ktp addresses. %v", err)
// 	}

// 	json.NewEncoder(w).Encode(&addresses)
// }

func GetKTPAddress(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	ha, err := getKTPAddress(&id)

	if err != nil {
		log.Fatalf("Unable to get ktp address. %v", err)
	}

	json.NewEncoder(w).Encode(&ha)
}

func DeleteKTPAddress(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteKTPAddress(&id)

	msg := fmt.Sprintf("KTP address deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      id,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreateKTPAddress(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var ha models.KtpAddress

	err := json.NewDecoder(r.Body).Decode(&ha)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	rowAffected, err := createKTPAddress(&ha)

	if err != nil {
		log.Fatalf("Nama ktp address tidak boleh sama.  %v", err)
	}

	msg := fmt.Sprintf("KTP address created successfully. Total rows/record affected %v", rowAffected)

	// format the reponse message
	res := Response{
		ID:      rowAffected,
		Message: msg,
	}

	json.NewEncoder(w).Encode(&res)

}

func UpdateKTPAddress(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var ha models.KtpAddress

	err := json.NewDecoder(r.Body).Decode(&ha)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateKTPAddress(&id, &ha)

	msg := fmt.Sprintf("KTP address updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getKTPAddress(id *int64) (models.KtpAddress, error) {

	var ha models.KtpAddress

	var sqlStatement = `SELECT 
		order_id, street, region, city, phone, zip
	FROM ktp_addresses
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

// func getAllKTPAddresses() ([]models.KTPAddress, error) {

// 	var addresses []models.KTPAddress

// 	var sqlStatement = `SELECT
// 		order_id, street, region, city, phone, zip
// 	FROM ktp_addresses
// 	ORDER BY name`

// 	rs, err := Sql().Query(sqlStatement)

// 	if err != nil {
// 		log.Fatalf("Unable to execute ktp addresses query %v", err)
// 	}

// 	defer rs.Close()

// 	for rs.Next() {
// 		var ha models.KTPAddress

// 		err := rs.Scan(&ha.OrderID, &ha.Street, &ha.Region, &ha.City, &ha.Phone, &ha.Zip)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		addresses = append(addresses, ha)
// 	}

// 	return addresses, err
// }

func deleteKTPAddress(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM ktp_addresses WHERE order_id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete ktp address. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createKTPAddress(ha *models.KtpAddress) (int64, error) {

	sqlStatement := `INSERT INTO ktp_addresses
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
		log.Fatalf("Unable to create ktp address. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Unable to create ktp address. %v", err)
	}

	return rowsAffected, err
}

func updateKTPAddress(id *int64, ha *models.KtpAddress) int64 {

	sqlStatement := `UPDATE ktp_addresses SET
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
		log.Fatalf("Unable to update ktp address. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating ktp address. %v", err)
	}

	return rowsAffected
}
