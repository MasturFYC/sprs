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

func AccTypeGetAll(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	allTypes, err := getAllAccTypes(db)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &allTypes)
}

func AccTypeGetItem(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acc_type, err := getAccType(c.Keys["db"].(*sql.DB), &id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &acc_type)
}

func AccTypeCreate(c *gin.Context) {

	var newType models.AccType

	err := c.BindJSON(&newType)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := createAccType(c.Keys["db"].(*sql.DB), &newType)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": rowsAffected, "message": "Account type created successfully"})

}

func AccTypeUpdate(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var acc_type models.AccType

	err := c.BindJSON(&acc_type)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRows, err := updateAccType(c.Keys["db"].(*sql.DB), &id, &acc_type)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Account type updated successfully. Total rows/record affected %v", updatedRows)

	c.JSON(http.StatusOK, gin.H{"rows": updatedRows, "message": msg})
}

func AccTypeDelete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	deletedRows, err := deleteAccType(db, &id)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Account type deleted successfully. Total rows/record affected %v", deletedRows)

	c.JSON(http.StatusOK, gin.H{"rows": deletedRows, "message": msg})
}

func getAccType(db *sql.DB, id *int) (models.AccType, error) {

	var acc models.AccType

	var sqlStatement = `SELECT 
		group_id, id, name, descriptions
	FROM acc_type
	WHERE id=$1`

	rs := db.QueryRow(sqlStatement, id)

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

func getAllAccTypes(db *sql.DB) ([]models.AccType, error) {

	var results = make([]models.AccType, 0)

	var sqlStatement = `SELECT group_id, id, name, descriptions FROM acc_type ORDER BY id`

	rs, err := db.Query(sqlStatement)

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

func createAccType(db *sql.DB, p *models.AccType) (int64, error) {

	sqlStatement := `INSERT INTO acc_type (group_id, id, name, descriptions) VALUES ($1, $2, $3, $4)`

	res, err := db.Exec(sqlStatement, p.GroupID, p.ID, p.Name, p.Descriptions)

	if err != nil {
		log.Printf("Unable to create account type. %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func updateAccType(db *sql.DB, id *int, p *models.AccType) (int64, error) {

	sqlStatement := `UPDATE acc_type SET
		group_id=$2, name=$3, descriptions=$4
		WHERE id=$1`

	res, err := db.Exec(sqlStatement,
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

	return rowsAffected, err
}

func deleteAccType(db *sql.DB, id *int) (int64, error) {

	sqlStatement := `DELETE FROM acc_type WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to delete account type. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}
