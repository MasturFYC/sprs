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

func GetTypes(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	types, err := getAllTypes(db)

	if err != nil {
		log.Fatalf("Unable to get all types. %v", err)
	}

	c.JSON(http.StatusOK, &types)
}

func GetType(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//		log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	types, err := getType(db, &id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &types)
}

func DeleteType(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	deletedRows := deleteType(db, &id)

	msg := fmt.Sprintf("Type deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func CreateType(c *gin.Context) {

	var t models.Type

	err := c.BindJSON(&t)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	id, err := createType(db, &t)

	if err != nil {
		//log.Fatalf("Nama type tidak boleh sama.  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t.ID = id

	c.JSON(http.StatusCreated, &t)

}

func UpdateType(c *gin.Context) {

	// create the postgres db connection

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var t models.Type

	err := c.BindJSON(&t)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateType(db, &id, &t)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Type updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getType(db *sql.DB, id *int64) (models.Type, error) {

	var t models.Type
	var sqlStatement = `SELECT id, name, wheel_id, merk_id FROM types WHERE id=$1`

	rs := db.QueryRow(sqlStatement, id)
	err := rs.Scan(&t.ID, &t.Name, &t.WheelID, &t.MerkID)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return t, nil
	case nil:

		w, err := getWheel(db, &t.WheelID)
		if err == nil {
			t.Wheel = w
		}
		m, err := getMerk(db, &t.MerkID)
		if err == nil {
			t.Merk = m
		}

		return t, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return t, err
}

func getAllTypes(db *sql.DB) ([]models.Type, error) {

	var types []models.Type

	var sqlStatement = `SELECT
		id, name, wheel_id, merk_id
	FROM types
	ORDER BY name`

	rs, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute types query %v", err)
	}

	defer rs.Close()

	for rs.Next() {
		var t models.Type

		err := rs.Scan(&t.ID, &t.Name, &t.WheelID, &t.MerkID)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		wheel, _ := getWheel(db, &t.WheelID)
		merk, _ := getMerk(db, &t.MerkID)
		t.Wheel = wheel
		t.Merk = merk

		types = append(types, t)
	}

	return types, err
}

func deleteType(db *sql.DB, id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM types WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete type. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createType(db *sql.DB, t *models.Type) (int64, error) {

	sqlStatement := `INSERT INTO types (name, wheel_id, merk_id) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, t.Name, t.WheelID, t.MerkID).Scan(&id)

	if err != nil {
		//log.Fatalf("Unable to create type. %v", err)
		return 0, err
	}

	t.ID = id

	return id, err
}

func updateType(db *sql.DB, id *int64, t *models.Type) (int64, error) {

	sqlStatement := `UPDATE types SET name=$2, wheel_id=$3, merk_id=$4 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, t.Name, t.WheelID, t.MerkID)

	if err != nil {
		//log.Fatalf("Unable to update type. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	//log.Fatalf("Error while updating type. %v", err)
	// 	return 0, err
	// }

	return rowsAffected, err
}
