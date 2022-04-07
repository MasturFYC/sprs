package controller

import (
	"fmt"
	"log"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTransactionDetails(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	details, err := getTransactionDetails(&id)

	if err != nil {
		log.Printf("Unable to get transaction detail. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, &details)
}

func CreateTransactionDetail(c *gin.Context) {

	var trx models.TrxDetail

	err := c.BindJSON(&trx)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction detail.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := createTransactionDetail(&trx)

	if err != nil {
		//log.Printf("(API) Unable to create transaction detail.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	trx.ID = id

	c.JSON(http.StatusOK, &trx)

}

func UpdateTransactionDetail(c *gin.Context) {

	// create the postgres db connection

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var detail models.TrxDetail

	err := c.BindJSON(&detail)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction detail.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRows, err := updateTransactionDetail(&id, &detail)

	if err != nil {
		//log.Printf("Unable to update transaction detail.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("TransactionDetail type updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func DeleteTransactionDetail(c *gin.Context) {

	trxid, err := strconv.ParseInt(c.Param("trxid"), 10, 64)

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedRows, err := deleteTransactionDetail(&trxid, &id)

	if err != nil {
		//log.Printf("Unable to delete transaction detail.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("TransactionDetail type deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      deletedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getTransactionDetails(trxID *int64) ([]models.TrxDetail, error) {

	var details []models.TrxDetail

	var sqlStatement = `SELECT 
		id, code_id, trx_id, debt, cred
	FROM trx_detail
	WHERE trx_id=$1
	ORDER BY id`

	rs, err := Sql().Query(sqlStatement, trxID)

	if err != nil {
		//log.Printf("Unable to execute transaction details query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.TrxDetail

		err := rs.Scan(
			&p.ID,
			&p.CodeID,
			&p.TrxID,
			&p.Debt,
			&p.Cred,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		details = append(details, p)
	}

	return details, err
}

func createTransactionDetail(p *models.TrxDetail) (int64, error) {

	sqlStatement := `INSERT INTO trx_detail
	(code_id, trx_id, debt, cred)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	var id int64

	err := Sql().QueryRow(sqlStatement,
		&p.CodeID,
		&p.TrxID,
		&p.Debt,
		&p.Cred).Scan(&id)

	if err != nil {
		log.Printf("Unable to create transaction detail. %v", err)
		return 0, err
	}

	return id, err
}

func updateTransactionDetail(id *int64, p *models.TrxDetail) (int64, error) {

	sqlStatement := `UPDATE trx_detail SET 
		code_id=$2,
		trx_id=$3,
		debt=$4,
		cred=$5
	WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		p.CodeID,
		p.TrxID,
		p.Debt,
		p.Cred,
	)

	if err != nil {
		log.Printf("Unable to update transaction detail. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating transaction detail. %v", err)
		return 0, err
	}

	return rowsAffected, err
}

func deleteTransactionDetail(trxid *int64, id *int) (int64, error) {

	sqlStatement := `DELETE FROM trx_detail WHERE trx_id=$1 AND id =$2`

	res, err := Sql().Exec(sqlStatement, trxid, id)

	if err != nil {
		log.Printf("Unable to delete transaction detail. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected, err
}
