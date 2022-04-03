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

	"github.com/MasturFYC/fyc"
	"github.com/gorilla/mux"
)

type invoice_create_param struct {
	Invoice   models.Invoice `json:"invoice"`
	DetailIDs []int64        `json:"detailIds"`
	Token     string         `json:"token"`
	Trx       models.Trx     `json:"transaction"`
}

func Invoice_GetSearch(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var t models.SearchGroup

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	invoices, err := invoices_search(&t.Txt)

	if err != nil || len(invoices) == 0 {
		//log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&invoices)

}

func Invoice_GetByFinance(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	invoices, err := invoices_by_finance(&id)

	if err != nil || len(invoices) == 0 {
		//log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&invoices)

}

func Invoice_GetByMonth(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	m, _ := strconv.Atoi(params["month"])
	y, _ := strconv.Atoi(params["year"])

	invoices, err := invoices_by_month(&m, &y)

	if err != nil || len(invoices) == 0 {
		//log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&invoices)

}

func Invoice_GetAll(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	invoices, err := invoice_get_all()

	if err != nil {
		log.Fatalf("Unable to get all merks. %v", err)
	}

	json.NewEncoder(w).Encode(&invoices)
}

// router invoice.go

func Invoice_GetOrders(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	finance_id, err := strconv.Atoi(params["financeId"])

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	invoice_id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	invoices, err := invoice_get_orders(&finance_id, &invoice_id)

	if err != nil || len(invoices) == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&invoices)
}

func Invoice_GetItem(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	invoice, err := invoice_get_item(&id)

	if err != nil {
		//log.Printf("Unable to get all account groups. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&invoice)
}

func Invoice_Create(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var param invoice_create_param

	err := json.NewDecoder(r.Body).Decode(&param)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := invoice_create(&param.Invoice, &param.Token)

	if err != nil {
		//log.Printf("(API) Unable to create transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(param.DetailIDs) > 0 {

		err = invoice_insert_details(param.DetailIDs, &id)

		if err != nil {
			//log.Printf("Unable to insert invoice details.  %v", err)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}

	param.Trx.RefID = id
	var stoken = fmt.Sprintf("%s%s%v", param.Token, param.Trx.Descriptions, id)
	param.Trx.Descriptions = fmt.Sprintf("%s%v", param.Trx.Descriptions, id)

	trxId, err := createTransaction(&param.Trx, stoken)

	if err != nil {
		//log.Printf("(API) Unable to create transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(param.Trx.Details) > 0 {

		err = bulkInsertDetails(param.Trx.Details, &trxId)

		if err != nil {
			//log.Printf("Unable to insert transaction details.  %v", err)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}

	res := Response{
		ID:      id,
		Message: "Invoice was succesfully inserted.",
	}

	json.NewEncoder(w).Encode(&res)

}

func Invoice_Delete(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)

	invoice_id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	invoice_delete(&invoice_id)

	// if err != nil {
	// 	//log.Printf("Unable to convert the string into int.  %v", err)
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }

	deletedRows := invoice_delete_transaction(&invoice_id)
	msg := fmt.Sprintf("Invoice deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      invoice_id,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func Invoice_Update(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	invoice_id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var param invoice_create_param

	err = json.NewDecoder(r.Body).Decode(&param)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := invoice_update(&invoice_id, &param.Invoice, &param.Token)

	if err != nil {
		//log.Printf("(API) Unable to update invoice.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	invocie_delete_details(&invoice_id)

	if len(param.DetailIDs) > 0 {

		err = invoice_insert_details(param.DetailIDs, &invoice_id)

		if err != nil {
			//log.Printf("Unable to insert invoice --- details.  %v", err)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}

	_, err = invoice_update_transaction(&param.Trx.ID, &param.Trx, param.Token)

	if err != nil {
		//log.Printf("(API) Unable to create transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(param.Trx.Details) > 0 {

		err = bulkInsertDetails(param.Trx.Details, &param.Trx.ID)

		if err != nil {
			//log.Printf("Unable to insert transaction details from invoices.  %v", err)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}

	res := Response{
		ID:      id,
		Message: "Transaction was succesfully inserted.",
	}

	json.NewEncoder(w).Encode(&res)

}

func invoice_update_transaction(id *int64, p *models.Trx, token string) (int64, error) {

	b := strings.Builder{}

	b.WriteString("UPDATE trx SET")
	b.WriteString(" descriptions=$2,")
	b.WriteString(" memo=$3,")
	b.WriteString(" trx_token=to_tsvector('indonesian', $4)")
	b.WriteString(" WHERE ref_id=$1")
	b.WriteString(" AND division='trx-invoice'")

	res, err := Sql().Exec(b.String(),
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

	if err != nil {
		log.Printf("Error while updating transaction. %v", err)
		return 0, err
	}

	return rowsAffected, err
}

func invoice_insert_details(ids []int64, id *int64) error {
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
	_, err := Sql().Exec(stmt, valueArgs...)
	return err
}

func invocie_delete_details(id *int64) (int64, error) {
	sqlStatement := `DELETE FROM invoice_details WHERE invoice_id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

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

func invoice_get_item(id *int64) (invoice_item, error) {
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

	rs := Sql().QueryRow(b.String(), id)

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

func invoice_delete_transaction(ref_id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM trx WHERE ref_id=$1 AND division='trx-invoice'`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, ref_id)

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

func invoice_delete(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM invoices WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

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

func invoice_create(inv *models.Invoice, token *string) (int64, error) {

	builder := strings.Builder{}
	builder.WriteString("INSERT INTO invoices")
	builder.WriteString(" (invoice_at, payment_term, due_at, salesman, finance_id, subtotal, ppn, tax, total, account_id, memo, token)")
	builder.WriteString(" VALUES")
	builder.WriteString(" ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, to_tsvector('indonesian', $12))")
	builder.WriteString(" RETURNING id")

	var id int64

	err := Sql().QueryRow(builder.String(),
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

func invoice_update(id *int64, inv *models.Invoice, token *string) (int64, error) {

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

	res, err := Sql().Exec(builder.String(),
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

func invoice_get_all() ([]invoice_all, error) {
	var invoices []invoice_all

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

	rs, err := Sql().Query(builder.String())

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

func invoice_get_orders(finance_id *int, invoice_id *int64) ([]invoice_order, error) {

	var invoices []invoice_order

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

	rs, err := Sql().Query(b.String(), finance_id, invoice_id)

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

func invoices_search(txt *string) ([]invoice_all, error) {
	var invoices []invoice_all

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

	rs, err := Sql().Query(b.String(), txt)

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

func invoices_by_month(month *int, year *int) ([]invoice_all, error) {
	var invoices []invoice_all

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

	rs, err := Sql().Query(builder.String(), month, year)

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

func invoices_by_finance(finance_id *int) ([]invoice_all, error) {
	var invoices []invoice_all
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

	rs, err := Sql().Query(b.String(), finance_id)

	//log.Println(b.String())

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

func invoice_get_item_customer(id *int64) (invoice_item, error) {
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

	rs := Sql().QueryRow(b.String(), id)

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
