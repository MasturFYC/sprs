package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

// func GetReceivables(w http.ResponseWriter, r *http.Request) {
// 	EnableCors(&w)

// 	addresses, err := getAllReceivables()

// 	if err != nil {
// 		log.Fatalf("Unable to get all receivablees. %v", err)
// 	}

// 	json.NewEncoder(w).Encode(&addresses)
// }

func GetReceivable(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	rv, err := getReceivable(&id)

	if err != nil {
		log.Fatalf("Unable to get receivable. %v", err)
	}

	json.NewEncoder(w).Encode(&rv)
}

func DeleteReceivable(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteReceivable(&id)

	msg := fmt.Sprintf("Receivale deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      id,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreateReceivable(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var rv models.Receivable

	err := json.NewDecoder(r.Body).Decode(&rv)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	rowAffected, err := createReceivable(&rv)

	if err != nil {
		log.Fatalf("Nama receivable tidak boleh sama.  %v", err)
	}

	msg := fmt.Sprintf("Receivale created successfully. Total rows/record affected %v", rowAffected)

	// format the reponse message
	res := Response{
		ID:      rowAffected,
		Message: msg,
	}

	json.NewEncoder(w).Encode(&res)

}

func UpdateReceivable(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var rv models.Receivable

	err := json.NewDecoder(r.Body).Decode(&rv)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateReceivable(&id, &rv)

	msg := fmt.Sprintf("Receivale updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getReceivable(id *int64) (models.Receivable, error) {

	var rv models.Receivable

	var sqlStatement = `SELECT 
		order_id, covenant_at, due_at, mortgage_by_month, mortgage_receivable, running_fine,
		rest_fine, bill_service, pay_deposit, rest_receivable, rest_base, day_period,
		mortgage_to, day_count
	FROM receivables
	WHERE order_id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(
		&rv.OrderID,
		&rv.CovenantAt,
		&rv.DueAt,
		&rv.MortgageByMonth,
		&rv.MortgageReceivable,
		&rv.RunningFine,
		&rv.RestFine,
		&rv.BillService,
		&rv.PayDeposit,
		&rv.RestReceivable,
		&rv.RestBase,
		&rv.DayPeriod,
		&rv.MortgageTo,
		&rv.DayCount,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return rv, nil
	case nil:
		return rv, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return rv, err
}

// func getAllReceivables() ([]models.Receivable, error) {

// 	var addresses []models.Receivable

// 	var sqlStatement = `SELECT
// 		order_id, street, region, city, phone, zip
// 	FROM receivables
// 	ORDER BY name`

// 	rs, err := Sql().Query(sqlStatement)

// 	if err != nil {
// 		log.Fatalf("Unable to execute receivablees query %v", err)
// 	}

// 	defer rs.Close()

// 	for rs.Next() {
// 		var ha models.Receivable

// 		err := rs.Scan(&ha.OrderID, &ha.Street, &ha.Region, &ha.City, &ha.Phone, &ha.Zip)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		addresses = append(addresses, ha)
// 	}

// 	return addresses, err
// }

func deleteReceivable(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM receivables WHERE order_id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete receivable. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createReceivable(rv *models.Receivable) (int64, error) {

	sqlStatement := `INSERT INTO receivables (
		order_id, covenant_at, due_at, mortgage_by_month, mortgage_receivable, running_fine,
		rest_fine, bill_service, pay_deposit, rest_receivable, rest_base, day_period,
		mortgage_to, day_count
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	res, err := Sql().Exec(sqlStatement,
		rv.OrderID,
		rv.CovenantAt,
		rv.DueAt,
		rv.MortgageByMonth,
		rv.MortgageReceivable,
		rv.RunningFine,
		rv.RestFine,
		rv.BillService,
		rv.PayDeposit,
		rv.RestReceivable,
		rv.RestBase,
		rv.DayPeriod,
		rv.MortgageTo,
		rv.DayCount,
	)

	if err != nil {
		log.Fatalf("Unable to create receivable. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Unable to create receivable. %v", err)
	}

	return rowsAffected, err
}

func updateReceivable(id *int64, rv *models.Receivable) int64 {

	sqlStatement := `UPDATE receivables SET
		covenant_at=$2,
		due_at=$3,
		mortgage_by_month=$4,
		mortgage_receivable=$5,
		running_fine=$6,
		rest_fine=$7,
		bill_service=$8,
		pay_deposit=$9,
		rest_receivable=$10,
		rest_base=$11,
		day_period=$12,
		mortgage_to=$13,
		day_count=$14
	WHERE order_id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		rv.CovenantAt,
		rv.DueAt,
		rv.MortgageByMonth,
		rv.MortgageReceivable,
		rv.RunningFine,
		rv.RestFine,
		rv.BillService,
		rv.PayDeposit,
		rv.RestReceivable,
		rv.RestBase,
		rv.DayPeriod,
		rv.MortgageTo,
		rv.DayCount,
	)

	if err != nil {
		log.Fatalf("Unable to update receivable. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating receivable. %v", err)
	}

	return rowsAffected
}
