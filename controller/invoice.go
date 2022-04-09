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

type invoice_create_param struct {
	Invoice   models.Invoice `json:"invoice"`
	DetailIDs []int64        `json:"detailIds"`
	Token     string         `json:"token"`
	Trx       models.Trx     `json:"transaction"`
}

func Invoice_GetSearch(c *gin.Context) {
	var t models.SearchGroup

	err := c.BindJSON(&t)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	invoices, err := invoices_search(db, &t.Txt)

	if err != nil {
		//log.Printf("Unable to get all account codes. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &invoices)

}

func InvoiceGetByFinance(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	db := c.Keys["db"].(*sql.DB)
	invoices, err := invoices_by_finance(db, &id)

	if err != nil {
		//log.Printf("Unable to get all account codes. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &invoices)

}

func InvoiceGetByMonth(c *gin.Context) {

	m, _ := strconv.Atoi(c.Param("month"))
	y, _ := strconv.Atoi(c.Param("year"))

	db := c.Keys["db"].(*sql.DB)
	invoices, err := invoices_by_month(db, &m, &y)

	if err != nil {
		//log.Printf("Unable to get all account codes. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &invoices)

}

func InvoiceGetAll(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	invoices, err := invoice_get_all(db)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &invoices)
}

// router invoice.go

func InvoiceGetOrders(c *gin.Context) {

	finance_id, err := strconv.Atoi(c.Param("financeId"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	invoice_id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)

	invoices, err := invoice_get_orders(db, &finance_id, &invoice_id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &invoices)
}

// /api/invoice/item/:id
func InvoiceGetItem(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	db := c.Keys["db"].(*sql.DB)
	invoice, err := invoice_get_item(db, &id)

	if err != nil {
		//log.Printf("Unable to get all account groups. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &invoice)
}

func InvoiceCreate(c *gin.Context) {
	var param invoice_create_param

	err := c.BindJSON(&param)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	id, err := invoice_create(db, &param.Invoice, &param.Token)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	if len(param.DetailIDs) > 0 {

		err = invoice_insert_details(db, param.DetailIDs, &id)

		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
	}

	param.Trx.RefID = id
	var stoken = fmt.Sprintf("%s%s%v", param.Token, param.Trx.Descriptions, id)
	param.Trx.Descriptions = fmt.Sprintf("%s%v", param.Trx.Descriptions, id)

	trxId, err := createTransaction(db, &param.Trx, stoken)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	if len(param.Trx.Details) > 0 {

		err = bulkInsertDetails(db, param.Trx.Details, &trxId)

		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
	}

	res := Response{
		ID:      id,
		Message: "Invoice was succesfully inserted.",
	}

	c.JSON(http.StatusOK, &res)

}

func InvoiceDelete(c *gin.Context) {

	invoice_id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	invoice_delete(db, &invoice_id)

	// if err != nil {
	// 	//log.Printf("Unable to convert the string into int.  %v", err)
	// 	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 	return
	// }

	deletedRows := invoice_delete_transaction(db, &invoice_id)
	msg := fmt.Sprintf("Invoice deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      invoice_id,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func InvoiceUpdate(c *gin.Context) {

	invoice_id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var param invoice_create_param

	err = c.BindJSON(&param)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction.  %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	id, err := invoice_update(db, &invoice_id, &param.Invoice, &param.Token)

	if err != nil {
		//log.Printf("(API) Unable to update invoice.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	invocie_delete_details(db, &invoice_id)

	if len(param.DetailIDs) > 0 {

		err = invoice_insert_details(db, param.DetailIDs, &invoice_id)

		if err != nil {
			//log.Printf("Unable to insert invoice --- details.  %v", err)
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
	}

	_, err = invoice_update_transaction(db, &param.Trx.ID, &param.Trx, param.Token)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	if len(param.Trx.Details) > 0 {

		err = bulkInsertDetails(db, param.Trx.Details, &param.Trx.ID)

		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
	}

	res := Response{
		ID:      id,
		Message: "Transaction was succesfully inserted.",
	}

	c.JSON(http.StatusOK, &res)

}

func invoice_update_transaction(db *sql.DB, id *int64, p *models.Trx, token string) (int64, error) {

	b := strings.Builder{}

	b.WriteString("UPDATE trx SET")
	b.WriteString(" descriptions=$2,")
	b.WriteString(" memo=$3,")
	b.WriteString(" trx_token=to_tsvector('indonesian', $4)")
	b.WriteString(" WHERE ref_id=$1")
	b.WriteString(" AND division='trx-invoice'")

	res, err := db.Exec(b.String(),
		p.RefID,
		p.Descriptions,
		p.Memo,
		token,
	)

	if err != nil {
		log.Printf("Unable to update transaction. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func invoice_insert_details(db *sql.DB, ids []int64, id *int64) error {
	valueStrings := make([]string, 0, len(ids))
	valueArgs := make([]interface{}, 0, len(ids)*2)
	i := 0
	for _, post := range ids {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, id)
		valueArgs = append(valueArgs, post)
		i++
	}
	stmt := fmt.Sprintf("INSERT INTO invoice_details (invoice_id, order_id) VALUES %s",
		strings.Join(valueStrings, ","))
	//log.Printf("%s %v", stmt, valueArgs)
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

func invocie_delete_details(db *sql.DB, id *int64) (int64, error) {
	sqlStatement := `DELETE FROM invoice_details WHERE invoice_id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete invoice_details. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected, err
}

type invoice_item struct {
	models.Invoice
	Finance     json.RawMessage  `json:"finance,omitempty"`
	Account     json.RawMessage  `json:"account,omitempty"`
	Details     json.RawMessage  `json:"details,omitempty"`
	Transaction *json.RawMessage `json:"transaction,omitempty"`
}

func invoice_get_item(db *sql.DB, id *int64) (invoice_item, error) {
	var item invoice_item
	var queryWheel = fyc.NestQuerySingle(`SELECT id, name, short_name as "shortName" FROM wheels WHERE id = t.wheel_id`)
	var queryMerk = fyc.NestQuerySingle(`SELECT id, name FROM merks WHERE id = t.merk_id`)

	var queryType = fyc.NestQuerySingle(fmt.Sprintf(`SELECT t.id, t.name, t.wheel_id AS "wheelId", t.merk_id AS "merkId", %s AS wheel, %s AS merk FROM types t WHERE t.id = u.type_id`,
		queryWheel,
		queryMerk))

	var queryUnit = fyc.NestQuerySingle(fmt.Sprintf(`SELECT u.order_id AS "orderId", u.nopol, u.year, u.frame_number AS "frameNumber",
		u.machine_number AS "machineNumber", u.color, u.type_id AS "typeId", u.warehouse_id AS "warehouseId", %s AS type
		FROM units u
		WHERE u.order_id = o.id`,
		queryType))

	var queryDetails = fyc.NestQuery(fmt.Sprintf(`SELECT o.id, o.name, o.order_at as "orderAt", o.printed_at AS "printedAt",
	o.bt_finance as "btFinance", o.bt_percent AS "btPercent", o.bt_matel AS "btMatel",
	o.user_name AS "userName", o.verified_by AS "verifiedBy",
	o.finance_id AS "financeId", o.branch_id AS "branchId",
	o.is_stnk AS "isStnk", o.stnk_price AS "stnkPrice", matrix, true AS "isSelected", %s AS unit 
	FROM orders o
	INNER JOIN invoice_details d ON d.order_id = o.id
	WHERE d.invoice_id = v.id
	ORDER BY d.order_id`, queryUnit))

	var querFinance = fyc.NestQuerySingle(`SELECT f.id, f.name, f.short_name AS "shortName", f.street, f.city, f.phone, f.cell, f.zip, f.email, f.group_id AS "groupId" FROM finances f WHERE f.id = v.finance_id`)
	var queryAccount = fyc.NestQuerySingle(`SELECT c.id, c.name, c.type_id AS "typeId", c.descriptions, c.is_active AS "isActive", c.receivable_option AS "receivableOption", c.is_auto_debet AS "isAutoDebet" FROM acc_code c WHERE c.id = v.account_id`)

	var queryTansactionDetails = fyc.NestQuery(`SELECT id, code_id AS "codeId", trx_id AS "trxId", debt, cred FROM trx_detail WHERE trx_id = x.id`)

	var queryTansaction = fyc.NestQuerySingle(fmt.Sprintf(`SELECT x.id, x.ref_id AS "refId", x.division, x.descriptions,
	x.trx_date AS "trxDate", x.memo, %s AS details
	FROM trx x WHERE x.ref_id = v.id AND x.division = 'trx-invoice'`, queryTansactionDetails))

	b := strings.Builder{}
	b.WriteString("SELECT v.id, v.invoice_at, v.payment_term, v.due_at,")
	b.WriteString(" v.salesman, v.finance_id, v.subtotal, v.ppn, v.tax, v.total, v.account_id, v.memo,")
	b.WriteString(querFinance)
	b.WriteString(" AS finance, ")
	b.WriteString(queryAccount)
	b.WriteString(" AS account,")
	//b.WriteString(" COALESCE(")
	b.WriteString(queryTansaction)
	//b.WriteString(", '{}') AS transaction, ")
	b.WriteString(" AS transaction, ")
	b.WriteString(queryDetails)
	b.WriteString(" AS details")
	b.WriteString(" FROM invoices v")
	b.WriteString(" WHERE v.id=$1")

	rs := db.QueryRow(b.String(), id)

	err := rs.Scan(
		&item.ID,
		&item.InvoiceAt,
		&item.PaymentTerm,
		&item.DueAt,
		&item.Salesman,
		&item.FinanceID,
		&item.Subtotal,
		&item.Ppn,
		&item.Tax,
		&item.Total,
		&item.AccountId,
		&item.Memo,
		&item.Finance,
		&item.Account,
		&item.Transaction,
		&item.Details,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return item, nil
	case nil:
		return item, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return item, err
}

func invoice_delete_transaction(db *sql.DB, ref_id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM trx WHERE ref_id=$1 AND division='trx-invoice'`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, ref_id)

	if err != nil {
		log.Fatalf("Unable to delete invoice. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func invoice_delete(db *sql.DB, id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM invoices WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete invoice. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func invoice_create(db *sql.DB, inv *models.Invoice, token *string) (int64, error) {

	builder := strings.Builder{}
	builder.WriteString("INSERT INTO invoices")
	builder.WriteString(" (invoice_at, payment_term, due_at, salesman, finance_id, subtotal, ppn, tax, total, account_id, memo, token)")
	builder.WriteString(" VALUES")
	builder.WriteString(" ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, to_tsvector('indonesian', $12))")
	builder.WriteString(" RETURNING id")

	var id int64

	err := db.QueryRow(builder.String(),
		inv.InvoiceAt,
		inv.PaymentTerm,
		inv.DueAt,
		inv.Salesman,
		inv.FinanceID,
		inv.Subtotal,
		inv.Ppn,
		inv.Tax,
		inv.Total,
		inv.AccountId,
		inv.Memo,
		token,
	).Scan(&id)

	if err != nil {
		log.Printf("Unable to create invoice. %v", err)
	}

	return id, err
}

func invoice_update(db *sql.DB, id *int64, inv *models.Invoice, token *string) (int64, error) {

	builder := strings.Builder{}
	builder.WriteString("UPDATE invoices SET")

	builder.WriteString(" invoice_at=$2")
	builder.WriteString(", payment_term=$3")
	builder.WriteString(", due_at=$4")
	builder.WriteString(", salesman=$5")
	builder.WriteString(", finance_id=$6")
	builder.WriteString(", subtotal=$7")
	builder.WriteString(", ppn=$8")
	builder.WriteString(", tax=$9")
	builder.WriteString(", total=$10")
	builder.WriteString(", account_id=$11")
	builder.WriteString(", memo=$12")
	builder.WriteString(", token=to_tsvector('indonesian', $13)")
	builder.WriteString(" WHERE id=$1")

	res, err := db.Exec(builder.String(),
		id,
		inv.InvoiceAt,
		inv.PaymentTerm,
		inv.DueAt,
		inv.Salesman,
		inv.FinanceID,
		inv.Subtotal,
		inv.Ppn,
		inv.Tax,
		inv.Total,
		inv.AccountId,
		inv.Memo,
		token,
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

type invoice_all struct {
	models.Invoice
	Finance json.RawMessage `json:"finance,omitempty"`
	Account json.RawMessage `json:"account,omitempty"`
}

func invoice_get_all(db *sql.DB) ([]invoice_all, error) {
	var invoices = make([]invoice_all, 0)

	builder := strings.Builder{}

	var querFinance = `SELECT f.id, f.name, f.short_name "shortName", f.street, f.city, f.phone, f.cell, f.zip, f.email, f.group_id AS "groupId" FROM finances f WHERE f.id = v.finance_id`
	var queryAccount = `SELECT c.id, c.name, c.type_id AS "typeId", c.descriptions, c.is_active AS "isActive", c.receivable_option AS "receivableOption", c.is_auto_debet AS "isAutoDebet" FROM acc_code c WHERE c.id = v.account_id`

	builder.WriteString("SELECT")
	builder.WriteString(" v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman, v.finance_id, v.subtotal, v.ppn, v.tax, v.total, v.account_id, v.memo, ")
	builder.WriteString(fyc.NestQuerySingle(querFinance))
	builder.WriteString(" AS finance,")
	builder.WriteString(fyc.NestQuerySingle(queryAccount))
	builder.WriteString(" AS account")
	builder.WriteString(" FROM invoices AS v")
	builder.WriteString(" ORDER BY v.id DESC")

	// var sqlStatement = fmt.Sprintf("SELECT v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman, v.finance_id, v.subtotal, v.ppn, v.tax, v.total, v.account_id, v.memo, %s AS finance, %s AS account FROM invoices AS v ORDER BY v.id DESC",
	// 	nestQuerySingle(querFinance),
	// 	nestQuerySingle(queryAccount),
	// )

	rs, err := db.Query(builder.String())

	if err != nil {
		log.Fatalf("Unable to execute merks query %v", err)
		return invoices, err
	}

	defer rs.Close()

	for rs.Next() {
		var item invoice_all

		err := rs.Scan(
			&item.ID,
			&item.InvoiceAt,
			&item.PaymentTerm,
			&item.DueAt,
			&item.Salesman,
			&item.FinanceID,
			&item.Subtotal,
			&item.Ppn,
			&item.Tax,
			&item.Total,
			&item.AccountId,
			&item.Memo,
			&item.Finance,
			&item.Account,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		invoices = append(invoices, item)
	}

	return invoices, err
}

type invoice_order struct {
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

	IsSelected bool             `json:"isSelected"`
	Branch     json.RawMessage  `json:"branch,omitempty"`
	Unit       *json.RawMessage `json:"unit,omitempty"`
}

func invoice_get_orders(db *sql.DB, finance_id *int, invoice_id *int64) ([]invoice_order, error) {

	var invoices = make([]invoice_order, 0)

	var queryWheel = `SELECT id, name, short_name as "shortName" FROM wheels WHERE id = t.wheel_id`
	var queryMerk = `SELECT id, name FROM merks WHERE id = t.merk_id`

	var queryTye = fmt.Sprintf(`SELECT t.id, t.name, t.wheel_id AS "wheelId", t.merk_id AS "merkId", %s AS wheel, %s AS merk FROM types t WHERE t.id = u.type_id`,
		fyc.NestQuerySingle(queryWheel),
		fyc.NestQuerySingle(queryMerk))

	var queryUnit = fyc.NestQuerySingle(fmt.Sprintf(`SELECT u.order_id AS "orderId", u.nopol, u.year, u.frame_number AS "frameNumber",
		u.machine_number AS "machineNumber", u.color, u.type_id AS "typeId", u.warehouse_id AS "warehouseId", %s AS type
		FROM units u
		WHERE u.order_id = o.id`,
		fyc.NestQuerySingle(queryTye)))

	var queryBranch = fyc.NestQuerySingle(`SELECT b.id, b.name, b.street, b.city, b.phone,
	b.cell, b.zip, b.head_branch AS "headBranch", b.email
	FROM branchs AS b
	WHERE b.id = o.branch_id`)

	b := strings.Builder{}

	b.WriteString("WITH RECURSIVE rs AS(")
	b.WriteString(" SELECT true as is_selected, o.id, o.name, o.order_at, o.printed_at, o.bt_finance, o.bt_percent, o.bt_matel,")
	b.WriteString(" o.user_name, o.verified_by, o.finance_id, o.branch_id,")
	b.WriteString(" o.is_stnk, o.stnk_price, o.matrix, ")
	b.WriteString(queryBranch)
	b.WriteString(" AS branch, ")
	b.WriteString(queryUnit)
	b.WriteString(" AS unit ")
	b.WriteString(" FROM orders AS o")
	b.WriteString(" WHERE o.id IN (SELECT d.order_id FROM invoice_details as d WHERE d.invoice_id = $2)")

	b.WriteString(" UNION ALL")

	b.WriteString(" SELECT false as is_selected, o.id, o.name, o.order_at, o.printed_at, o.bt_finance, o.bt_percent, o.bt_matel,")
	b.WriteString(" o.user_name, o.verified_by, o.finance_id, o.branch_id,")
	b.WriteString(" o.is_stnk, o.stnk_price, o.matrix, ")
	b.WriteString(queryBranch)
	b.WriteString(" AS branch, ")
	b.WriteString(queryUnit)
	b.WriteString(" AS unit")
	b.WriteString(" FROM orders AS o")
	b.WriteString(" WHERE o.finance_id=$1")
	b.WriteString(" AND o.verified_by IS NOT NULL")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM invoice_details)")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM lents)")
	b.WriteString(")")
	// -- WHERE 0 = $2
	// -- WHERE d.invoice_id = $2

	b.WriteString(" SELECT t.is_selected, t.id, t.name, t.order_at, t.printed_at, t.bt_finance,")
	b.WriteString(" t.bt_percent, t.bt_matel, t.user_name, t.verified_by,")
	b.WriteString(" t.finance_id, t.branch_id, t.is_stnk, t.stnk_price, t.matrix, t.branch, t.unit")
	b.WriteString(" FROM rs AS t")
	b.WriteString(" ORDER BY t.is_selected DESC, t.finance_id, t.id DESC")
	// `,
	// 	,
	// 	,
	// 	nestQuerySingle(queryBranch),
	// 	nestQuerySingle(queryUnit),
	// )

	rs, err := db.Query(b.String(), finance_id, invoice_id)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var invoice invoice_order

		err := rs.Scan(
			&invoice.IsSelected,
			&invoice.ID,
			&invoice.Name,
			&invoice.OrderAt,
			&invoice.PrintedAt,
			&invoice.BtFinance,
			&invoice.BtPercent,
			&invoice.BtMatel,
			//	&invoice.Ppn,
			//	&invoice.Nominal,
			//			&invoice.Subtotal,
			&invoice.UserName,
			&invoice.VerifiedBy,
			//	&invoice.ValidatedBy,
			&invoice.FinanceID,
			&invoice.BranchID,
			&invoice.IsStnk,
			&invoice.StnkPrice,
			&invoice.Matrix,
			&invoice.Branch,
			&invoice.Unit,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		invoices = append(invoices, invoice)
	}

	return invoices, err

}

func invoices_search(db *sql.DB, txt *string) ([]invoice_all, error) {
	var invoices = make([]invoice_all, 0)

	b := strings.Builder{}

	var querFinance = `SELECT f.id, f.name, f.short_name "shortName", f.street, f.city, f.phone, f.cell, f.zip, f.email, f.group_id AS "groupId" FROM finances f WHERE f.id = v.finance_id`
	var queryAccount = `SELECT c.id, c.name, c.type_id AS "typeId", c.descriptions, c.is_active AS "isActive", c.receivable_option AS "receivableOption", c.is_auto_debet AS "isAutoDebet" FROM acc_code c WHERE c.id = v.account_id`

	b.WriteString("SELECT v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman,")
	b.WriteString(" v.finance_id, v.subtotal, v.ppn, v.tax, v.total, v.account_id, v.memo,")
	b.WriteString(fyc.NestQuerySingle(querFinance))
	b.WriteString("	AS finance, ")
	b.WriteString(fyc.NestQuerySingle(queryAccount))
	b.WriteString(" AS account ")
	b.WriteString(" FROM invoices AS v")
	b.WriteString(" WHERE token @@ to_tsquery('indonesian', $1)")
	b.WriteString(" ORDER BY v.id DESC")

	rs, err := db.Query(b.String(), txt)

	if err != nil {
		log.Fatalf("Unable to execute merks query %v", err)
		return invoices, err
	}

	defer rs.Close()

	for rs.Next() {
		var item invoice_all

		err := rs.Scan(
			&item.ID,
			&item.InvoiceAt,
			&item.PaymentTerm,
			&item.DueAt,
			&item.Salesman,
			&item.FinanceID,
			&item.Subtotal,
			&item.Ppn,
			&item.Tax,
			&item.Total,
			&item.AccountId,
			&item.Memo,
			&item.Finance,
			&item.Account,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		invoices = append(invoices, item)
	}

	return invoices, err
}

func invoices_by_month(db *sql.DB, month *int, year *int) ([]invoice_all, error) {
	var invoices = make([]invoice_all, 0)

	builder := strings.Builder{}

	var querFinance = `SELECT f.id, f.name, f.short_name "shortName", f.street, f.city, f.phone, f.cell, f.zip, f.email, f.group_id AS "groupId" FROM finances f WHERE f.id = v.finance_id`
	var queryAccount = `SELECT c.id, c.name, c.type_id AS "typeId", c.descriptions, c.is_active AS "isActive", c.receivable_option AS "receivableOption", c.is_auto_debet AS "isAutoDebet" FROM acc_code c WHERE c.id = v.account_id`

	// var sqlStatement = fmt.Sprintf(`
	// 	%s AS finance, %s AS account
	// 	FROM invoices AS v
	// 	WHERE EXTRACT(MONTH FROM v.invoice_at)=$1 AND EXTRACT(YEAR FROM v.invoice_at)=$2 OR 0=$1
	// 	ORDER BY v.id DESC`,
	// 	nestQuerySingle(querFinance),
	// 	nestQuerySingle(queryAccount),
	// )

	builder.WriteString("SELECT v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman, v.finance_id, v.subtotal, v.ppn, v.tax, v.total, v.account_id, v.memo, ")
	builder.WriteString(fyc.NestQuerySingle(querFinance))
	builder.WriteString(" AS finance, ")
	builder.WriteString(fyc.NestQuerySingle(queryAccount))
	builder.WriteString(" AS account")
	builder.WriteString(" FROM invoices AS v")
	builder.WriteString(" WHERE EXTRACT(MONTH FROM v.invoice_at)=$1")
	builder.WriteString(" AND EXTRACT(YEAR FROM v.invoice_at)=$2")
	builder.WriteString(" OR 0=$1")
	builder.WriteString(" ORDER BY v.id DESC")

	rs, err := db.Query(builder.String(), month, year)

	if err != nil {
		log.Fatalf("Unable to execute merks query %v", err)
		return invoices, err
	}

	defer rs.Close()

	for rs.Next() {
		var item invoice_all

		err := rs.Scan(
			&item.ID,
			&item.InvoiceAt,
			&item.PaymentTerm,
			&item.DueAt,
			&item.Salesman,
			&item.FinanceID,
			&item.Subtotal,
			&item.Ppn,
			&item.Tax,
			&item.Total,
			&item.AccountId,
			&item.Memo,
			&item.Finance,
			&item.Account,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		invoices = append(invoices, item)
	}

	return invoices, err
}

func invoices_by_finance(db *sql.DB, finance_id *int) ([]invoice_all, error) {
	var invoices = make([]invoice_all, 0)
	var querFinance = `SELECT f.id, f.name, f.short_name "shortName", f.street, f.city, f.phone, f.cell, f.zip, f.email, f.group_id AS "groupId" FROM finances f WHERE f.id = v.finance_id`
	var queryAccount = `SELECT c.id, c.name, c.type_id AS "typeId", c.descriptions, c.is_active AS "isActive", c.receivable_option AS "receivableOption", c.is_auto_debet AS "isAutoDebet" FROM acc_code c WHERE c.id = v.account_id`

	b := strings.Builder{}

	b.WriteString("SELECT v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman,")
	b.WriteString(" v.finance_id, v.subtotal, v.ppn, v.tax, v.total, v.account_id, v.memo, ")
	b.WriteString(fyc.NestQuerySingle(querFinance))
	b.WriteString(" AS finance, ")
	b.WriteString(fyc.NestQuerySingle(queryAccount))
	b.WriteString(" AS account")
	b.WriteString(" FROM invoices AS v")
	b.WriteString(" WHERE v.finance_id=$1 OR 0=$1")
	b.WriteString(" ORDER BY v.id DESC")

	rs, err := db.Query(b.String(), finance_id)

	if err != nil {
		log.Fatalf("Unable to execute merks query %v", err)
		return invoices, err
	}

	defer rs.Close()

	for rs.Next() {
		var item invoice_all

		err := rs.Scan(
			&item.ID,
			&item.InvoiceAt,
			&item.PaymentTerm,
			&item.DueAt,
			&item.Salesman,
			&item.FinanceID,
			&item.Subtotal,
			&item.Ppn,
			&item.Tax,
			&item.Total,
			&item.AccountId,
			&item.Memo,
			&item.Finance,
			&item.Account,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		invoices = append(invoices, item)
	}

	return invoices, err
}

func invoice_get_item_customer(db *sql.DB, id *int64) (invoice_item, error) {
	var item invoice_item
	var queryWheel = fyc.NestQuerySingle(`SELECT id, name, short_name as "shortName" FROM wheels WHERE id = t.wheel_id`)
	var queryMerk = fyc.NestQuerySingle(`SELECT id, name FROM merks WHERE id = t.merk_id`)

	var queryType = fyc.NestQuerySingle(fmt.Sprintf(`SELECT t.id, t.name, t.wheel_id AS "wheelId", t.merk_id AS "merkId", %s AS wheel, %s AS merk FROM types t WHERE t.id = u.type_id`,
		queryWheel,
		queryMerk))

	var queryUnit = fyc.NestQuerySingle(fmt.Sprintf(`SELECT u.order_id AS "orderId", u.nopol, u.year, u.frame_number AS "frameNumber",
		u.machine_number AS "machineNumber", u.color, u.type_id AS "typeId", u.warehouse_id AS "warehouseId", %s AS type
		FROM units u
		WHERE u.order_id = o.id`,
		queryType))

	var q_customer = fyc.NestQuerySingle("SELECT order_id, name, agreement_number, payment_type FROM customers WHERE order_id=o.id")

	var queryDetails = fyc.NestQuery(fmt.Sprintf(`SELECT o.id, o.name, o.order_at as "orderAt", o.printed_at AS "printedAt",
	o.bt_finance as "btFinance", o.bt_percent AS "btPercent", o.bt_matel AS "btMatel",
	o.user_name AS "userName", o.verified_by AS "verifiedBy",
	o.finance_id AS "financeId", o.branch_id AS "branchId",
	o.is_stnk AS "isStnk", o.stnk_price AS "stnkPrice", matrix, true AS "isSelected", %s AS unit,
	%s as customer
	FROM orders o
	INNER JOIN invoice_details d ON d.order_id = o.id
	WHERE d.invoice_id = v.id
	ORDER BY d.order_id`, queryUnit, q_customer))

	var querFinance = fyc.NestQuerySingle(`SELECT f.id, f.name, f.short_name AS "shortName", f.street, f.city, f.phone, f.cell, f.zip, f.email, f.group_id AS "groupId" FROM finances f WHERE f.id = v.finance_id`)
	var queryAccount = fyc.NestQuerySingle(`SELECT c.id, c.name, c.type_id AS "typeId", c.descriptions, c.is_active AS "isActive", c.receivable_option AS "receivableOption", c.is_auto_debet AS "isAutoDebet" FROM acc_code c WHERE c.id = v.account_id`)

	var queryTansactionDetails = fyc.NestQuery(`SELECT id, code_id AS "codeId", trx_id AS "trxId", debt, cred FROM trx_detail WHERE trx_id = x.id`)

	var queryTansaction = fyc.NestQuerySingle(fmt.Sprintf(`SELECT x.id, x.ref_id AS "refId", x.division, x.descriptions,
	x.trx_date AS "trxDate", x.memo, %s AS details
	FROM trx x WHERE x.ref_id = v.id AND x.division = 'trx-invoice'`, queryTansactionDetails))

	b := strings.Builder{}
	b.WriteString("SELECT v.id, v.invoice_at, v.payment_term, v.due_at,")
	b.WriteString(" v.salesman, v.finance_id, v.subtotal, v.ppn, v.tax, v.total, v.account_id, v.memo,")
	b.WriteString(querFinance)
	b.WriteString(" AS finance, ")
	b.WriteString(queryAccount)
	b.WriteString(" AS account,")
	//b.WriteString(" COALESCE(")
	b.WriteString(queryTansaction)
	//b.WriteString(", '{}') AS transaction, ")
	b.WriteString(" AS transaction, ")
	b.WriteString(queryDetails)
	b.WriteString(" AS details")
	b.WriteString(" FROM invoices v")
	b.WriteString(" WHERE v.id=$1")

	rs := db.QueryRow(b.String(), id)

	err := rs.Scan(
		&item.ID,
		&item.InvoiceAt,
		&item.PaymentTerm,
		&item.DueAt,
		&item.Salesman,
		&item.FinanceID,
		&item.Subtotal,
		&item.Ppn,
		&item.Tax,
		&item.Total,
		&item.AccountId,
		&item.Memo,
		&item.Finance,
		&item.Account,
		&item.Transaction,
		&item.Details,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return item, nil
	case nil:
		return item, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return item, err
}
