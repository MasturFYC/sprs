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

func GetWheels(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	wheels, err := getAllWheels(db)

	if err != nil {
		log.Fatalf("Unable to get all wheels. %v", err)
	}

	c.JSON(http.StatusOK, &wheels)
}

func GetWheel(c *gin.Context) {

	// id = order id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	db := c.Keys["db"].(*sql.DB)
	wheels, err := getWheel(db, &id)

	if err != nil {
		log.Fatalf("Unable to get wheel. %v", err)
	}

	c.JSON(http.StatusOK, &wheels)
}

func DeleteWheel(c *gin.Context) {

	// id = order id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	db := c.Keys["db"].(*sql.DB)
	deletedRows := deleteWheel(db, &id)

	msg := fmt.Sprintf("Wheel deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func CreateWheel(c *gin.Context) {

	var wheel models.Wheel

	err := c.BindJSON(&wheel)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	id, err := createWheel(db, &wheel)

	if err != nil {
		//log.Fatalf("Nama wheel tidak boleh sama.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	wheel.ID = id

	c.JSON(http.StatusOK, &wheel)

}

func UpdateWheel(c *gin.Context) {

	// create the postgres db connection

	id, _ := strconv.Atoi(c.Param("id"))

	var wheel models.Wheel

	err := c.BindJSON(&wheel)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateWheel(db, &id, &wheel)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Wheel updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getWheel(db *sql.DB, id *int) (models.Wheel, error) {

	var wheel models.Wheel

	var sqlStatement = `SELECT id, name, short_name FROM wheels WHERE id=$1`

	rs := db.QueryRow(sqlStatement, id)

	err := rs.Scan(&wheel.ID, &wheel.Name, &wheel.ShortName)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return wheel, nil
	case nil:
		return wheel, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return wheel, err
}

func getAllWheels(db *sql.DB) ([]models.Wheel, error) {

	var wheels []models.Wheel

	// wheels = append(wheels, models.Wheel{ID: 0, Name: "", ShortName: ""})

	var sqlStatement = `SELECT id, name, short_name	FROM wheels ORDER BY short_name`

	rs, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute wheels query %v", err)
	}

	defer rs.Close()

	for rs.Next() {
		var wheel models.Wheel

		err := rs.Scan(&wheel.ID, &wheel.Name, &wheel.ShortName)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		wheels = append(wheels, wheel)
	}

	return wheels, err
}

func deleteWheel(db *sql.DB, id *int) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM wheels WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete wheel. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createWheel(db *sql.DB, wheel *models.Wheel) (int, error) {

	sqlStatement := `INSERT INTO wheels (name, short_name) VALUES ($1, $2) RETURNING id`

	var id int

	err := db.QueryRow(sqlStatement, wheel.Name, wheel.ShortName).Scan(&id)

	if err != nil {
		log.Printf("Unable to create wheel. %v", err)
	}

	return id, err
}

func updateWheel(db *sql.DB, id *int, wheel *models.Wheel) (int64, error) {

	sqlStatement := `UPDATE wheels SET name=$2, short_name=$3 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, wheel.Name, wheel.ShortName)

	if err != nil {
		log.Printf("Unable to update wheel. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating wheel. %v", err)
	}

	return rowsAffected, err
}
