package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/MasturFYC/fyc"
	"github.com/gin-gonic/gin"
)

func SearchOrders(c *gin.Context) {

	var t models.SearchGroup

	err := c.BindJSON(&t)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	orders, err := searchOrders(db, &t.Txt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &orders)
}

func GetOrdersByFinance(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	orders, err := get_order_by_finance(db, &id)

	if err != nil {
		//log.Printf("Unable to get all account codes. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
		//var test []models.AccCode
		//c.JSON(http.StatusOK, test)
		//return
	}

	c.JSON(http.StatusOK, &orders)
}

func GetOrdersByBranch(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	acc_codes, err := get_order_by_branch(db, &id)

	if err != nil {
		//log.Printf("Unable to get all account codes. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
		//var test []models.AccCode
		//c.JSON(http.StatusOK, test)
		//return
	}

	c.JSON(http.StatusOK, &acc_codes)
}

func GetOrdersByMonth(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	//log.Printf("%d============", id)

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	acc_codes, err := get_order_by_month(db, &id)

	if err != nil {
		//log.Printf("Unable to get all account codes. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
		//var test []models.AccCode
		//c.JSON(http.StatusOK, test)
		//return
	}

	c.JSON(http.StatusOK, &acc_codes)
}

func GetOrders(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)

	addresses, err := getAllOrders(db)

	if err != nil {
		//log.Fatalf("Unable to get all orderes. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &addresses)
}

