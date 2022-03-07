package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func GetTransactionsByType(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_codes, err := getTransactionsByType(&id)

	if err != nil {
		log.Printf("Unable to get all transactions. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_codes)
}

func SearchTransactions(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	var txt = params["txt"]

	trxs, err := searchTransactions(&txt)

	if err != nil {
		log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&trxs)
}

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

	id, err := strconv.ParseInt(params["id"], 10, 64)

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

	if len(trx.Details) > 0 {

		err = bulkInsertDetails(trx.Details, &id)

		if err != nil {
			log.Printf("Unable to insert transaction details.  %v", err)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
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

	id, err := strconv.ParseInt(params["id"], 10, 64)

	var trx models.Trx

	err = json.NewDecoder(r.Body).Decode(&trx)

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

	if len(trx.Details) > 0 {

		_, err = deleteDetailsByOrder(&id)
		if err != nil {
			log.Printf("Unable to delete all details by transaction.  %v", err)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		var newId int64 = 0

		err = bulkInsertDetails(trx.Details, &newId)

		if err != nil {
			log.Printf("Unable to insert transaction details.  %v", err)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
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

func deleteDetailsByOrder(id *int64) (int64, error) {
	sqlStatement := `DELETE FROM trx_detail WHERE trx_id=$1`

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

func bulkInsertDetails(rows []models.TrxDetail, id *int64) error {
	valueStrings := make([]string, 0, len(rows))
	valueArgs := make([]interface{}, 0, len(rows)*5)
	i := 0
	for _, post := range rows {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		valueArgs = append(valueArgs, post.ID)
		valueArgs = append(valueArgs, post.AccCodeID)
		valueArgs = append(valueArgs, setTrxID(id, &post.TrxID))
		valueArgs = append(valueArgs, post.Debt)
		valueArgs = append(valueArgs, post.Cred)
		i++
	}
	stmt := fmt.Sprintf("INSERT INTO trx_detail (id, acc_code_id, trx_id, debt, cred) VALUES %s",
		strings.Join(valueStrings, ","))
	//log.Printf("%s %v", stmt, valueArgs)
	_, err := Sql().Exec(stmt, valueArgs...)
	return err
}

func setTrxID(id *int64, id2 *int64) int64 {
	if (*id) == 0 {
		return (*id2)
	}
	return (*id)
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

func getTransaction(id *int64) (models.Trx, error) {

	var trx models.Trx

	var sqlStatement = `SELECT 
		id, trx_type_id, ref_id, division, trx_date, descriptions, memo,
		(select sum(d.debt) as debt from trx_detail d where d.trx_id = t.id) saldo
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
		&trx.Saldo,
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
		t.id, t.trx_type_id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,
		(select sum(d.debt) as debt from trx_detail d where d.trx_id = t.id) saldo
	FROM trx t
	ORDER BY t.id DESC`

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
			&p.Saldo,
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
	(trx_type_id, ref_id, division, trx_date, descriptions, memo, trx_token)
	VALUES ($1, $2, $3, $4, $5, $6, to_tsvector('indonesian', $7))
	RETURNING id`

	var id int64
	token := fmt.Sprintf("%d %s %s", p.ID, p.Descriptions, p.Memo)

	err := Sql().QueryRow(sqlStatement,
		p.TrxTypeID,
		p.RefID,
		p.Division,
		p.TrxDate,
		p.Descriptions,
		p.Memo,
		token,
	).Scan(&id)

	if err != nil {
		log.Printf("Unable to create transaction. %v", err)
		return 0, err
	}

	return id, err
}

func updateTransaction(id *int64, p *models.Trx) (int64, error) {

	sqlStatement := `UPDATE trx SET 
		trx_type_id=$2,
		ref_id=$3,
		division=$4,
		trx_date=$5,
		descriptions=$6,
		memo=$7,
		trx_token=to_tsvector('indonesian', $8)
	WHERE id=$1`

	token := fmt.Sprintf("%d %s %s", p.ID, p.Descriptions, p.Memo)

	res, err := Sql().Exec(sqlStatement,
		id,
		p.TrxTypeID,
		p.RefID,
		p.Division,
		p.TrxDate,
		p.Descriptions,
		p.Memo,
		token,
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

func searchTransactions(txt *string) ([]models.Trx, error) {

	var results []models.Trx

	var sqlStatement = `SELECT 
	t.id, t.trx_type_id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,
	(select sum(d.debt) as debt from trx_detail d where d.trx_id = t.id) saldo
	FROM trx t
	WHERE t.trx_token @@ to_tsquery('indonesian', $1)
	ORDER BY t.id DESC`

	rs, err := Sql().Query(sqlStatement, txt)

	if err != nil {
		log.Printf("Unable to execute transactions code query %v", err)
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
			&p.Saldo,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func getTransactionsByType(id *int64) ([]models.Trx, error) {

	var results []models.Trx

	var sqlStatement = `SELECT 
	t.id, t.trx_type_id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,
	(select sum(d.debt) as debt from trx_detail d where d.trx_id = t.id) saldo
	FROM trx t
	WHERE t.trx_type_id=$1
	ORDER BY t.id DESC`

	rs, err := Sql().Query(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to execute transactions query %v", err)
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
			&p.Saldo,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}
