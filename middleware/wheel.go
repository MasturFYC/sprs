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

func GetWheels(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	wheels, err := getAllWheels()

	if err != nil {
		log.Fatalf("Unable to get all wheels. %v", err)
	}

	json.NewEncoder(w).Encode(&wheels)
}

func GetWheel(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	wheels, err := getWheel(&id)

	if err != nil {
		log.Fatalf("Unable to get wheel. %v", err)
	}

	json.NewEncoder(w).Encode(&wheels)
}

func DeleteWheel(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteWheel(&id)

	msg := fmt.Sprintf("Wheel deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreateWheel(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var wheel models.Wheel

	err := json.NewDecoder(r.Body).Decode(&wheel)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	id, err := createWheel(&wheel)

	if err != nil {
		log.Fatalf("Nama wheel tidak boleh sama.  %v", err)
	}

	wheel.ID = id

	json.NewEncoder(w).Encode(&wheel)

}

func UpdateWheel(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var wheel models.Wheel

	err := json.NewDecoder(r.Body).Decode(&wheel)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateWheel(&id, &wheel)

	msg := fmt.Sprintf("Wheel updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getWheel(id *int) (models.Wheel, error) {

	var wheel models.Wheel

	var sqlStatement = `SELECT id, name, short_name FROM wheels WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&wheel.ID, &wheel.Name, &wheel.ShortName)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return wheel, nil
	case nil:
		return wheel, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return wheel, err
}

func getAllWheels() ([]models.Wheel, error) {

	var wheels []models.Wheel

	var sqlStatement = `SELECT id, name, short_name	FROM wheels`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute wheels query %v", err)
	}

	defer rs.Close()

	for rs.Next() {
		var wheel models.Wheel

		err := rs.Scan(&wheel.ID, &wheel.Name, &wheel.ShortName)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		wheels = append(wheels, wheel)
	}

	return wheels, err
}

func deleteWheel(id *int) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM wheels WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete wheel. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createWheel(wheel *models.Wheel) (int, error) {

	sqlStatement := `INSERT INTO wheels (name, short_name) VALUES ($1) RETURNING id`

	var id int

	err := Sql().QueryRow(sqlStatement, wheel.Name, wheel.ShortName).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to create wheel. %v", err)
	}

	wheel.ID = id

	return id, err
}

func updateWheel(id *int, wheel *models.Wheel) int64 {

	sqlStatement := `UPDATE wheels SET name=$2, short_name=$3 WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id, wheel.Name, wheel.ShortName)

	if err != nil {
		log.Fatalf("Unable to update wheel. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating wheel. %v", err)
	}

	return rowsAffected
}
