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

func GetCustomers(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	customers, err := getAllCustomer(db)

	if err != nil {
		//		log.Fatalf("Unable to get all customers. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &customers)
}

func CustomerGetItem(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	customers, err := getCustomer(db, &id)

	if err != nil {
		//log.Fatalf("Unable to get category. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &customers)
}

func CustomerDelete(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	deletedRows, err := deleteCustomer(db, &id)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)

		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Customer deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      id,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func CustomerCreate(c *gin.Context) {

	var cust models.Customer

	err := c.BindJSON(&cust)

	if err != nil {
		//log.Printf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("%v", cust)

	db := c.Keys["db"].(*sql.DB)
	_, err = createCustomer(db, &cust)

	if err != nil {
		//log.Printf("Nama customers tidak boleh sama.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &cust)

}

func CustomerUpdate(c *gin.Context) {
	// create the postgres db connection

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var cust models.Customer

	err := c.BindJSON(&cust)

	if err != nil {
		//log.Printf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	_, err = updateCustomer(db, &id, &cust)

	if err != nil {
		//log.Printf("Unable to update customer.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cust)
}

func getCustomer(db *sql.DB, id *int64) (models.Customer, error) {

	var cust = models.Customer{}

	var sqlStatement = `SELECT 
		order_id, name, agreement_number, payment_type
	FROM customers
	WHERE order_id=$1`

	rs := db.QueryRow(sqlStatement, id)

	err := rs.Scan(&cust.OrderID, &cust.Name, &cust.AgreementNumber, &cust.PaymentType)

	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		return cust, err
	case nil:
		return cust, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return cust, err
}

func getAllCustomer(db *sql.DB) ([]models.Customer, error) {

	var customers = make([]models.Customer, 0)

	var sqlStatement = `SELECT 
		order_id, name, agreement_number, payment_type
	FROM customers
	ORDER BY name`

	rs, err := db.Query(sqlStatement)

	if err != nil {
		//log.Fatalf("Unable to execute customers query %v", err)
		return customers, err
	}

	defer rs.Close()

	for rs.Next() {
		var cust models.Customer

		err := rs.Scan(&cust.OrderID, &cust.Name, &cust.AgreementNumber, &cust.PaymentType)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		customers = append(customers, cust)
	}

	return customers, err
}

func deleteCustomer(db *sql.DB, id *int64) (int64, error) {
	sqlStatement := `DELETE FROM customers WHERE order_id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}

func createCustomer(db *sql.DB, cust *models.Customer) (int64, error) {

	sqlStatement := `INSERT INTO customers 
	(order_id, name, agreement_number, payment_type) 
	VALUES 
	($1, $2, $3, $4)`

	res, err := db.Exec(sqlStatement,
		cust.OrderID,
		cust.Name,
		cust.AgreementNumber,
		cust.PaymentType,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func updateCustomer(db *sql.DB, id *int64, cust *models.Customer) (int64, error) {

	sqlStatement := `UPDATE customers SET
		name=$2, agreement_number=$3, payment_type=$4
	WHERE order_id=$1`

	res, err := db.Exec(sqlStatement,
		id,
		cust.Name,
		cust.AgreementNumber,
		cust.PaymentType,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}
