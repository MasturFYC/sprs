package controller

import (
	"database/sql"
	"fmt"
	"log"

	"fyc.com/sprs/models"
	"github.com/gin-gonic/gin"

	"net/http"

	"strconv"
)

// func GetOfficeAddresses(c *gin.Context) {
//

// 	addresses, err := getAllOfficeAddresses()

// 	if err != nil {
// 		log.Fatalf("Unable to get all office addresses. %v", err)
// 	}

// 	c.JSON(http.StatusOK, &addresses)
// }

func GetOfficeAddress(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//		log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ha, err := getOfficeAddress(&id)

	if err != nil {
		//		log.Fatalf("Unable to get office address. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &ha)
}

func DeleteOfficeAddress(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//		log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedRows := deleteOfficeAddress(&id)

	msg := fmt.Sprintf("Office address deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      id,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func CreateOfficeAddress(c *gin.Context) {

	var ha models.OfficeAddress

	err := c.BindJSON(&ha)

	if err != nil {
		//		log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rowAffected, err := createOfficeAddress(&ha)

	if err != nil {
		//		log.Fatalf("Nama office address tidak boleh sama.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Office address created successfully. Total rows/record affected %v", rowAffected)

	// format the reponse message
	res := Response{
		ID:      rowAffected,
		Message: msg,
	}

	c.JSON(http.StatusOK, &res)

}

func UpdateOfficeAddress(c *gin.Context) {

	// create the officegres db connection

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var ha models.OfficeAddress

	err := c.BindJSON(&ha)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRows := updateOfficeAddress(&id, &ha)

	msg := fmt.Sprintf("Office address updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getOfficeAddress(id *int64) (models.OfficeAddress, error) {

	var ha models.OfficeAddress

	var sqlStatement = `SELECT 
		order_id, street, region, city, phone, zip
	FROM office_addresses
	WHERE order_id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&ha.OrderID, &ha.Street, &ha.Region, &ha.City, &ha.Phone, &ha.Zip)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return ha, nil
	case nil:
		return ha, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return ha, err
}

// func getAllOfficeAddresses() ([]models.OfficeAddress, error) {

// 	var addresses []models.OfficeAddress

// 	var sqlStatement = `SELECT
// 		order_id, street, region, city, phone, zip
// 	FROM office_addresses
// 	ORDER BY name`

// 	rs, err := Sql().Query(sqlStatement)

// 	if err != nil {
// 		log.Fatalf("Unable to execute office addresses query %v", err)
// 	}

// 	defer rs.Close()

// 	for rs.Next() {
// 		var ha models.OfficeAddress

// 		err := rs.Scan(&ha.OrderID, &ha.Street, &ha.Region, &ha.City, &ha.Phone, &ha.Zip)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		addresses = append(addresses, ha)
// 	}

// 	return addresses, err
// }

func deleteOfficeAddress(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM office_addresses WHERE order_id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete office address. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createOfficeAddress(ha *models.OfficeAddress) (int64, error) {

	sqlStatement := `INSERT INTO office_addresses
	(order_id, street, region, city, phone, zip) 
	VALUES 
	($1, $2, $3, $4, $5, $6)`

	res, err := Sql().Exec(sqlStatement,
		ha.OrderID,
		ha.Street,
		ha.Region,
		ha.City,
		ha.Phone,
		ha.Zip,
	)

	if err != nil {
		log.Fatalf("Unable to create office address. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Unable to create office address. %v", err)
	}

	return rowsAffected, err
}

func updateOfficeAddress(id *int64, ha *models.OfficeAddress) int64 {

	sqlStatement := `UPDATE office_addresses SET
		street=$2, region=$3, city=$4, phone=$5, zip=$6
	WHERE order_id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		ha.Street,
		ha.Region,
		ha.City,
		ha.Phone,
		ha.Zip,
	)

	if err != nil {
		log.Fatalf("Unable to update office address. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating office address. %v", err)
	}

	return rowsAffected
}