func GetOrder(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	rv, err := getOrder(db, &id)

	if err != nil {
		//log.Printf("Unable to get order. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &rv)
}

func DeleteOrder(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//		log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	deletedRows := deleteOrder(db, &id)

	msg := fmt.Sprintf("Order deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      id,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func Order_GetNameSeq(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	id, err := create_name_seq(db)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	res := Response{
		ID:      id,
		Message: "Nama urut baru order",
	}

	c.JSON(http.StatusOK, &res)

}

type st_order_create struct {
	models.Order
	Finance models.Finance `json:"finance,omitempty"`
	Branch  models.Branch  `json:"branch,omitempty"`
}

func CreateOrder(c *gin.Context) {

	var order st_order_create

	err := c.BindJSON(&order)

	if err != nil {
		//log.Printf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	id, err := createOrder(db, &order)

	if err != nil {
		//log.Printf("Nomor order tidak boleh sama.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	order.ID = id

	c.JSON(http.StatusOK, &order)

}

func UpdateOrder(c *gin.Context) {

	// create the postgres db connection

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var rv st_order_update

	err := c.BindJSON(&rv)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateOrder(db, &id, &rv)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Order updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func Order_GetInvoiced(c *gin.Context) {

	m, _ := strconv.Atoi(c.Param("month"))
	y, _ := strconv.Atoi(c.Param("year"))
	fid, _ := strconv.Atoi(c.Param("financeId"))

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// invoice_id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	db := c.Keys["db"].(*sql.DB)
	orders, err := order_get_invoiced(db, &m, &y, &fid)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &orders)
}

func getOrder(db *sql.DB, id *int64) (order_all, error) {

	var o order_all
	b := create_order_query()
	b.WriteString(" FROM orders AS o")
	b.WriteString(" WHERE o.id=$1")

	rs := db.QueryRow(b.String(), id)

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
		&o.Finance,
		&o.Branch,
		&o.Unit,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return o, nil
	case nil:

		//set_child(&o)

		return o, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return o, err
}

type order_all struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	OrderAt   string  `json:"orderAt"`
	PrintedAt string  `json:"printedAt"`
	BtFinance float64 `json:"btFinance"`
	BtPercent float32 `json:"btPercent"`
	BtMatel   float64 `json:"btMatel"`
	UserName  string  `json:"userName"`

	VerifiedBy models.NullString `json:"verifiedBy"`

	FinanceID int     `json:"financeId"`
	BranchID  int     `json:"branchId"`
	IsStnk    bool    `json:"isStnk"`
	StnkPrice float64 `json:"stnkPrice"`
	Matrix    float64 `json:"matrix"`

	Finance json.RawMessage  `json:"finance,omitempty"`
	Branch  json.RawMessage  `json:"branch,omitempty"`
	Unit    *json.RawMessage `json:"unit,omitempty"`
}

func getAllOrders(db *sql.DB) ([]order_all, error) {

	var orders = make([]order_all, 0)

	b := create_order_query()
	b.WriteString(" FROM orders AS o")
	b.WriteString(" ORDER BY o.id DESC")

	rs, err := db.Query(b.String())

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var o order_all

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
			&o.Finance,
			&o.Branch,
			&o.Unit,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		//	set_child(&o)

		orders = append(orders, o)
	}

	return orders, err
}

func deleteOrder(db *sql.DB, id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM orders WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

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

func create_name_seq(db *sql.DB) (int64, error) {

	sqlStatement := "SELECT nextval('order_name_seq'::regclass) AS id"

	var id int64

	err := db.QueryRow(sqlStatement).Scan(&id)

	if err != nil {
		log.Printf("Unable to create order name sequence. %v", err)
	}

	return id, err
}

func createOrder(db *sql.DB, p *st_order_create) (int64, error) {

	sqlStatement := `INSERT INTO orders (
		name, order_at, printed_at, bt_finance, bt_percent, bt_matel,
		user_name, verified_by, finance_id, branch_id,
		is_stnk, stnk_price, matrix, token
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,
		to_tsvector('indonesian', $14))
	RETURNING id`

	var id int64
	token := create_new_token(p)

	err := db.QueryRow(sqlStatement,
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

type st_order_update struct {
	order_unit
	Finance models.Finance `json:"finance,omitempty"`
	Branch  models.Branch  `json:"branch,omitempty"`
}

func create_new_token(p *st_order_create) string {
	builder := strings.Builder{}

	builder.WriteString(p.Name)
	builder.WriteString(" ")
	builder.WriteString(strings.Replace(fyc.CreateIndonesianDate(p.OrderAt, true), " ", "-", -1))
	builder.WriteString(" finance ")
	builder.WriteString(p.Finance.Name)
	builder.WriteString(" ")
	builder.WriteString(p.Finance.ShortName)
	builder.WriteString(" cabang ")
	builder.WriteString(p.Branch.Name)
	builder.WriteString(" ")
	builder.WriteString(p.Branch.HeadBranch)
	builder.WriteString(" ")

	if p.IsStnk {
		builder.WriteString(" stnk-ada ")
	} else {
		builder.WriteString(" stnk-tidak-ada ")
	}

	return builder.String()

}

func create_token(p *st_order_update) string {
	builder := strings.Builder{}

	builder.WriteString(p.Name)
	builder.WriteString(" ")
	builder.WriteString(strings.Replace(fyc.CreateIndonesianDate(p.OrderAt, true), " ", "-", -1))
	builder.WriteString(" finance ")
	builder.WriteString(p.Finance.Name)
	builder.WriteString(" ")
	builder.WriteString(p.Finance.ShortName)
	builder.WriteString(" cabang ")
	builder.WriteString(p.Branch.Name)
	builder.WriteString(" ")
	builder.WriteString(p.Branch.HeadBranch)
	builder.WriteString(" ")

	if p.IsStnk {
		builder.WriteString(" stnk-ada ")
	} else {
		builder.WriteString(" stnk-tidak-ada ")
	}

	if p.Unit.TypeID > 0 {
		builder.WriteString(create_unit_token(&p.Unit))
	}

	return builder.String()

}

func updateOrder(db *sql.DB, id *int64, p *st_order_update) (int64, error) {

	sqlStatement := `UPDATE orders SET
		name=$2, order_at=$3, printed_at=$4, bt_finance=$5, bt_percent=$6, bt_matel=$7, 
		user_name=$8, verified_by=$9, finance_id=$10, branch_id=$11,
		is_stnk=$12, stnk_price=$13, matrix=$14, token=to_tsvector('indonesian', $15)
	WHERE id=$1`

	token := create_token(p)

	res, err := db.Exec(sqlStatement,
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

	return rowsAffected, err
}

func create_order_query() *strings.Builder {
	b := strings.Builder{}

	q_wheel := `SELECT id, name, short_name AS "shortName" FROM wheels WHERE id = t.wheel_id`
	q_merk := "SELECT id, name FROM merks WHERE id = t.merk_id"
	q_type := fmt.Sprintf(`SELECT t.id, t.name, t.wheel_id AS "wheelId", t.merk_id AS "merkId",
		%s as wheel, %s as merk
		FROM types t
		WHERE t.id = u.type_id`,
		fyc.NestQuerySingle(q_wheel),
		fyc.NestQuerySingle(q_merk),
	)

	q_warehouse := "SELECT id, name FROM warehouses WHERE id = u.warehouse_id"
	q_unit := fmt.Sprintf(`SELECT u.order_id AS "orderId", u.nopol, u.year, u.frame_number AS "frameNumber", 
	u.machine_number AS "machineNumber", u.color, u.type_id AS "typeId", u.warehouse_id AS "warehouseId",
	%s as type, %s as warehouse
	FROM units AS u WHERE u.order_id = o.id`,
		fyc.NestQuerySingle(q_type),
		fyc.NestQuerySingle(q_warehouse))

	b.WriteString("SELECT")
	b.WriteString(" o.id, o.name, o.order_at, o.printed_at, o.bt_finance, o.bt_percent, o.bt_matel,")
	b.WriteString(" o.user_name, o.verified_by, o.finance_id, o.branch_id, o.is_stnk, o.stnk_price, o.matrix, ")
	b.WriteString(fyc.NestQuerySingle(`SELECT id, name, short_name AS "shortName", street, city, phone, cell, zip, email, group_id AS "groupId" FROM finances WHERE id = o.finance_id`))
	b.WriteString(" AS finance, ")
	b.WriteString(fyc.NestQuerySingle(`SELECT id, name, head_branch AS "headBranch", street, city, phone, cell, zip, email FROM branchs WHERE id = o.branch_id`))
	b.WriteString(" AS branch, ")
	//b.WriteString(" COALESCE(")
	b.WriteString(fyc.NestQuerySingle(q_unit))
	//b.WriteString(", '{}') ")
	b.WriteString(" AS unit ")
	return &b
}

func searchOrders(db *sql.DB, txt *string) ([]order_all, error) {

	var orders = make([]order_all, 0)
	b := create_order_query()
	b.WriteString(" FROM orders AS o")
	b.WriteString(" WHERE token @@ to_tsquery('indonesian', $1)")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM invoice_details)")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM lents)")
	b.WriteString(" ORDER BY o.order_at, o.id")

	rs, err := db.Query(b.String(), txt)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return orders, err
	}

	defer rs.Close()

	for rs.Next() {
		var o order_all

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
			&o.Finance,
			&o.Branch,
			&o.Unit,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		//	set_child(&o)

		orders = append(orders, o)
	}

	return orders, err
}

func get_order_by_finance(db *sql.DB, id *int) ([]order_all, error) {

	var orders = make([]order_all, 0)
	b := create_order_query()
	b.WriteString(" FROM orders AS o")
	b.WriteString(" WHERE o.finance_id=$1 OR 0=$1")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM invoice_details)")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM lents)")
	b.WriteString(" ORDER BY o.order_at, o.id")

	rs, err := db.Query(b.String(), id)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var o order_all

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
			&o.Finance,
			&o.Branch,
			&o.Unit,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		//set_child(&o)

		orders = append(orders, o)
	}

	return orders, err
}

