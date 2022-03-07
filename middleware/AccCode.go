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

func GetAccountCodeProps(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	acc_codes, err := getAllAccCodeProps()

	if err != nil {
		log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_codes)
}

func GetAccountCodes(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	acc_codes, err := getAllAccCodes()

	if err != nil {
		log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_codes)
}

func SearchAccountCodeByName(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	var txt = params["txt"]

	acc_codes, err := searchAccCodeByName(&txt)

	if err != nil {
		log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_codes)
}

func GetAccountCodeByType(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_codes, err := getAccCodeByType(&id)

	if err != nil {
		log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_codes)
}

func GetAccountCode(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_code, err := getAccCode(&id)

	if err != nil {
		log.Printf("Unable to get account code. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(&acc_code)
}

func CreateAccountCode(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var acc_code models.AccCode

	err := json.NewDecoder(r.Body).Decode(&acc_code)

	if err != nil {
		log.Printf("Unable to decode the request body to account code.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	rowsAffected, err := createAccCode(&acc_code)

	if err != nil {
		log.Printf("(API) Unable to create account code.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	res := Response{
		ID:      rowsAffected,
		Message: "Account type created successfully",
	}

	json.NewEncoder(w).Encode(&res)

}

func UpdateAccountCode(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var acc_code models.AccCode

	err := json.NewDecoder(r.Body).Decode(&acc_code)

	if err != nil {
		log.Printf("Unable to decode the request body to account code.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := updateAccCode(&id, &acc_code)

	if err != nil {
		log.Printf("Unable to update account code.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Account type updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func DeleteAccountCode(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows, err := deleteAccCode(&id)

	if err != nil {
		log.Printf("Unable to delete account.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Account type deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      deletedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getAccCode(id *int) (models.AccCode, error) {

	var acc models.AccCode

	var sqlStatement = `SELECT 
		id, name, acc_type_id, descriptions
	FROM acc_code
	WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&acc.ID, &acc.Name, &acc.AccTypeID, &acc.Descriptions)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return acc, nil
	case nil:
		return acc, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return acc, err
}

func getAccCodeByType(id *int) ([]models.AccCode, error) {

	var results []models.AccCode

	var sqlStatement = `SELECT 
		id, name, acc_type_id, descriptions
	FROM acc_code
	WHERE acc_type_id=$1
	ORDER BY acc_type_id, id`

	rs, err := Sql().Query(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to execute account code query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCode

		err := rs.Scan(&p.ID, &p.Name, &p.AccTypeID, &p.Descriptions)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func searchAccCodeByName(txt *string) ([]models.AccCode, error) {

	var results []models.AccCode

	var sqlStatement = `SELECT 
		id, name, acc_type_id, descriptions
	FROM acc_code
	WHERE token_name @@ to_tsquery($1)
	ORDER BY acc_type_id, id`

	rs, err := Sql().Query(sqlStatement, txt)

	if err != nil {
		log.Printf("Unable to execute account code query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCode

		err := rs.Scan(&p.ID, &p.Name, &p.AccTypeID, &p.Descriptions)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func getAllAccCodes() ([]models.AccCode, error) {

	var results []models.AccCode

	var sqlStatement = `SELECT 
		id, name, acc_type_id, descriptions
	FROM acc_code
	ORDER BY acc_type_id, id`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute account code query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCode

		err := rs.Scan(&p.ID, &p.Name, &p.AccTypeID, &p.Descriptions)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func createAccCode(p *models.AccCode) (int64, error) {

	sqlStatement := `INSERT INTO 
	acc_code (id, name, acc_type_id, descriptions, token_name)
	VALUES ($1, $2, $3, $4, to_tsvector($5))`

	token := fmt.Sprintf("%s %s", p.Name, p.Descriptions)

	res, err := Sql().Exec(sqlStatement,
		p.ID,
		p.Name,
		p.AccTypeID,
		p.Descriptions,
		token,
	)

	if err != nil {
		log.Printf("Unable to create account code. %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Unable to create account code. %v", err)
	}

	return rowsAffected, err
}

func updateAccCode(id *int, p *models.AccCode) (int64, error) {

	sqlStatement := `UPDATE acc_code SET 
	id=$2, name=$3, acc_type_id=$4, descriptions=$5,
	token_name=to_tsvector($6)
	WHERE id=$1`

	token := fmt.Sprintf("%s %s", p.Name, p.Descriptions)

	res, err := Sql().Exec(sqlStatement,
		id,
		p.ID,
		p.Name,
		p.AccTypeID,
		p.Descriptions,
		token,
	)

	if err != nil {
		log.Printf("Unable to update account code. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating account code. %v", err)
		return 0, err
	}

	return rowsAffected, err
}

func deleteAccCode(id *int) (int64, error) {

	sqlStatement := `DELETE FROM acc_code WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to delete account code. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected, err
}

func getAllAccCodeProps() ([]models.AccCodeType, error) {

	var results []models.AccCodeType

	var sqlStatement = `SELECT 
		c.id, c.name, t.id as type_id, t.name AS type_name, c.descriptions
	FROM acc_code c
	INNER JOIN acc_type t ON t.id = c.acc_type_id
	ORDER BY t.id, c.id`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute account code property query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCodeType

		err := rs.Scan(
			&p.ID,
			&p.Name,
			&p.TypeID,
			&p.TypeName,
			&p.Descriptions,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}
