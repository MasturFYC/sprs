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

func GetOrders(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	addresses, err := getAllOrders()

	if err != nil {
		log.Fatalf("Unable to get all orderes. %v", err)
	}

	json.NewEncoder(w).Encode(&addresses)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	rv, err := getOrder(&id)

	if err != nil {
		log.Printf("Unable to get order. %v", err)
	}

	json.NewEncoder(w).Encode(&rv)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	log.Printf("id to remove.  %v", id)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteOrder(&id)

	msg := fmt.Sprintf("Order deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      id,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var order models.Order

	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusRequestedRangeNotSatisfiable), http.StatusRequestedRangeNotSatisfiable)
		return
	}

	id, err := createOrder(&order)

	if err != nil {
		log.Printf("Nomor order tidak boleh sama.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	order.ID = id

	json.NewEncoder(w).Encode(&order)

}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var rv models.Order

	err := json.NewDecoder(r.Body).Decode(&rv)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows, err := updateOrder(&id, &rv)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Order updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getOrder(id *int64) (models.Order, error) {

	var o models.Order

	var sqlStatement = `SELECT 
		id, name, order_at, printed_at, bt_finance, bt_percent, bt_matel, ppn,
		nominal, subtotal, user_name, verified_by, validated_by, finance_id, branch_id,
		is_stnk, stnk_price
	FROM orders
	WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(
		&o.ID,
		&o.Name,
		&o.OrderAt,
		&o.PrintedAt,
		&o.BtFinance,
		&o.BtPercent,
		&o.BtMatel,
		&o.Ppn,
		&o.Nominal,
		&o.Subtotal,
		&o.UserName,
		&o.VerifiedBy,
		&o.ValidatedBy,
		&o.FinanceID,
		&o.BranchID,
		&o.IsStnk,
		&o.StnkPrice,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return o, nil
	case nil:

		finance, _ := getFinance(&o.FinanceID)
		o.Finance = finance
		branch, _ := getBranch(&o.BranchID)
		o.Branch = branch
		cust, _ := getCustomer(&o.ID)
		o.Customer = cust
		receivable, _ := getReceivable(&o.ID)
		o.Receivable = receivable
		unit, _ := getUnit(&o.ID)
		o.Unit = unit
		actions, _ := getAllActions(&o.ID)
		o.Actions = actions
		task, _ := getTask(&o.ID)
		o.Task = task
		home, _ := getHomeAddress(&o.ID)
		o.HomeAddress = home
		office, _ := getOfficeAddress(&o.ID)
		o.OfficeAddress = office
		post, _ := getPostAddress(&o.ID)
		o.PostAddress = post
		ktp, _ := getKTPAddress(&o.ID)
		o.KtpAddress = ktp

		return o, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return o, err
}

func getAllOrders() ([]models.Order, error) {

	var orders []models.Order

	var sqlStatement = `SELECT
		id, name, order_at, printed_at, bt_finance, bt_percent, bt_matel, ppn,
		nominal, subtotal, user_name, verified_by, validated_by, finance_id, branch_id,
		is_stnk, stnk_price
	FROM orders
	ORDER BY id DESC`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var o models.Order

		err := rs.Scan(
			&o.ID,
			&o.Name,
			&o.OrderAt,
			&o.PrintedAt,
			&o.BtFinance,
			&o.BtPercent,
			&o.BtMatel,
			&o.Ppn,
			&o.Nominal,
			&o.Subtotal,
			&o.UserName,
			&o.VerifiedBy,
			&o.ValidatedBy,
			&o.FinanceID,
			&o.BranchID,
			&o.IsStnk,
			&o.StnkPrice,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		finance, _ := getFinance(&o.FinanceID)
		o.Finance = finance

		branch, _ := getBranch(&o.BranchID)
		o.Branch = branch

		cust, err := getCustomer(&o.ID)
		o.Customer = cust

		receivable, _ := getReceivable(&o.ID)
		o.Receivable = receivable

		unit, _ := getUnit(&o.ID)
		o.Unit = unit

		actions, _ := getAllActions(&o.ID)
		o.Actions = actions

		task, _ := getTask(&o.ID)
		o.Task = task

		home, _ := getHomeAddress(&o.ID)
		o.HomeAddress = home

		office, _ := getOfficeAddress(&o.ID)
		o.OfficeAddress = office

		post, _ := getPostAddress(&o.ID)
		o.PostAddress = post

		ktp, _ := getKTPAddress(&o.ID)
		o.KtpAddress = ktp

		orders = append(orders, o)
	}

	return orders, err
}

func deleteOrder(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM orders WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete order. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createOrder(o *models.Order) (int64, error) {

	sqlStatement := `INSERT INTO orders (
		name, order_at, printed_at, bt_finance, bt_percent, bt_matel, ppn,
		nominal, subtotal, user_name, verified_by, validated_by, finance_id, branch_id,
		is_stnk, stnk_price
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	RETURNING id`

	var id int64

	err := Sql().QueryRow(sqlStatement,
		o.Name,
		o.OrderAt,
		o.PrintedAt,
		o.BtFinance,
		o.BtPercent,
		o.BtMatel,
		o.Ppn,
		o.Nominal,
		o.Subtotal,
		o.UserName,
		o.VerifiedBy,
		o.ValidatedBy,
		o.FinanceID,
		o.BranchID,
		o.IsStnk,
		o.StnkPrice,
	).Scan(&id)

	if err != nil {
		log.Printf("Unable to create order. %v", err)
	}

	return id, err
}

func updateOrder(id *int64, o *models.Order) (int64, error) {

	sqlStatement := `UPDATE orders SET
		name=$2, order_at=$3, printed_at=$4, bt_finance=$5, bt_percent=$6, bt_matel=$7, ppn=$8,
		nominal=$9, subtotal=$10, user_name=$11, verified_by=$12, validated_by=$13, finance_id=$14, branch_id=$15,
		is_stnk=$16, stnk_price=$17
	WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		o.Name,
		o.OrderAt,
		o.PrintedAt,
		o.BtFinance,
		o.BtPercent,
		o.BtMatel,
		o.Ppn,
		o.Nominal,
		o.Subtotal,
		o.UserName,
		o.VerifiedBy,
		o.ValidatedBy,
		o.FinanceID,
		o.BranchID,
		o.IsStnk,
		o.StnkPrice,
	)

	if err != nil {
		log.Printf("Unable to update order. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating order. %v", err)
	}

	return rowsAffected, err
}
