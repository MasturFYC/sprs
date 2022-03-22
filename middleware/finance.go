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

func GetFinances(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	finances, err := getAllFinances()

	if err != nil {
		//		log.Fatalf("Unable to get all finances. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&finances)
}

func GetFinance(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	finances, err := getFinance(&id)

	if err != nil {
		//log.Fatalf("Unable to get finance. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&finances)
}

func DeleteFinance(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows, err := deleteFinance(&id)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Finance deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreateFinance(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var finance models.Finance

	err := json.NewDecoder(r.Body).Decode(&finance)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := createFinance(&finance)

	if err != nil {
		//log.Fatalf("Nama finance tidak boleh sama.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	finance.ID = id

	json.NewEncoder(w).Encode(&finance)

}

func UpdateFinance(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var finance models.Finance

	err := json.NewDecoder(r.Body).Decode(&finance)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := updateFinance(&id, &finance)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Finance updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getFinance(id *int) (models.Finance, error) {

	var finance models.Finance

	var sqlStatement = `SELECT 
		id, name, short_name, street, city, phone, cell, zip, email
	FROM finances
	WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&finance.ID, &finance.Name, &finance.ShortName, &finance.Street,
		&finance.City, &finance.Phone, &finance.Cell, &finance.Zip, &finance.Email)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return finance, nil
	case nil:
		return finance, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return finance, err
}

func getAllFinances() ([]models.Finance, error) {

	var finances []models.Finance

	var sqlStatement = `SELECT 
		id, name, short_name, street, city, phone, cell, zip, email
	FROM finances
	ORDER BY name`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		// log.Fatalf("Unable to execute finances query %v", err)
		return finances, err
	}

	defer rs.Close()

	for rs.Next() {
		var finance models.Finance

		err := rs.Scan(&finance.ID, &finance.Name, &finance.ShortName, &finance.Street,
			&finance.City, &finance.Phone, &finance.Cell, &finance.Zip, &finance.Email)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		finances = append(finances, finance)
	}

	return finances, err
}

func deleteFinance(id *int) (int64, error) {
	// create the delete sql query
	sqlStatement := `DELETE FROM finances WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		//log.Fatalf("Unable to delete finance. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Fatalf("Error while checking the affected rows. %v", err)
	// }

	return rowsAffected, err
}

func createFinance(finance *models.Finance) (int, error) {

	sqlStatement := `INSERT INTO finances 
	(name, short_name, street, city, phone, cell, zip, email) 
	VALUES 
	($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`

	var id int

	err := Sql().QueryRow(sqlStatement,
		finance.Name,
		finance.ShortName,
		finance.Street,
		finance.City,
		finance.Phone,
		finance.Cell,
		finance.Zip,
		finance.Email,
	).Scan(&id)

	// if err != nil {
	// 	log.Printf("Unable to create finance. %v", err)
	// }

	return id, err
}

func updateFinance(id *int, finance *models.Finance) (int64, error) {

	sqlStatement := `UPDATE finances SET
		name=$2, short_name=$3, street=$4, city=$5, phone=$6, cell=$7, zip=$8, email=$9
	WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		finance.Name,
		finance.ShortName,
		finance.Street,
		finance.City,
		finance.Phone,
		finance.Cell,
		finance.Zip,
		finance.Email,
	)

	if err != nil {
		//log.Printf("Unable to update finance. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Printf("Error while updating finance. %v", err)
	// }

	return rowsAffected, err
}
