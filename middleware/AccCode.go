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

type account_specific struct {
	ID           int32             `json:"id"`
	Name         string            `json:"name"`
	Descriptions models.NullString `json:"descriptions"`
}

func Account_GetSpec(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	spec_id, err := strconv.Atoi(params["id"])

	accounts, err := get_accounts_spec(&spec_id)

	if err != nil || len(accounts) == 0 {
		log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&accounts)
}

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

	if err != nil || len(acc_codes) == 0 {
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

	if err != nil || len(acc_codes) == 0 {
		//log.Printf("Unable to get all account codes. %v", err)
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
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_codes, err := getAccCodeByType(&id)

	if err != nil || len(acc_codes) == 0 {
		//log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
		//var test []models.AccCode
		//json.NewEncoder(w).Encode(test)
		//return
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

func getAccCode(id *int) (models.AccInfo, error) {

	var p models.AccInfo

	var sqlStatement = `SELECT 
		c.type_id, c.id, c.name, c.descriptions, c.is_active, c.is_auto_debet, c.receivable_option,
		t.name as type_name, g.name as group_name, t.descriptions as type_desc, g.descriptions as group_desc

	FROM acc_code c
	INNER JOIN acc_type t ON t.id = c.type_id
	INNER JOIN acc_group g ON g.id = t.group_id
	WHERE c.id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(
		&p.TypeID,
		&p.ID,
		&p.Name,
		&p.Descriptions,
		&p.IsActive,
		&p.IsAutoDebet,
		&p.ReceivableOption,
		&p.TypeName,
		&p.GroupName,
		&p.TypeDesc,
		&p.GroupDesc,
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

	// return empty user on error
	return p, err
}

func getAccCodeByType(id *int) ([]models.AccCode, error) {

	var results []models.AccCode

	var sqlStatement = `SELECT 
		type_id, id, name, descriptions, is_active, is_auto_debet, receivable_option
	FROM acc_code
	WHERE type_id=$1 OR 0=$1
	ORDER BY id`

	rs, err := Sql().Query(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to execute account code query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCode

		err := rs.Scan(
			&p.TypeID,
			&p.ID,
			&p.Name,
			&p.Descriptions,
			&p.IsActive,
			&p.IsAutoDebet,
			&p.ReceivableOption,
		)

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
		type_id, id, name, descriptions, is_active, is_auto_debet, receivable_option
	FROM acc_code
	WHERE token_name @@ to_tsquery('indonesian', $1)
	ORDER BY id`

	rs, err := Sql().Query(sqlStatement, txt)

	if err != nil {
		log.Printf("Unable to execute account code query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCode

		err := rs.Scan(
			&p.TypeID,
			&p.ID,
			&p.Name,
			&p.Descriptions,
			&p.IsActive,
			&p.IsAutoDebet,
			&p.ReceivableOption,
		)

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
		type_id, id, name, descriptions, is_active, is_auto_debet, receivable_option
	FROM acc_code
	ORDER BY id`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute account code query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCode

		err := rs.Scan(
			&p.TypeID,
			&p.ID,
			&p.Name,
			&p.Descriptions,
			&p.IsActive,
			&p.IsAutoDebet,
			&p.ReceivableOption,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func createAccCode(p *models.AccCode) (int64, error) {

	sqlStatement := `INSERT INTO 
	acc_code (type_id, id, name, descriptions, is_active, is_auto_debet, receivable_option, token_name)
	VALUES ($1, $2, $3, $4, $5, $6, $7, to_tsvector('indonesian', $8))`

	token := fmt.Sprintf("%s %s", p.Name, p.Descriptions)

	res, err := Sql().Exec(sqlStatement,
		p.TypeID,
		p.ID,
		p.Name,
		p.Descriptions,
		p.IsActive,
		p.IsAutoDebet,
		p.ReceivableOption,
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
	type_id=$2, name=$3, descriptions=$4, is_active=$5, is_auto_debet=$6,
	receivable_option=$7,
	token_name=to_tsvector('indonesian', $8)	
	WHERE id=$1`

	token := fmt.Sprintf("%s %s", p.Name, p.Descriptions)

	res, err := Sql().Exec(sqlStatement,
		id,
		p.TypeID,
		p.Name,
		p.Descriptions,
		p.IsActive,
		p.IsAutoDebet,
		p.ReceivableOption,
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
		c.type_id, c.id, c.name, t.name AS type_name, c.descriptions, c.is_active
	FROM acc_code c
	INNER JOIN acc_type t ON t.id = c.type_id
	ORDER BY c.id`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute account code property query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCodeType

		err := rs.Scan(
			&p.TypeID,
			&p.ID,
			&p.Name,
			&p.TypeName,
			&p.Descriptions,
			&p.IsActive,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func get_accounts_spec(specId *int) ([]account_specific, error) {

	var results []account_specific

	var sqlStatement = `SELECT 
		id, name, descriptions
	FROM acc_code
	WHERE receivable_option = $1
	ORDER BY id`

	rs, err := Sql().Query(sqlStatement, specId)

	if err != nil {
		log.Printf("Unable to execute account code property query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p account_specific

		err := rs.Scan(
			&p.ID,
			&p.Name,
			&p.Descriptions,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}
