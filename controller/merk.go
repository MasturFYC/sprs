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

	db := c.Keys["db"].(*sql.DB)
	merks, err := getAllMerks(db)

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
	db := c.Keys["db"].(*sql.DB)
	merks, err := getMerk(db, &id)

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

	db := c.Keys["db"].(*sql.DB)
	deletedRows := deleteMerk(db, &id)

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

	db := c.Keys["db"].(*sql.DB)
	id, err := createMerk(db, &merk)

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

	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateMerk(db, &id, &merk)

	if err != nil {
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

func getMerk(db *sql.DB, id *int) (models.Merk, error) {

	var merk models.Merk

	var sqlStatement = `SELECT id, name FROM merks WHERE id=$1`

	rs := db.QueryRow(sqlStatement, id)

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

func getAllMerks(db *sql.DB) ([]models.Merk, error) {

	var merks = make([]models.Merk, 0)

	var sqlStatement = `SELECT id, name FROM merks ORDER BY name`

	rs, err := db.Query(sqlStatement)

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

func deleteMerk(db *sql.DB, id *int) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM merks WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

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

func createMerk(db *sql.DB, merk *models.Merk) (int, error) {

	sqlStatement := `INSERT INTO merks (name) VALUES ($1) RETURNING id`

	var id int

	err := db.QueryRow(sqlStatement, merk.Name).Scan(&id)

	return id, err
}

func updateMerk(db *sql.DB, id *int, merk *models.Merk) (int64, error) {

	sqlStatement := `UPDATE merks SET name=$2 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, merk.Name)

	if err != nil {
		log.Printf("Unable to update merk. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}
