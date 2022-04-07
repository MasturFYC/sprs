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

// func GetPostAddresses(c *gin.Context) {
//

// 	addresses, err := getAllPostAddresses()

// 	if err != nil {
// 		log.Fatalf("Unable to get all post addresses. %v", err)
// 	}

// 	c.JSON(http.StatusOK, &addresses)
// }

func GetPostAddress(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ha, err := getPostAddress(&id)

	if err != nil {
		//log.Fatalf("Unable to get post address. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &ha)
}

func DeletePostAddress(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedRows := deletePostAddress(&id)

	msg := fmt.Sprintf("Post address deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      id,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func CreatePostAddress(c *gin.Context) {

	var ha models.PostAddress

	err := c.BindJSON(&ha)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rowAffected, err := createPostAddress(&ha)

	if err != nil {
		//log.Fatalf("Nama post address tidak boleh sama.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Post address created successfully. Total rows/record affected %v", rowAffected)

	// format the reponse message
	res := Response{
		ID:      rowAffected,
		Message: msg,
	}

	c.JSON(http.StatusOK, &res)

}

func UpdatePostAddress(c *gin.Context) {

	// create the postgres db connection

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var ha models.PostAddress

	err := c.BindJSON(&ha)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	updatedRows := updatePostAddress(&id, &ha)

	msg := fmt.Sprintf("Post address updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getPostAddress(id *int64) (models.PostAddress, error) {

	var ha models.PostAddress

	var sqlStatement = `SELECT 
		order_id, street, region, city, phone, zip
	FROM post_addresses
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

// func getAllPostAddresses() ([]models.PostAddress, error) {

// 	var addresses []models.PostAddress

// 	var sqlStatement = `SELECT
// 		order_id, street, region, city, phone, zip
// 	FROM post_addresses
// 	ORDER BY name`

// 	rs, err := Sql().Query(sqlStatement)

// 	if err != nil {
// 		log.Fatalf("Unable to execute post addresses query %v", err)
// 	}

// 	defer rs.Close()

// 	for rs.Next() {
// 		var ha models.PostAddress

// 		err := rs.Scan(&ha.OrderID, &ha.Street, &ha.Region, &ha.City, &ha.Phone, &ha.Zip)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		addresses = append(addresses, ha)
// 	}

// 	return addresses, err
// }

func deletePostAddress(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM post_addresses WHERE order_id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete post address. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createPostAddress(ha *models.PostAddress) (int64, error) {

	sqlStatement := `INSERT INTO post_addresses
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
		log.Fatalf("Unable to create post address. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Unable to create post address. %v", err)
	}

	return rowsAffected, err
}

func updatePostAddress(id *int64, ha *models.PostAddress) int64 {

	sqlStatement := `UPDATE post_addresses SET
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
		log.Fatalf("Unable to update post address. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating post address. %v", err)
	}

	return rowsAffected
}
