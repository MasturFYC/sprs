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

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	customers, err := getAllCustomer()

	if err != nil {
		//		log.Fatalf("Unable to get all customers. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	customers, err := getCustomer(&id)

	if err != nil {
		//log.Fatalf("Unable to get category. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&customers)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows, err := deleteCustomer(&id)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Customer deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      id,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var cust models.Customer

	err := json.NewDecoder(r.Body).Decode(&cust)

	if err != nil {
		//log.Printf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusCreated), http.StatusCreated)
		return
	}

	_, err = createCustomer(&cust)

	if err != nil {
		//log.Printf("Nama customers tidak boleh sama.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	json.NewEncoder(w).Encode(&cust)

}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var cust models.Customer

	err := json.NewDecoder(r.Body).Decode(&cust)

	if err != nil {
		//log.Printf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusCreated), http.StatusCreated)
		return
	}

	_, err = updateCustomer(&id, &cust)

	if err != nil {
		//log.Printf("Unable to update customer.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	json.NewEncoder(w).Encode(cust)
}

func getCustomer(id *int64) (models.Customer, error) {

	var cust models.Customer

	var sqlStatement = `SELECT 
		order_id, name, agreement_number, payment_type
	FROM customers
	WHERE order_id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&cust.OrderID, &cust.Name, &cust.AgreementNumber, &cust.PaymentType)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return cust, err
	case nil:
		return cust, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return cust, err
}

func getAllCustomer() ([]models.Customer, error) {

	var customers []models.Customer

	var sqlStatement = `SELECT 
		order_id, name, agreement_number, payment_type
	FROM customers
	ORDER BY name`

	rs, err := Sql().Query(sqlStatement)

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

func deleteCustomer(id *int64) (int64, error) {
	// create the delete sql query
	sqlStatement := `DELETE FROM customers WHERE order_id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		//log.Fatalf("Unable to delete customer. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Fatalf("Error while checking the affected rows. %v", err)
	// }

	return rowsAffected, err
}

func createCustomer(cust *models.Customer) (int64, error) {

	sqlStatement := `INSERT INTO customers 
	(order_id, name, agreement_number, payment_type) 
	VALUES 
	($1, $2, $3, $4)`

	res, err := Sql().Exec(sqlStatement,
		cust.OrderID,
		cust.Name,
		cust.AgreementNumber,
		cust.PaymentType,
	)

	if err != nil {
		//log.Printf("Unable to create customer. %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	// if err != nil || rowsAffected == 0 {
	// 	log.Printf("Unable to create customer. %v", err)
	// }

	return rowsAffected, err
}

func updateCustomer(id *int64, cust *models.Customer) (int64, error) {

	sqlStatement := `UPDATE customers SET
		name=$2, agreement_number=$3, payment_type=$4
	WHERE order_id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		cust.Name,
		cust.AgreementNumber,
		cust.PaymentType,
	)

	if err != nil {
		//log.Printf("Unable to update customer. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil || rowsAffected == 0 {
	// 	log.Printf("Error while updating customer. %v", err)
	// }

	return rowsAffected, err
}
