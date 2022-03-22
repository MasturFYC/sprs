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

func GetTypes(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	types, err := getAllTypes()

	if err != nil {
		log.Fatalf("Unable to get all types. %v", err)
	}

	json.NewEncoder(w).Encode(&types)
}

func GetType(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	types, err := getType(&id)

	if err != nil {
		log.Fatalf("Unable to get type. %v", err)
	}

	json.NewEncoder(w).Encode(&types)
}

func DeleteType(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteType(&id)

	msg := fmt.Sprintf("Type deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreateType(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var t models.Type

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := createType(&t)

	if err != nil {
		//log.Fatalf("Nama type tidak boleh sama.  %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	t.ID = id

	json.NewEncoder(w).Encode(&t)

}

func UpdateType(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var t models.Type

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := updateType(&id, &t)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("Type updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getType(id *int64) (models.Type, error) {

	var t models.Type

	var sqlStatement = `SELECT id, name, wheel_id, merk_id FROM types WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&t.ID, &t.Name, &t.WheelID, &t.MerkID)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return t, nil
	case nil:

		w, err := getWheel(&t.WheelID)
		if err == nil {
			t.Wheel = w
		}
		m, err := getMerk(&t.MerkID)
		if err == nil {
			t.Merk = m
		}

		return t, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return t, err
}

func getAllTypes() ([]models.Type, error) {

	var types []models.Type

	var sqlStatement = `SELECT
		id, name, wheel_id, merk_id
	FROM types
	ORDER BY name`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute types query %v", err)
	}

	defer rs.Close()

	for rs.Next() {
		var t models.Type

		err := rs.Scan(&t.ID, &t.Name, &t.WheelID, &t.MerkID)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		wheel, _ := getWheel(&t.WheelID)
		merk, _ := getMerk(&t.MerkID)
		t.Wheel = wheel
		t.Merk = merk

		types = append(types, t)
	}

	return types, err
}

func deleteType(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM types WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete type. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createType(t *models.Type) (int64, error) {

	sqlStatement := `INSERT INTO types (name, wheel_id, merk_id) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := Sql().QueryRow(sqlStatement, t.Name, t.WheelID, t.MerkID).Scan(&id)

	if err != nil {
		//log.Fatalf("Unable to create type. %v", err)
		return 0, err
	}

	t.ID = id

	return id, err
}

func updateType(id *int64, t *models.Type) (int64, error) {

	sqlStatement := `UPDATE types SET name=$2, wheel_id=$3, merk_id=$4 WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id, t.Name, t.WheelID, t.MerkID)

	if err != nil {
		//log.Fatalf("Unable to update type. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	//log.Fatalf("Error while updating type. %v", err)
	// 	return 0, err
	// }

	return rowsAffected, err
}
