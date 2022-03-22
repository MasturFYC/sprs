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

func GetAccountTypes(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	acc_types, err := getAllAccTypes()

	if err != nil || len(acc_types) == 0 {
		//log.Printf("Unable to get all account types. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_types)
}

func GetAccountType(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_type, err := getAccType(&id)

	if err != nil {
		//log.Printf("Unable to get account type. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_type)
}

func CreateAccountType(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var acc_type models.AccType

	err := json.NewDecoder(r.Body).Decode(&acc_type)

	if err != nil {
		//log.Printf("Unable to decode the request body to account type.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rowsAffected, err := createAccType(&acc_type)

	if err != nil {
		log.Printf("(API) Unable to create account type.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	res := Response{
		ID:      rowsAffected,
		Message: "Account type created successfully",
	}

	json.NewEncoder(w).Encode(&res)

}

func UpdateAccountType(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var acc_type models.AccType

	err := json.NewDecoder(r.Body).Decode(&acc_type)

	if err != nil {
		//log.Printf("Unable to decode the request body to account type.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := updateAccType(&id, &acc_type)

	if err != nil {
		//log.Printf("Unable to update account type.  %v", err)
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

func DeleteAccountType(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows, err := deleteAccType(&id)

	if err != nil {
		//log.Printf("Unable to delete account.  %v", err)
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

func getAccType(id *int) (models.AccType, error) {

	var acc models.AccType

	var sqlStatement = `SELECT 
		group_id, id, name, descriptions
	FROM acc_type
	WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&acc.GroupID, &acc.ID, &acc.Name, acc.Descriptions)

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

func getAllAccTypes() ([]models.AccType, error) {

	var results []models.AccType

	var sqlStatement = `SELECT group_id, id, name, descriptions FROM acc_type ORDER BY id`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute account type query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccType

		err := rs.Scan(&p.GroupID, &p.ID, &p.Name, &p.Descriptions)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func createAccType(p *models.AccType) (int64, error) {

	sqlStatement := `INSERT INTO acc_type (group_id, id, name, descriptions) VALUES ($1, $2, $3, $4)`

	res, err := Sql().Exec(sqlStatement, p.GroupID, p.ID, p.Name, p.Descriptions)

	if err != nil {
		log.Printf("Unable to create account type. %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Unable to create account type. %v", err)
	}

	return rowsAffected, err
}

func updateAccType(id *int, p *models.AccType) (int64, error) {

	sqlStatement := `UPDATE acc_type SET
		group_id=$2, name=$3, descriptions=$4
		WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		p.GroupID,
		p.Name,
		p.Descriptions,
	)

	if err != nil {
		log.Printf("Unable to update account type. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating account type. %v", err)
		return 0, err
	}

	return rowsAffected, err
}

func deleteAccType(id *int) (int64, error) {

	sqlStatement := `DELETE FROM acc_type WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to delete account type. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected, err
}
