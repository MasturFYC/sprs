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

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	trxs, err := getAllTransactions()

	if err != nil {
		log.Printf("Unable to get all transaction. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&trxs)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	trx, err := getTransaction(&id)

	if err != nil {
		log.Printf("Unable to get transaction. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(&trx)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var trx models.Trx

	err := json.NewDecoder(r.Body).Decode(&trx)

	if err != nil {
		log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	id, err := createTransaction(&trx)

	if err != nil {
		log.Printf("(API) Unable to create transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	res := Response{
		ID:      id,
		Message: "Transaction was succesfully inserted.",
	}

	json.NewEncoder(w).Encode(&res)

}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var trx models.Trx

	err := json.NewDecoder(r.Body).Decode(&trx)

	if err != nil {
		log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := updateTransaction(&id, &trx)

	if err != nil {
		log.Printf("Unable to update transaction.  %v", err)
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

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows, err := deleteTransaction(&id)

	if err != nil {
		log.Printf("Unable to delete transaction.  %v", err)
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

func getTransaction(id *int) (models.Trx, error) {

	var trx models.Trx

	var sqlStatement = `SELECT 
		id, trx_type_id, ref_id, division, trx_date, descriptions, memo
	FROM trx
	WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(
		&trx.ID,
		&trx.TrxTypeID,
		&trx.RefID,
		&trx.Division,
		&trx.TrxDate,
		&trx.Descriptions,
		&trx.Memo,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return trx, nil
	case nil:
		return trx, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return trx, err
}

func getAllTransactions() ([]models.Trx, error) {

	var results []models.Trx

	var sqlStatement = `SELECT 
		id, trx_type_id, ref_id, division, trx_date, descriptions, memo
	FROM trx
	ORDER BY id DESC`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute transaction query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.Trx

		err := rs.Scan(
			&p.ID,
			&p.TrxTypeID,
			&p.RefID,
			&p.Division,
			&p.TrxDate,
			&p.Descriptions,
			&p.Memo,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// d, _ := getTransactionDetails(&p.ID)
		// p.Details = d

		results = append(results, p)
	}

	return results, err
}

func createTransaction(p *models.Trx) (int64, error) {

	sqlStatement := `INSERT INTO trx 
	(trx_type_id, ref_id, division, trx_date, descriptions, memo)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`

	var id int64

	err := Sql().QueryRow(sqlStatement,
		p.TrxTypeID,
		p.RefID,
		p.Division,
		p.TrxDate,
		p.Descriptions,
		p.Memo,
	).Scan(&id)

	if err != nil {
		log.Printf("Unable to create transaction. %v", err)
		return 0, err
	}

	return id, err
}

func updateTransaction(id *int, p *models.Trx) (int64, error) {

	sqlStatement := `UPDATE trx SET 
		trx_type_id=$2,
		ref_id=$3,
		division=$4,
		trx_date=$5,
		descriptions=$6,
		memo=$7
	WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		p.TrxTypeID,
		p.RefID,
		p.Division,
		p.TrxDate,
		p.Descriptions,
		p.Memo,
	)

	if err != nil {
		log.Printf("Unable to update transaction. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating transaction. %v", err)
		return 0, err
	}

	return rowsAffected, err
}

func deleteTransaction(id *int) (int64, error) {

	sqlStatement := `DELETE FROM trx WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to delete transaction. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected, err
}
