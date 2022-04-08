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

func FinanceGetAll(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	finances, err := getAllFinances(db)

	if err != nil {
		//		log.Fatalf("Unable to get all finances. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &finances)
}

func FinanceGetItem(c *gin.Context) {

	// id = order id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	finances, err := getFinance(db, &id)

	if err != nil {
		//log.Fatalf("Unable to get finance. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &finances)
}

func FinanceDelete(c *gin.Context) {

	// id = order id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	deletedRows, err := deleteFinance(db, &id)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Finance deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func FinanceCreate(c *gin.Context) {

	var finance models.Finance

	err := c.BindJSON(&finance)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	id, err := createFinance(db, &finance)

	if err != nil {
		//log.Fatalf("Nama finance tidak boleh sama.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	finance.ID = id

	c.JSON(http.StatusOK, &finance)

}

func FinanceUpdate(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var finance models.Finance

	err := c.BindJSON(&finance)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateFinance(db, &id, &finance)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Finance updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getFinance(db *sql.DB, id *int) (models.Finance, error) {

	var p models.Finance

	var sqlStatement = `SELECT 
		id, name, short_name, street, city, phone, cell, zip, email, group_id
	FROM finances
	WHERE id=$1`

	rs := db.QueryRow(sqlStatement, id)

	err := rs.Scan(
		&p.ID,
		&p.Name,
		&p.ShortName,
		&p.Street,
		&p.City,
		&p.Phone,
		&p.Cell,
		&p.Zip,
		&p.Email,
		&p.GroupID,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return p, nil
	case nil:
		return p, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return p, err
}

func getAllFinances(db *sql.DB) ([]models.Finance, error) {

	var finances []models.Finance

	var sqlStatement = `SELECT 
		id, name, short_name, street, city, phone, cell, zip, email, group_id
	FROM finances
	ORDER BY name`

	rs, err := db.Query(sqlStatement)

	if err != nil {
		// log.Fatalf("Unable to execute finances query %v", err)
		return finances, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.Finance

		err := rs.Scan(
			&p.ID,
			&p.Name,
			&p.ShortName,
			&p.Street,
			&p.City,
			&p.Phone,
			&p.Cell,
			&p.Zip,
			&p.Email,
			&p.GroupID)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		finances = append(finances, p)
	}

	return finances, err
}

func deleteFinance(db *sql.DB, id *int) (int64, error) {
	// create the delete sql query
	sqlStatement := `DELETE FROM finances WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		//log.Fatalf("Unable to delete finance. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Fatalf("Error while checking the affected rows. %v", err)
	// }

	return rowsAffected, err
}

func createFinance(db *sql.DB, finance *models.Finance) (int, error) {

	sqlStatement := `INSERT INTO finances 
	(name, short_name, street, city, phone, cell, zip, email, group_id) 
	VALUES 
	($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id`

	var id int

	err := db.QueryRow(sqlStatement,
		finance.Name,
		finance.ShortName,
		finance.Street,
		finance.City,
		finance.Phone,
		finance.Cell,
		finance.Zip,
		finance.Email,
		finance.GroupID,
	).Scan(&id)

	// if err != nil {
	// 	log.Printf("Unable to create finance. %v", err)
	// }

	return id, err
}

func updateFinance(db *sql.DB, id *int, finance *models.Finance) (int64, error) {

	sqlStatement := `UPDATE finances SET
		name=$2, short_name=$3, street=$4, city=$5, phone=$6, cell=$7, zip=$8, email=$9, group_id=$10
	WHERE id=$1`

	res, err := db.Exec(sqlStatement,
		id,
		finance.Name,
		finance.ShortName,
		finance.Street,
		finance.City,
		finance.Phone,
		finance.Cell,
		finance.Zip,
		finance.Email,
		finance.GroupID,
	)

	if err != nil {
		//log.Printf("Unable to update finance. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Printf("Error while updating finance. %v", err)
	// }

	return rowsAffected, err
}