func get_order_by_branch(db *sql.DB, id *int) ([]order_all, error) {

	var orders = make([]order_all, 0)
	b := create_order_query()
	b.WriteString(" FROM orders AS o")
	b.WriteString(" WHERE o.branch_id=$1 OR 0=$1")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM invoice_details)")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM lents)")
	b.WriteString(" ORDER BY o.order_at, o.id")

	rs, err := db.Query(b.String(), id)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var o order_all

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
			&o.Finance,
			&o.Branch,
			&o.Unit,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// set_child(&o)

		orders = append(orders, o)
	}

	return orders, err
}

/*
// func set_child(o *models.Order) {
// 	finance, _ := getFinance(&o.FinanceID)
// 	o.Finance = finance

// 	branch, _ := getBranch(&o.BranchID)
// 	o.Branch = branch

// 	cust, _ := getCustomer(&o.ID)
// 	o.Customer = cust

// 	// receivable, _ := getReceivable(&o.ID)
// 	// o.Receivable = receivable

// 	unit, _ := getUnit(&o.ID)
// 	o.Unit = unit

// 	actions, _ := getAllActions(&o.ID)
// 	o.Actions = actions

// 	task, _ := getTask(&o.ID)
// 	o.Task = task

// 	home, _ := getHomeAddress(&o.ID)
// 	o.HomeAddress = home

// 	office, _ := getOfficeAddress(&o.ID)
// 	o.OfficeAddress = office

// 	post, _ := getPostAddress(&o.ID)
// 	o.PostAddress = post

// 	ktp, _ := getKTPAddress(&o.ID)
// 	o.KtpAddress = ktp
// }
*/

