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

func GetWarehouses(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	warehouses, err := getAllWarehouses(db)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
		//log.Fatalf("Unable to get all warehouses. %v", err)
	}

	c.JSON(http.StatusOK, &warehouses)
}

func GetWarehouse(c *gin.Context) {

	// id = order id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		//		log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	warehouses, err := getWarehouse(db, &id)

	if err != nil {
		//log.Fatalf("Unable to get warehouse. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &warehouses)
}

func DeleteWarehouse(c *gin.Context) {

	// id = order id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	deletedRows, err := deleteWarehouse(db, &id)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Warehouse deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func CreateWarehouse(c *gin.Context) {

	var warehouse models.Warehouse

	err := c.BindJSON(&warehouse)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	id, err := createWarehouse(db, &warehouse)

	if err != nil {
		//log.Fatalf("Nama warehouse tidak boleh sama.  %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	warehouse.ID = id

	c.JSON(http.StatusOK, &warehouse)

}

func UpdateWarehouse(c *gin.Context) {

	// create the postgres db connection

	id, _ := strconv.Atoi(c.Param("id"))

	var warehouse models.Warehouse

	err := c.BindJSON(&warehouse)

	if err != nil {
		//		log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateWarehouse(db, &id, &warehouse)

	if err != nil {
		//		log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return

	}

	msg := fmt.Sprintf("Warehouse updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getWarehouse(db *sql.DB, id *int) (models.Warehouse, error) {

	var warehouse models.Warehouse

	var sqlStatement = `SELECT id, name, descriptions FROM warehouses WHERE id=$1`

	rs := db.QueryRow(sqlStatement, id)

	err := rs.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Descriptions)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return warehouse, nil
	case nil:
		return warehouse, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return warehouse, err
}

func getAllWarehouses(db *sql.DB) ([]models.Warehouse, error) {

	var warehouses = make([]models.Warehouse, 0)

	var sqlStatement = `SELECT id, name, descriptions FROM warehouses`

	rs, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute warehouses query %v", err)
	}

	defer rs.Close()

	for rs.Next() {
		var warehouse models.Warehouse

		err := rs.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Descriptions)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		warehouses = append(warehouses, warehouse)
	}

	return warehouses, err
}

func deleteWarehouse(db *sql.DB, id *int) (int64, error) {
	// create the delete sql query
	sqlStatement := `DELETE FROM warehouses WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		//log.Fatalf("Unable to delete warehouse. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Fatalf("Error while checking the affected rows. %v", err)
	// }

	return rowsAffected, err
}

func createWarehouse(db *sql.DB, warehouse *models.Warehouse) (int, error) {

	sqlStatement := `INSERT INTO warehouses (name, descriptions) VALUES ($1, $2) RETURNING id`

	var id int

	err := db.QueryRow(sqlStatement,
		warehouse.Name,
		warehouse.Descriptions,
	).Scan(&id)

	if err != nil {
		//log.Fatalf("Unable to create warehouse. %v", err)
		return 0, err
	}

	warehouse.ID = id

	return id, nil
}

func updateWarehouse(db *sql.DB, id *int, warehouse *models.Warehouse) (int64, error) {

	sqlStatement := `UPDATE warehouses SET name=$2, descriptions=$3 WHERE id=$1`

	res, err := db.Exec(sqlStatement,
		id,
		warehouse.Name,
		warehouse.Descriptions,
	)

	if err != nil {
		//log.Fatalf("Unable to update warehouse. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Fatalf("Error while updating warehouse. %v", err)
	// }

	return rowsAffected, err
}
