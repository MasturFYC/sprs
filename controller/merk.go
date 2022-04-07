package controller

import (
	"database/sql"
	"fmt"
	"log"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func MerkGetAll(c *gin.Context) {

	merks, err := getAllMerks()

	if err != nil {
		log.Fatalf("Unable to get all merks. %v", err)
	}

	c.JSON(http.StatusOK, &merks)
}

func MerkGetItem(c *gin.Context) {

	// id = order id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	merks, err := getMerk(&id)

	if err != nil {
		log.Fatalf("Unable to get merk. %v", err)
	}

	c.JSON(http.StatusOK, &merks)
}

func MerkDelete(c *gin.Context) {

	// id = order id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteMerk(&id)

	msg := fmt.Sprintf("Merk deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func MerkCreate(c *gin.Context) {

	var merk models.Merk

	err := c.BindJSON(&merk)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := createMerk(&merk)

	if err != nil {
		//log.Fatalf("Nama merk tidak boleh sama.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	merk.ID = id

	c.JSON(http.StatusOK, &merk)

}

func MerkUpdate(c *gin.Context) {

	// create the postgres db connection

	id, _ := strconv.Atoi(c.Param("id"))

	var merk models.Merk

	err := c.BindJSON(&merk)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRows := updateMerk(&id, &merk)

	if updatedRows == 0 {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Merk updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getMerk(id *int) (models.Merk, error) {

	var merk models.Merk

	var sqlStatement = `SELECT id, name FROM merks WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&merk.ID, &merk.Name)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return merk, nil
	case nil:
		return merk, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return merk, err
}

func getAllMerks() ([]models.Merk, error) {

	var merks []models.Merk

	var sqlStatement = `SELECT id, name FROM merks ORDER BY name`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute merks query %v", err)
	}

	defer rs.Close()

	for rs.Next() {
		var merk models.Merk

		err := rs.Scan(&merk.ID, &merk.Name)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		merks = append(merks, merk)
	}

	return merks, err
}

func deleteMerk(id *int) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM merks WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete merk. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createMerk(merk *models.Merk) (int, error) {

	sqlStatement := `INSERT INTO merks (name) VALUES ($1) RETURNING id`

	var id int

	err := Sql().QueryRow(sqlStatement, merk.Name).Scan(&id)

	if err != nil {
		log.Printf("Unable to create merk. %v\n", err)
	}

	return id, err
}

func updateMerk(id *int, merk *models.Merk) int64 {

	sqlStatement := `UPDATE merks SET name=$2 WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id, merk.Name)

	if err != nil {
		log.Printf("Unable to update merk. %v", err)
		return 0
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating merk. %v", err)
	}

	return rowsAffected
}
