package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"fyc.com/sprs/models"
	"github.com/gin-gonic/gin"
)

func BranchGetAll(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	branchs, err := getAllBranchs(db)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &branchs)
}

func BranchGetItem(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	branch, err := getBranch(db, &id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &branch)
}

func BranchDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	deletedRows := deleteBranch(db, &id)

	msg := fmt.Sprintf("Branch deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func BranchCreate(c *gin.Context) {

	var branch models.Branch

	err := c.BindJSON(&branch)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	id, err := createBranch(db, &branch)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	branch.ID = id

	c.JSON(http.StatusOK, &branch)

}

func BranchUpdate(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var branch models.Branch

	err := c.BindJSON(&branch)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateBranch(db, &id, &branch)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Branch updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getBranch(db *sql.DB, id *int) (models.Branch, error) {

	var branch models.Branch

	var sqlStatement = `SELECT 
		id, name, head_branch, street, city, phone, cell, zip, email
	FROM branchs
	WHERE id=$1`

	rs := db.QueryRow(sqlStatement, id)

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

func getAllBranchs(db *sql.DB) ([]models.Branch, error) {

	var branchs = make([]models.Branch, 0)

	var sqlStatement = `SELECT 
		id, name, head_branch, street, city, phone, cell, zip, email
	FROM branchs
	ORDER BY name`

	rs, err := db.Query(sqlStatement)

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

func deleteBranch(db *sql.DB, id *int) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM branchs WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

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

func createBranch(db *sql.DB, branch *models.Branch) (int, error) {

	sqlStatement := `INSERT INTO branchs
		(name, head_branch, street, city, phone, cell, zip, email) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`

	var id int

	err := db.QueryRow(sqlStatement,
		branch.Name,
		branch.HeadBranch,
		branch.Street,
		branch.City,
		branch.Phone,
		branch.Cell,
		branch.Zip,
		branch.Email,
	).Scan(&id)

	return id, err
}

func updateBranch(db *sql.DB, id *int, branch *models.Branch) (int64, error) {

	sqlStatement := `UPDATE branchs SET
		name=$2, head_branch=$3, street=$4, city=$5,
		phone=$6, cell=$7, zip=$8, email=$9
	WHERE id=$1`

	res, err := db.Exec(sqlStatement,
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

	return rowsAffected, err
}
