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

func GetAccGroups(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	groups, err := getAllAccGroups()

	if err != nil {
		log.Printf("Unable to get all account groups. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&groups)
}

func GetAccGroup(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_group, err := getAccGroup(&id)

	if err != nil {
		log.Printf("Unable to get account group. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(&acc_group)
}

func CreateAccGroup(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var acc_group models.AccGroup

	err := json.NewDecoder(r.Body).Decode(&acc_group)

	if err != nil {
		log.Printf("Unable to decode the request body to account group.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	rowsAffected, err := createAccGroup(&acc_group)

	if err != nil {
		log.Printf("(API) Unable to create account group.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	res := Response{
		ID:      rowsAffected,
		Message: "Account group created successfully",
	}

	json.NewEncoder(w).Encode(&res)

}

func UpdateAccGroup(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var acc_group models.AccGroup

	err := json.NewDecoder(r.Body).Decode(&acc_group)

	if err != nil {
		log.Printf("Unable to decode the request body to account group.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := updateAccGroup(&id, &acc_group)

	if err != nil {
		log.Printf("Unable to update account group.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Account group updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func DeleteAccGroup(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows, err := deleteAccGroup(&id)

	if err != nil {
		log.Printf("Unable to delete account group.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Account group deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      deletedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getAccGroup(id *int) (models.AccGroup, error) {

	var acc_group models.AccGroup

	var sqlStatement = `SELECT 
		id, name, descriptions
	FROM acc_group
	WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&acc_group.ID, &acc_group.Name, &acc_group.Descriptions)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return acc_group, nil
	case nil:
		return acc_group, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return acc_group, err
}

func getAllAccGroups() ([]models.AccGroup, error) {

	var results []models.AccGroup

	var sqlStatement = `SELECT 
		id, name, descriptions
	FROM acc_group
	ORDER BY id`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute account group query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccGroup

		err := rs.Scan(&p.ID, &p.Name, &p.Descriptions)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func createAccGroup(p *models.AccGroup) (int64, error) {

	sqlStatement := `INSERT INTO acc_group (id, name, descriptions) VALUES ($1, $2, $3)`

	res, err := Sql().Exec(sqlStatement, p.ID, p.Name, p.Descriptions)

	if err != nil {
		log.Printf("Unable to create account group. %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Unable to create account group. %v", err)
	}

	return rowsAffected, err
}

func updateAccGroup(id *int, p *models.AccGroup) (int64, error) {

	sqlStatement := `UPDATE acc_group SET
	id=$2, name=$3, descriptions=$4
	WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		p.ID,
		p.Name,
		p.Descriptions,
	)

	if err != nil {
		log.Printf("Unable to update account group. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating account group. %v", err)
		return 0, err
	}

	return rowsAffected, err
}

func deleteAccGroup(id *int) (int64, error) {

	sqlStatement := `DELETE FROM acc_group WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to delete account group. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected, err
}
