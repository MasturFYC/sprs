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

func GetTransactionTypes(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	trx_types, err := getAllTrxTypes()

	if err != nil {
		log.Printf("Unable to get all transaction types. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&trx_types)
}

func GetTransactionType(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	trx_type, err := getTrxType(&id)

	if err != nil {
		log.Printf("Unable to get transaction type. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(&trx_type)
}

func CreateTransactionType(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var trx_type models.TrxType

	err := json.NewDecoder(r.Body).Decode(&trx_type)

	if err != nil {
		log.Printf("Unable to decode the request body to transaction type.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	rowsAffected, err := createTrxType(&trx_type)

	if err != nil {
		log.Printf("(API) Unable to create transaction type.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	res := Response{
		ID:      rowsAffected,
		Message: "Transaction type created successfully",
	}

	json.NewEncoder(w).Encode(&res)

}

func UpdateTransactionType(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var trx_type models.TrxType

	err := json.NewDecoder(r.Body).Decode(&trx_type)

	if err != nil {
		log.Printf("Unable to decode the request body to transaction type.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := updateTrxType(&id, &trx_type)

	if err != nil {
		log.Printf("Unable to update transaction type.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Transaction type updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func DeleteTransactionType(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows, err := deleteTrxType(&id)

	if err != nil {
		log.Printf("Unable to delete transaction type.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Transaction type deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      deletedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getTrxType(id *int) (models.TrxType, error) {

	var trx_type models.TrxType

	var sqlStatement = `SELECT 
		id, name, descriptions
	FROM trx_type
	WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&trx_type.ID, &trx_type.Name, &trx_type.Descriptions)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return trx_type, nil
	case nil:
		return trx_type, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return trx_type, err
}

func getAllTrxTypes() ([]models.TrxType, error) {

	var results []models.TrxType

	var sqlStatement = `SELECT 
		id, name, descriptions
	FROM trx_type
	ORDER BY id`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute transaction type query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.TrxType

		err := rs.Scan(&p.ID, &p.Name, &p.Descriptions)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func createTrxType(p *models.TrxType) (int64, error) {

	sqlStatement := `INSERT INTO trx_type (id, name, descriptions) VALUES ($1, $2, $3)`

	res, err := Sql().Exec(sqlStatement, p.ID, p.Name, p.Descriptions)

	if err != nil {
		log.Printf("Unable to create transaction type. %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Unable to create transaction type. %v", err)
	}

	return rowsAffected, err
}

func updateTrxType(id *int, p *models.TrxType) (int64, error) {

	sqlStatement := `UPDATE trx_type SET
	id=$2, name=$3, descriptions=$4
	WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		p.ID,
		p.Name,
		p.Descriptions,
	)

	if err != nil {
		log.Printf("Unable to update transaction type. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating transaction type. %v", err)
		return 0, err
	}

	return rowsAffected, err
}

func deleteTrxType(id *int) (int64, error) {

	sqlStatement := `DELETE FROM trx_type WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to delete transaction type. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected, err
}
