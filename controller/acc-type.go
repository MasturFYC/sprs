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

	allTypes, err := getAllAccTypes()

	if err != nil || len(allTypes) == 0 {
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

	acc_type, err := getAccType(&id)

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

	rowsAffected, err := createAccType(&newType)

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

	updatedRows, err := updateAccType(&id, &acc_type)

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

	deletedRows, err := deleteAccType(&id)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Account type deleted successfully. Total rows/record affected %v", deletedRows)

	c.JSON(http.StatusOK, gin.H{"rows": deletedRows, "message": msg})
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
