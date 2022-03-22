package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func SearchOrders(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	var t models.SearchGroup

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_codes, err := searchOrders(&t.Txt)

	if err != nil || len(acc_codes) == 0 {
		//log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_codes)
}

func GetOrdersByFinance(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	orders, err := get_order_by_finance(&id)

	if err != nil || len(orders) == 0 {
		//log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
		//var test []models.AccCode
		//json.NewEncoder(w).Encode(test)
		//return
	}

	json.NewEncoder(w).Encode(&orders)
}

func GetOrdersByBranch(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_codes, err := get_order_by_branch(&id)

	if err != nil || len(acc_codes) == 0 {
		//log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
		//var test []models.AccCode
		//json.NewEncoder(w).Encode(test)
		//return
	}

	json.NewEncoder(w).Encode(&acc_codes)
}

func GetOrdersByMonth(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_codes, err := get_order_by_month(&id)

	if err != nil || len(acc_codes) == 0 {
		//log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
		//var test []models.AccCode
		//json.NewEncoder(w).Encode(test)
		//return
	}

	json.NewEncoder(w).Encode(&acc_codes)
}

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

func Order_GetNameSeq(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	id, err := create_name_seq()

	if err != nil {
		//log.Printf("Nama order tidak boleh sama.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	res := Response{
		ID:      id,
		Message: "Nama urut baru order",
	}

	json.NewEncoder(w).Encode(&res)

}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var order models.Order

	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		//log.Printf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusRequestedRangeNotSatisfiable), http.StatusRequestedRangeNotSatisfiable)
		return
	}

	id, err := createOrder(&order)

	if err != nil {
		//log.Printf("Nomor order tidak boleh sama.  %v", err)
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
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
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
		id, name, order_at, printed_at, bt_finance, bt_percent, bt_matel, 
		user_name, verified_by, finance_id, branch_id, is_stnk, stnk_price, matrix
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
		// &o.Ppn,
		// &o.Nominal,
		// &o.Subtotal,
		&o.UserName,
		&o.VerifiedBy,
		// &o.ValidatedBy,
		&o.FinanceID,
		&o.BranchID,
		&o.IsStnk,
		&o.StnkPrice,
		&o.Matrix,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return o, nil
	case nil:

		set_child(&o)

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
		id, name, order_at, printed_at, bt_finance, bt_percent, bt_matel, 
		user_name, verified_by, finance_id, branch_id, is_stnk, stnk_price, matrix
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
			// &o.Ppn,
			// &o.Nominal,
			// &o.Subtotal,
			&o.UserName,
			&o.VerifiedBy,
			//&o.ValidatedBy,
			&o.FinanceID,
			&o.BranchID,
			&o.IsStnk,
			&o.StnkPrice,
			&o.Matrix,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		set_child(&o)

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

func create_name_seq() (int64, error) {

	sqlStatement := "SELECT nextval('order_name_seq'::regclass) AS id"

	var id int64

	err := Sql().QueryRow(sqlStatement).Scan(&id)

	if err != nil {
		log.Printf("Unable to create order name sequence. %v", err)
	}

	return id, err
}

func createOrder(p *models.Order) (int64, error) {

	sqlStatement := `INSERT INTO orders (
		name, order_at, printed_at, bt_finance, bt_percent, bt_matel,
		user_name, verified_by, finance_id, branch_id,
		is_stnk, stnk_price, matrix, token
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,
		to_tsvector('indonesian', $14))
	RETURNING id`

	var id int64

	token := strings.Join(create_token(p), " ")

	err := Sql().QueryRow(sqlStatement,
		p.Name,
		p.OrderAt,
		p.PrintedAt,
		p.BtFinance,
		p.BtPercent,
		p.BtMatel,
		// p.Ppn,
		// p.Nominal,
		//	p.Subtotal,
		p.UserName,
		p.VerifiedBy,
		//p.ValidatedBy,
		p.FinanceID,
		p.BranchID,
		p.IsStnk,
		p.StnkPrice,
		p.Matrix,
		token,
	).Scan(&id)

	if err != nil {
		log.Printf("Unable to create order. %v", err)
	}

	return id, err
}

func create_token(p *models.Order) []string {
	var s []string

	s = append(s, p.Name,
		p.OrderAt,
		p.Finance.Name,
		p.Finance.ShortName,
		p.Branch.Name,
		p.Branch.HeadBranch,
	)
	if p.IsStnk {
		s = append(s, "stnk-ada")
	} else {
		s = append(s, "stnk-tidak-ada")
	}

	if p.Unit.TypeID > 0 {

		s = append(s, p.Unit.Nopol, p.Unit.Type.Name)

		if p.Unit.WarehouseID > 0 {
			s = append(s, p.Unit.Warehouse.Name)
		}

		if p.Unit.FrameNumber != "" {
			s = append(s, string(p.Unit.FrameNumber))
		}

		if p.Unit.MachineNumber != "" {
			s = append(s, string(p.Unit.MachineNumber))
		}
		if p.Unit.Color != "" {
			s = append(s, string(p.Unit.Color))
		}

		if p.Unit.Year != 0 {
			s = append(s, strconv.FormatInt(p.Unit.Year, 10))
		}

		// if p.Unit.Dealer != "" {
		// 	s = append(s, p.Unit.Dealer)
		// }

		// if p.Unit.Surveyor != "" {
		// 	s = append(s, p.Unit.Surveyor)
		// }

		// if p.Unit.BpkbName != "" {
		// 	s = append(s, p.Unit.BpkbName)
		// }

		if p.Unit.Type.MerkID > 0 {
			s = append(s, p.Unit.Type.Merk.Name)
		}
		if p.Unit.Type.WheelID > 0 {
			s = append(s, p.Unit.Type.Wheel.Name)
			s = append(s, p.Unit.Type.Wheel.ShortName)
		}
	}

	return s

}

func updateOrder(id *int64, p *models.Order) (int64, error) {

	sqlStatement := `UPDATE orders SET
		name=$2, order_at=$3, printed_at=$4, bt_finance=$5, bt_percent=$6, bt_matel=$7, 
		user_name=$8, verified_by=$9, finance_id=$10, branch_id=$11,
		is_stnk=$12, stnk_price=$13, matrix=$14, token=to_tsvector('indonesian', $15)
	WHERE id=$1`

	token := strings.Join(create_token(p), " ")

	res, err := Sql().Exec(sqlStatement,
		id,
		p.Name,
		p.OrderAt,
		p.PrintedAt,
		p.BtFinance,
		p.BtPercent,
		p.BtMatel,
		// p.Ppn,
		// p.Nominal,
		// p.Subtotal,
		p.UserName,
		p.VerifiedBy,
		//p.ValidatedBy,
		p.FinanceID,
		p.BranchID,
		p.IsStnk,
		p.StnkPrice,
		p.Matrix,
		token,
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

func searchOrders(txt *string) ([]models.Order, error) {

	var orders []models.Order

	var sqlStatement = `SELECT
		id, name, order_at, printed_at, bt_finance, bt_percent, bt_matel,
		user_name, verified_by, finance_id, branch_id,
		is_stnk, stnk_price, matrix
	FROM orders
	WHERE token @@ to_tsquery('indonesian', $1)
	ORDER BY id DESC`

	rs, err := Sql().Query(sqlStatement, txt)

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
			// &o.Ppn,
			// &o.Nominal,
			// &o.Subtotal,
			&o.UserName,
			&o.VerifiedBy,
			//&o.ValidatedBy,
			&o.FinanceID,
			&o.BranchID,
			&o.IsStnk,
			&o.StnkPrice,
			&o.Matrix,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		set_child(&o)

		orders = append(orders, o)
	}

	return orders, err
}

func get_order_by_finance(id *int) ([]models.Order, error) {

	var orders []models.Order

	var sqlStatement = `SELECT
		id, name, order_at, printed_at, bt_finance, bt_percent, bt_matel, 
		user_name, verified_by, finance_id, branch_id,
		is_stnk, stnk_price, matrix
	FROM orders
	WHERE finance_id=$1
	ORDER BY finance_id, id DESC`

	rs, err := Sql().Query(sqlStatement, id)

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
			// &o.Ppn,
			// &o.Nominal,
			// &o.Subtotal,
			&o.UserName,
			&o.VerifiedBy,
			//&o.ValidatedBy,
			&o.FinanceID,
			&o.BranchID,
			&o.IsStnk,
			&o.StnkPrice,
			&o.Matrix,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		set_child(&o)

		orders = append(orders, o)
	}

	return orders, err
}

func get_order_by_branch(id *int) ([]models.Order, error) {

	var orders []models.Order

	var sqlStatement = `SELECT
		id, name, order_at, printed_at, bt_finance, bt_percent, bt_matel,
		user_name, verified_by, finance_id, branch_id,
		is_stnk, stnk_price, matrix
	FROM orders
	WHERE branch_id=$1
	ORDER BY branch_id, id DESC`

	rs, err := Sql().Query(sqlStatement, id)

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
			// &o.Ppn,
			// &o.Nominal,
			// &o.Subtotal,
			&o.UserName,
			&o.VerifiedBy,
			//&o.ValidatedBy,
			&o.FinanceID,
			&o.BranchID,
			&o.IsStnk,
			&o.StnkPrice,
			&o.Matrix,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		set_child(&o)

		orders = append(orders, o)
	}

	return orders, err
}

func set_child(o *models.Order) {
	finance, _ := getFinance(&o.FinanceID)
	o.Finance = finance

	branch, _ := getBranch(&o.BranchID)
	o.Branch = branch

	cust, _ := getCustomer(&o.ID)
	o.Customer = cust

	// receivable, _ := getReceivable(&o.ID)
	// o.Receivable = receivable

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
}

func get_order_by_month(id *int) ([]models.Order, error) {

	var orders []models.Order

	var sqlStatement = `SELECT
		id, name, order_at, printed_at, bt_finance, bt_percent, bt_matel, 
		user_name, verified_by, finance_id, branch_id,
		is_stnk, stnk_price, matrix
	FROM orders
	WHERE EXTRACT(MONTH from order_at)=$1
	ORDER BY branch_id, id DESC`

	rs, err := Sql().Query(sqlStatement, id)

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
			// &o.Ppn,
			// &o.Nominal,
			// &o.Subtotal,
			&o.UserName,
			&o.VerifiedBy,
			//&o.ValidatedBy,
			&o.FinanceID,
			&o.BranchID,
			&o.IsStnk,
			&o.StnkPrice,
			&o.Matrix,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		set_child(&o)

		orders = append(orders, o)
	}

	return orders, err
}
