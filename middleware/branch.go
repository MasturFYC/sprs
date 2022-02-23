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

func GetBranchs(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	branchs, err := getAllBranchs()

	if err != nil {
		log.Fatalf("Unable to get all branchs. %v", err)
	}

	json.NewEncoder(w).Encode(&branchs)
}

func GetBranch(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	branch, err := getBranch(&id)

	if err != nil {
		log.Fatalf("Unable to get branch. %v", err)
	}

	json.NewEncoder(w).Encode(&branch)
}

func DeleteBranch(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteBranch(&id)

	msg := fmt.Sprintf("Branch deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreateBranch(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var branch models.Branch

	err := json.NewDecoder(r.Body).Decode(&branch)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	id, err := createBranch(&branch)

	if err != nil {
		log.Fatalf("Nama Branch tidak boleh sama.  %v", err)
	}

	branch.ID = id

	json.NewEncoder(w).Encode(&branch)

}

func UpdateBranch(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var branch models.Branch

	err := json.NewDecoder(r.Body).Decode(&branch)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateBranch(&id, &branch)

	msg := fmt.Sprintf("Branch updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getBranch(id *int) (models.Branch, error) {

	var branch models.Branch

	var sqlStatement = `SELECT 
		id, name, head_branch, street, city, phone, cell, zip, email
	FROM branchs
	WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&branch.ID, &branch.Name, &branch.HeadBranch, &branch.Street,
		&branch.City, &branch.Phone, &branch.Cell, &branch.Zip, &branch.Email)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return branch, nil
	case nil:
		return branch, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return branch, err
}

func getAllBranchs() ([]models.Branch, error) {

	var branchs []models.Branch

	var sqlStatement = `SELECT 
		id, name, head_branch, street, city, phone, cell, zip, email
	FROM branchs
	ORDER BY name`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute branch query %v", err)
	}

	defer rs.Close()

	for rs.Next() {
		var branch models.Branch

		err := rs.Scan(&branch.ID, &branch.Name, &branch.HeadBranch, &branch.Street,
			&branch.City, &branch.Phone, &branch.Cell, &branch.Zip, &branch.Email)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		branchs = append(branchs, branch)
	}

	return branchs, err
}

func deleteBranch(id *int) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM branchs WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete branch. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createBranch(branch *models.Branch) (int, error) {

	sqlStatement := `INSERT INTO branchs
		(name, head_branch, street, city, phone, cell, zip, email) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`

	var id int

	err := Sql().QueryRow(sqlStatement,
		branch.Name,
		branch.HeadBranch,
		branch.Street,
		branch.City,
		branch.Phone,
		branch.Cell,
		branch.Zip,
		branch.Email,
	).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to create branch. %v", err)
	}

	return id, err
}

func updateBranch(id *int, branch *models.Branch) int64 {

	sqlStatement := `UPDATE branchs SET
		name=$2, head_branch=$3, street=$4, city=$5,
		phone=$6, cell=$7, zip=$8, email=$9
	WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		branch.Name,
		branch.HeadBranch,
		branch.Street,
		branch.City,
		branch.Phone,
		branch.Cell,
		branch.Zip,
		branch.Email,
	)

	if err != nil {
		log.Fatalf("Unable to update branch. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating category. %v", err)
	}

	return rowsAffected
}