func get_order_by_month(db *sql.DB, id *int) ([]order_all, error) {

	var orders = make([]order_all, 0)
	b := create_order_query()
	b.WriteString(" FROM orders AS o")
	b.WriteString(" WHERE EXTRACT(MONTH from o.order_at)=$1 OR 0 = $1")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM invoice_details)")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM lents)")
	b.WriteString(" ORDER BY o.order_at, o.id")
	rs, err := db.Query(b.String(), id)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var o order_all

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
			&o.Finance,
			&o.Branch,
			&o.Unit,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		//set_child(&o)

		orders = append(orders, o)
	}

	return orders, err
}

type order_invoiced struct {
	ID int64 `json:"id"`

	Name      string  `json:"name"`
	OrderAt   string  `json:"orderAt"`
	BtFinance float64 `json:"btFinance"`
	BtPercent float32 `json:"btPercent"`
	BtMatel   float64 `json:"btMatel"`

	//VerifiedBy models.NullString `json:"verifiedBy"`

	IsStnk    bool    `json:"isStnk"`
	StnkPrice float64 `json:"stnkPrice"`
	Status    int     `json:"status"`
	FinanceId int     `json:"financeId"`

	Branch  json.RawMessage  `json:"branch,omitempty"`
	Unit    *json.RawMessage `json:"unit,omitempty"`
	Finance json.RawMessage  `json:"finance,omitempty"`
}

