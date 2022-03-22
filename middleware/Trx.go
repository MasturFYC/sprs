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

type local_detail struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	AccCodeID int32   `json:"accCodeId"`
	Debt      float64 `json:"debt"`
	Cred      float64 `json:"cred"`
}

type local_trx struct {
	models.Trx
	Details json.RawMessage `json:"details,omitempty"`
}

func GetTransactionsByMonth(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_codes, err := get_trx_by_month(&id)

	if err != nil {
		//log.Printf("Unable to get all transactions. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_codes)
}

func GetTransactionsByGroup(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_codes, err := getTransactionsByGroup(&id)

	if err != nil {
		//log.Printf("Unable to get all transactions. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_codes)
}

func SearchTransactions(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	var t models.SearchGroup

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	trxs, err := searchTransactions(&t.Txt)

	if err != nil {
		//log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&trxs)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	trxs, err := get_all_transactions()

	if err != nil {
		//log.Printf("Unable to get all transaction. %v", err)
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
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	trx, err := getTransaction(&id)

	if err != nil {
		//log.Printf("Unable to get transaction. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&trx)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var param models.TrxDetailsToken

	err := json.NewDecoder(r.Body).Decode(&param)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := createTransaction(&param.Trx, param.Token)

	if err != nil {
		//log.Printf("(API) Unable to create transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(param.Details) > 0 {

		err = bulkInsertDetails(param.Details, &id)

		if err != nil {
			//log.Printf("Unable to insert transaction details.  %v", err)
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

	var trx models.TrxDetailsToken

	err = json.NewDecoder(r.Body).Decode(&trx)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := updateTransaction(&id, &trx.Trx, trx.Token)

	if err != nil {
		//log.Printf("Unable to update transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// if len(trx.Details) > 0 {

	// 	_, err = deleteDetailsByOrder(&id)
	// 	if err != nil {
	// 		log.Printf("Unable to delete all details by transaction.  %v", err)
	// 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 		return
	// 	}

	// 	var newId int64 = 0

	err = bulkInsertDetails(trx.Details, &id)

	if err != nil {
		//log.Printf("Unable to insert transaction details (message from command).  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	//	}

	msg := fmt.Sprintf("Transaction updated successfully. Total rows/record affected %v", updatedRows)

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
		//log.Printf("Unable to delete transaction. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Fatalf("Error while checking the affected rows. %v", err)
	// }

	return rowsAffected, err

}

func bulkInsertDetails(rows []models.TrxDetail, id *int64) error {
	valueStrings := make([]string, 0, len(rows))
	valueArgs := make([]interface{}, 0, len(rows)*5)
	i := 0
	for _, post := range rows {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		valueArgs = append(valueArgs, post.ID)
		valueArgs = append(valueArgs, post.CodeID)
		valueArgs = append(valueArgs, setTrxID(id, &post.TrxID))
		valueArgs = append(valueArgs, post.Debt)
		valueArgs = append(valueArgs, post.Cred)
		i++
	}
	stmt := fmt.Sprintf("INSERT INTO trx_detail (id, code_id, trx_id, debt, cred) VALUES %s",
		strings.Join(valueStrings, ","))
	//log.Printf("%s %v", stmt, valueArgs)
	_, err := Sql().Exec(stmt+" ON CONFLICT (trx_id, id) DO UPDATE SET code_id=EXCLUDED.code_id, trx_id=EXCLUDED.trx_id, debt=EXCLUDED.debt, cred=EXCLUDED.cred", valueArgs...)
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

	msg := fmt.Sprintf("Transaction deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      deletedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getTransaction(id *int64) (local_trx, error) {

	var p local_trx

	var sqlStatement = `SELECT t.id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,
(SELECT COALESCE(sum(s.debt),0) AS debt FROM trx_detail s WHERE s.trx_id = t.id) saldo,
COALESCE
(
(
SELECT array_to_json(array_agg(row_to_json(x))) FROM
(
SELECT c.id, c.name, d.code_id, d.debt, d.cred
FROM trx_detail d
INNER JOIN acc_code c ON c.id = d.code_id
WHERE d.trx_id=t.id 
ORDER BY d.id
) x
),
'[]'
) AS details
FROM trx t
WHERE t.id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(
		&p.ID,
		&p.RefID,
		&p.Division,
		&p.TrxDate,
		&p.Descriptions,
		&p.Memo,
		&p.Saldo,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return p, nil
	case nil:
		return p, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// d, _ := get_details(&p.ID)
	// p.Details = d

	// return empty user on error
	return p, err
}

func get_all_transactions() ([]local_trx, error) {

	var results []local_trx

	var sqlStatement = `SELECT t.id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,
(SELECT COALESCE(sum(s.debt),0) AS debt FROM trx_detail s WHERE s.trx_id = t.id) saldo,
COALESCE
(
(
SELECT array_to_json(array_agg(row_to_json(x))) FROM
(
SELECT c.id, c.name, d.code_id, d.debt, d.cred
FROM trx_detail d
INNER JOIN acc_code c ON c.id = d.code_id
WHERE d.trx_id=t.id 
ORDER BY d.id
) x
),
'[]'
) AS details
			
FROM trx t
ORDER BY t.id DESC`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute transaction query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p local_trx

		err := rs.Scan(
			&p.ID,
			&p.RefID,
			&p.Division,
			&p.TrxDate,
			&p.Descriptions,
			&p.Memo,
			&p.Saldo,
			&p.Details,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// d, _ := get_details(&p.ID)
		// p.Details = d

		results = append(results, p)
	}

	return results, err
}

func createTransaction(p *models.Trx, token string) (int64, error) {

	sqlStatement := `INSERT INTO trx 
	(ref_id, division, trx_date, descriptions, memo, trx_token)
	VALUES ($1, $2, $3, $4, $5, to_tsvector('indonesian', $6))
	RETURNING id`

	var id int64

	err := Sql().QueryRow(sqlStatement,
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

func updateTransaction(id *int64, p *models.Trx, token string) (int64, error) {

	sqlStatement := `UPDATE trx SET 
		ref_id=$2,
		division=$3,
		trx_date=$4,
		descriptions=$5,
		memo=$6,
		trx_token=to_tsvector('indonesian', $7)
	WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
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

func searchTransactions(txt *string) ([]local_trx, error) {

	var results []local_trx

	var sqlStatement = `SELECT t.id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,
(SELECT COALESCE(sum(s.debt),0) AS debt FROM trx_detail s WHERE s.trx_id = t.id) saldo,
COALESCE
(
(
SELECT array_to_json(array_agg(row_to_json(x))) FROM
(
SELECT c.id, c.name, d.code_id, d.debt, d.cred
FROM trx_detail d
INNER JOIN acc_code c ON c.id = d.code_id
WHERE d.trx_id=t.id 
ORDER BY d.id
) x
),
'[]'
) AS details
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
		var p local_trx

		err := rs.Scan(
			&p.ID,
			&p.RefID,
			&p.Division,
			&p.TrxDate,
			&p.Descriptions,
			&p.Memo,
			&p.Saldo,
			&p.Details,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// d, _ := get_details(&p.ID)
		// p.Details = d

		results = append(results, p)
	}

	return results, err
}

func getTransactionsByGroup(id *int64) ([]local_trx, error) {

	var results []local_trx

	var sqlStatement = `SELECT t.id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,
(SELECT COALESCE(sum(s.debt),0) AS debt FROM trx_detail s WHERE s.trx_id = t.id) saldo,
COALESCE
(
(
SELECT array_to_json(array_agg(row_to_json(x))) FROM
(
SELECT c.id, c.name, d.code_id, d.debt, d.cred
FROM trx_detail d
INNER JOIN acc_code c ON c.id = d.code_id
WHERE d.trx_id=t.id 
ORDER BY d.id
) x
),
'[]'
) AS details
FROM trx t
INNER JOIN trx_detail d ON d.trx_id = d.id
INNER JOIN acc_type e ON e.id = d.code_id
INNER JOIN acc_code c ON c.id = e.group_id
WHERE c.group_id=$1
ORDER BY t.id DESC`

	rs, err := Sql().Query(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to execute transactions query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p local_trx

		err := rs.Scan(
			&p.ID,
			&p.RefID,
			&p.Division,
			&p.TrxDate,
			&p.Descriptions,
			&p.Memo,
			&p.Saldo,
			&p.Details,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		// d, _ := get_details(&p.ID)
		// p.Details = d

		results = append(results, p)
	}

	return results, err
}

/*
func get_details(trxID *int64) ([]local_detail, error) {

	var details []local_detail

	var sqlStatement = `SELECT
	c.id, c.name, d.code_id, d.debt, d.cred
	FROM trx_detail d
	INNER JOIN acc_code c ON c.id = d.code_id
	WHERE d.trx_id=$1
	-- AND c.receivable_option != 1
	ORDER BY d.id`

	rs, err := Sql().Query(sqlStatement, trxID)

	if err != nil {
		log.Printf("Unable to execute transaction details query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p local_detail

		err := rs.Scan(
			&p.ID,
			&p.Name,
			&p.AccCodeID,
			&p.Debt,
			&p.Cred,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		details = append(details, p)
	}

	return details, err
}
*/
func get_trx_by_month(id *int) ([]local_trx, error) {

	var results []local_trx

	var sqlStatement = `SELECT t.id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,
(SELECT COALESCE(sum(s.debt),0) AS debt FROM trx_detail s WHERE s.trx_id = t.id) saldo,
COALESCE
(
(
SELECT array_to_json(array_agg(row_to_json(x))) FROM
(
SELECT c.id, c.name, d.code_id, d.debt, d.cred
FROM trx_detail d
INNER JOIN acc_code c ON c.id = d.code_id
WHERE d.trx_id=t.id 
ORDER BY d.id
) x
),
'[]'
) AS details
FROM trx t
WHERE EXTRACT(MONTH from t.trx_date)=$1
ORDER BY t.id DESC`

	rs, err := Sql().Query(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to execute transactions query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p local_trx

		err := rs.Scan(
			&p.ID,
			&p.RefID,
			&p.Division,
			&p.TrxDate,
			&p.Descriptions,
			&p.Memo,
			&p.Saldo,
			&p.Details,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		// d, _ := get_details(&p.ID)
		// p.Details = d

		results = append(results, p)
	}

	return results, err
}