func order_get_invoiced(db *sql.DB, m *int, y *int, fid *int) ([]order_invoiced, error) {
	var orders = make([]order_invoiced, 0)

	var q_finance = `SELECT f.name, f.short_name "shortName" FROM finances f WHERE f.id = t.finance_id`

	var queryWheel = `SELECT name, short_name as "shortName" FROM wheels WHERE id = t.wheel_id`
	var queryMerk = `SELECT name FROM merks WHERE id = t.merk_id`

	var queryTye = fmt.Sprintf(`SELECT t.name, %s AS wheel, %s AS merk FROM types t WHERE t.id = u.type_id`,
		fyc.NestQuerySingle(queryWheel),
		fyc.NestQuerySingle(queryMerk))

	var queryUnit = fyc.NestQuerySingle(fmt.Sprintf(`SELECT u.nopol, u.year,
		%s AS type
		FROM units u
		WHERE u.order_id = t.id`,
		fyc.NestQuerySingle(queryTye)))

	var queryBranch = fyc.NestQuerySingle(`SELECT b.name FROM branchs AS b WHERE b.id = t.branch_id`)

	b := strings.Builder{}

	b.WriteString("WITH RECURSIVE rs AS(")

	b.WriteString(" SELECT 0 as status, o.id, o.name, o.order_at, o.bt_finance,")
	b.WriteString(" o.bt_percent, o.bt_matel, o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")
	//	b.WriteString(" WHERE (EXTRACT(MONTH from o.order_at)=$1 AND EXTRACT(YEAR from o.order_at)=$2")
	b.WriteString(" WHERE o.verified_by IS NULL")
	b.WriteString(" AND (o.finance_id=$3 OR 0=$3)")
	//	b.WriteString(" WHERE o.id IN (SELECT d.order_id FROM invoice_details as d WHERE d.invoice_id = $2)")

	b.WriteString(" UNION ALL")

	b.WriteString(" SELECT 1 as status, o.id, v.id::text as name, v.invoice_at,")
	b.WriteString(" o.bt_finance, v.ppn AS bt_percent,")
	b.WriteString(" o.bt_finance - (o.bt_finance * (v.ppn / 100.0)) AS bt_matel,")
	b.WriteString(" o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")
	b.WriteString(" INNER JOIN invoice_details d ON d.order_id = o.id")
	b.WriteString(" INNER JOIN invoices v ON v.id = d.invoice_id")
	b.WriteString(" WHERE (EXTRACT(MONTH from v.invoice_at)=$1 AND EXTRACT(YEAR from v.invoice_at)=$2 OR 0=$1)")
	b.WriteString(" AND (o.finance_id=$3 OR 0=$3)")
	b.WriteString(" AND o.id IN (SELECT order_id FROM invoice_details)")
	//	b.WriteString(" WHERE o.id IN (SELECT d.order_id FROM invoice_details as d WHERE d.invoice_id = $2)")

	b.WriteString(" UNION ALL")

	b.WriteString(" SELECT 2 as status, o.id, o.name, o.order_at, o.bt_finance,")
	b.WriteString(" o.bt_percent, o.bt_matel, o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")
	//b.WriteString(" WHERE (EXTRACT(MONTH from o.order_at)=$1 AND EXTRACT(YEAR from o.order_at)=$2")
	b.WriteString(" WHERE o.id NOT IN (SELECT order_id FROM invoice_details)")
	b.WriteString(" AND (o.finance_id=$3 OR 0=$3) AND o.verified_by IS NOT NULL")
	b.WriteString(")")
	// -- WHERE 0 = $2
	// -- WHERE d.invoice_id = $2

	b.WriteString(" SELECT t.status, t.id, t.name, t.order_at, t.bt_finance,")
	b.WriteString(" t.bt_percent, t.bt_matel,")
	b.WriteString(" t.is_stnk, t.stnk_price, t.finance_id, ")
	b.WriteString(queryBranch)
	b.WriteString(" AS branch, ")
	b.WriteString(queryUnit)
	b.WriteString(" AS unit, ")
	b.WriteString(fyc.NestQuerySingle(q_finance))
	b.WriteString(" AS finance ")
	b.WriteString(" FROM rs AS t")
	//b.WriteString(" WHERE t.id=1")
	b.WriteString(" ORDER BY t.status DESC, t.order_at")

	rs, err := db.Query(b.String(), m, y, fid)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p order_invoiced

		err := rs.Scan(&p.Status, &p.ID, &p.Name, &p.OrderAt, &p.BtFinance,
			&p.BtPercent, &p.BtMatel,
			&p.IsStnk, &p.StnkPrice,
			&p.FinanceId,
			&p.Branch,
			&p.Unit,
			&p.Finance,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		orders = append(orders, p)
	}

	return orders, err
}
