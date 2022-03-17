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
		log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
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

	id, err := strconv.Atoi(params["id"])

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

	m, err := strconv.Atoi(params["month"])
	y, err := strconv.Atoi(params["year"])

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

	id, err := strconv.ParseInt(params["id"], 10, 64)

	invoice, err := invoice_get_item(&id)

	if err != nil {
		log.Printf("Unable to get all account groups. %v", err)
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
		log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	id, err := invoice_create(&param.Invoice, &param.Token)

	if err != nil {
		log.Printf("(API) Unable to create transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(param.DetailIDs) > 0 {

		err = invoice_insert_details(param.DetailIDs, &id)

		if err != nil {
			log.Printf("Unable to insert invoice details.  %v", err)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}

	param.Trx.RefID = id
	var stoken = fmt.Sprintf("%s%s%v", param.Token, param.Trx.Descriptions, id)
	param.Trx.Descriptions = fmt.Sprintf("%s%v", param.Trx.Descriptions, id)

	trxId, err := createTransaction(&param.Trx, stoken)

	if err != nil {
		log.Printf("(API) Unable to create transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(param.Trx.Details) > 0 {

		err = bulkInsertDetails(param.Trx.Details, &trxId)

		if err != nil {
			log.Printf("Unable to insert transaction details.  %v", err)
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
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows := invoice_delete(&invoice_id)
	deletedRows = invoice_delete_transaction(&invoice_id)
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
		log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var param invoice_create_param

	err = json.NewDecoder(r.Body).Decode(&param)

	if err != nil {
		log.Printf("Unable to decode the request body to transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	id, err := invoice_update(&invoice_id, &param.Invoice, &param.Token)

	if err != nil {
		log.Printf("(API) Unable to update invoice.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	invocie_delete_details(&invoice_id)

	if len(param.DetailIDs) > 0 {

		err = invoice_insert_details(param.DetailIDs, &invoice_id)

		if err != nil {
			log.Printf("Unable to insert invoice --- details.  %v", err)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}

	_, err = invocie_update_transaction(&param.Trx.ID, &param.Trx, param.Token)

	if err != nil {
		log.Printf("(API) Unable to create transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(param.Trx.Details) > 0 {

		err = bulkInsertDetails(param.Trx.Details, &param.Trx.ID)

		if err != nil {
			log.Printf("Unable to insert transaction details from invoices.  %v", err)
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

func invocie_update_transaction(id *int64, p *models.Trx, token string) (int64, error) {

	sqlStatement := `UPDATE trx SET
		descriptions=$2,
		memo=$3,
		trx_token=to_tsvector('indonesian', $4)
	WHERE ref_id=$1`

	res, err := Sql().Exec(sqlStatement,
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
	Finance     json.RawMessage `json:"finance,ommitempty"`
	Account     json.RawMessage `json:"account,ommitempty"`
	Details     json.RawMessage `json:"details,ommitempty"`
	Transaction json.RawMessage `json:"transaction,ommitempty"`
}

func invoice_get_item(id *int64) (invoice_item, error) {
	var item invoice_item
	var queryWheel = `SELECT w.id, w.name, w.short_name as "shortName" FROM wheels w WHERE w.id = t.wheel_id`
	var queryMerk = `SELECT m.id, m.name FROM merks m WHERE m.id = t.merk_id`

	var queryTye = fmt.Sprintf(`SELECT t.id, t.name, t.wheel_id AS "wheelId", t.merk_id AS "merkId", %s AS wheel, %s AS merk FROM types t WHERE t.id = u.type_id`,
		nestQuerySingle(queryWheel),
		nestQuerySingle(queryMerk))

	var queryUnit = fmt.Sprintf(`SELECT u.order_id AS "orderId", u.nopol, u.year, u.frame_number AS "frameNumber",
		u.machine_number AS "machineNumber", u.bpkb_name AS "bpkbName",
		color, dealer, surveyor, u.type_id AS "typeId", u.warehouse_id AS "warehouseId", %s AS type
		FROM units u
		WHERE u.order_id = o.id`,
		nestQuerySingle(queryTye))

	var queryDetails = fmt.Sprintf(`SELECT o.id, o.name, o.order_at as "orderAt", o.printed_at AS "printedAt",
	o.bt_finance as "btFinance", o.bt_percent AS "btPercent", o.bt_matel AS "btMatel", o.ppn,
		o.nominal, o.subtotal, o.user_name AS "userName", o.verified_by AS "verifiedBy",
		o.validated_by AS "validatedBy", o.finance_id AS "financeId", o.branch_id AS "branchId",
		o.is_stnk AS "isStnk", o.stnk_price AS "stnkPrice", matrix, true AS "isSelected", %s AS unit 
	FROM orders o
	INNER JOIN invoice_details d ON d.order_id = o.id
	WHERE d.invoice_id = v.id`, nestQuerySingle(queryUnit))

	var querFinance = `SELECT f.id, f.name, f.short_name AS "shortName", f.street, f.city, f.phone, f.cell, f.zip, f.email FROM finances f WHERE f.id = v.finance_id`
	var queryAccount = `SELECT c.id, c.name, c.type_id AS "typeId", c.descriptions, c.is_active AS "isActive", c.receivable_option AS "receivableOption", c.is_auto_debet AS "isAutoDebet" FROM acc_code c WHERE c.id = v.account_id`

	var queryTansactionDetails = `SELECT id, code_id AS "codeId", trx_id AS "trxId", debt, cred FROM trx_detail WHERE trx_id = x.id`

	var queryTansaction = fmt.Sprintf(`SELECT x.id, x.ref_id AS "refId", x.division, x.descriptions,
	x.trx_date AS "trxDate", x.memo, %s AS details
	FROM trx x WHERE x.ref_id = v.id`, nestQuery(queryTansactionDetails))

	var sqlStatement = fmt.Sprintf(`SELECT v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman, v.finance_id, memo, total, tax, account_id,
%s AS finance,
%s AS account,
COALESCE(%s, '{}') AS transaction,
%s AS details
FROM invoices v
WHERE v.id=$1`,
		nestQuerySingle(querFinance),
		nestQuerySingle(queryAccount),
		nestQuerySingle(queryTansaction),
		nestQuery(queryDetails))

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(
		&item.ID,
		&item.InvoiceAt,
		&item.PaymentTerm,
		&item.DueAt,
		&item.Salesman,
		&item.FinanceID,
		&item.Memo,
		&item.Total,
		&item.Tax,
		&item.AccountId,
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
	sqlStatement := `DELETE FROM trx WHERE ref_id=$1`

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

	sqlStatement := `INSERT INTO invoices 
	(invoice_at, payment_term, due_at, salesman, finance_id, memo, total, tax, account_id, token) 
	VALUES 
	($1, $2, $3, $4, $5, $6, $7, $8, $9, to_tsvector('indonesian', $10))
	RETURNING id`

	var id int64

	err := Sql().QueryRow(sqlStatement,
		inv.InvoiceAt,
		inv.PaymentTerm,
		inv.DueAt,
		inv.Salesman,
		inv.FinanceID,
		inv.Memo,
		inv.Total,
		inv.Tax,
		inv.AccountId,
		token,
	).Scan(&id)

	if err != nil {
		log.Printf("Unable to create invoice. %v", err)
	}

	return id, err
}

func invoice_update(id *int64, inv *models.Invoice, token *string) (int64, error) {

	sqlStatement := `UPDATE invoices SET
	invoice_at=$1,
	payment_term=$2,
	due_at=$3,
	salesman=$4,
	finance_id=$5,
	memo=$6,
	total=$7,
	tax=$8,
	account_id=$9,
	token=to_tsvector('indonesian', $10)
	WHERE id=$11`

	res, err := Sql().Exec(sqlStatement,
		inv.InvoiceAt,
		inv.PaymentTerm,
		inv.DueAt,
		inv.Salesman,
		inv.FinanceID,
		inv.Memo,
		inv.Total,
		inv.Tax,
		inv.AccountId,
		token,
		id,
	)

	if err != nil {
		log.Printf("Unable to update finance. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating finance. %v", err)
	}

	return rowsAffected, err
}

type invoice_all struct {
	models.Invoice
	Finance json.RawMessage `json:"finance,ommitempty"`
	Account json.RawMessage `json:"account,ommitempty"`
}

func invoice_get_all() ([]invoice_all, error) {
	var invoices []invoice_all
	var querFinance = `SELECT f.id, f.name, f.short_name "shortName", f.street, f.city, f.phone, f.cell, f.zip, f.email FROM finances f WHERE f.id = v.finance_id`
	var queryAccount = `SELECT c.id, c.name, c.type_id AS "typeId", c.descriptions, c.is_active AS "isActive", c.receivable_option AS "receivableOption", c.is_auto_debet AS "isAutoDebet" FROM acc_code c WHERE c.id = v.account_id`

	var sqlStatement = fmt.Sprintf("SELECT v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman, v.finance_id, memo, total, tax, account_id, %s AS finance, %s AS account FROM invoices AS v ORDER BY v.id DESC",
		nestQuerySingle(querFinance),
		nestQuerySingle(queryAccount),
	)

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute merks query %v", err)
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
			&item.Memo,
			&item.Total,
			&item.Tax,
			&item.AccountId,
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
	models.Order
	IsSelected bool            `json:"isSelected"`
	Branch     json.RawMessage `json:"branch,omitempty"`
	Unit       json.RawMessage `json:"unit,omitempty"`
}

func invoice_get_orders(finance_id *int, invoice_id *int64) ([]invoice_order, error) {

	var invoices []invoice_order

	var queryWheel = `SELECT id, name, short_name as "shortName" FROM wheels WHERE id = t.wheel_id`
	var queryMerk = `SELECT id, name FROM merks WHERE id = t.merk_id`

	var queryTye = fmt.Sprintf(`SELECT t.id, t.name, t.wheel_id AS "wheelId", t.merk_id AS "merkId", %s AS wheel, %s AS merk FROM types t WHERE t.id = u.type_id`,
		nestQuerySingle(queryWheel),
		nestQuerySingle(queryMerk))

	var queryUnit = fmt.Sprintf(`SELECT u.order_id AS "orderId", u.nopol, u.year, u.frame_number AS "frameNumber",
		u.machine_number AS "machineNumber", u.bpkb_name AS "bpkbName",
		color, dealer, surveyor, u.type_id AS "typeId", u.warehouse_id AS "warehouseId", %s AS type
		FROM units u
		WHERE u.order_id = o.id`,
		nestQuerySingle(queryTye))

	var queryBranch = `SELECT b.id, b.name, b.street, b.city, b.phone,
	b.cell, b.zip, b.head_branch AS "headBranch", b.email
	FROM branchs AS b
	WHERE b.id = o.branch_id`

	var sqlStatement = fmt.Sprintf(`WITH RECURSIVE rs AS(
		SELECT true as is_selected, o.id, o.name, o.order_at, o.printed_at, o.bt_finance, o.bt_percent, o.bt_matel, o.ppn,
		o.nominal, o.subtotal, o.user_name, o.verified_by, o.validated_by, o.finance_id, o.branch_id,
		o.is_stnk, o.stnk_price, o.matrix,
		%s AS branch,
		%s AS unit
	FROM orders AS o
	WHERE o.id IN (SELECT d.order_id FROM invoice_details as d WHERE d.invoice_id = $2)
	
	UNION ALL

	SELECT false as is_selected, o.id, o.name, o.order_at, o.printed_at, o.bt_finance, o.bt_percent, o.bt_matel, o.ppn,
		o.nominal, o.subtotal, o.user_name, o.verified_by, o.validated_by, o.finance_id, o.branch_id,
		o.is_stnk, o.stnk_price, o.matrix,
		%s AS branch,
		%s AS unit
	FROM orders AS o
	WHERE o.finance_id=$1 AND o.verified_by IS NOT NULL
	AND o.id NOT IN (
    SELECT order_id
    FROM invoice_details
		-- WHERE 0 = $2
    -- WHERE d.invoice_id = $2
)
	)
	
	SELECT t.is_selected, t.id, t.name, t.order_at, t.printed_at, t.bt_finance, t.bt_percent, t.bt_matel, t.ppn,
		t.nominal, t.subtotal, t.user_name, t.verified_by, t.validated_by, t.finance_id, t.branch_id,
		t.is_stnk, t.stnk_price, t.matrix,
		t.branch,
		t.unit
		FROM rs AS t
		ORDER BY t.is_selected DESC, t.finance_id, t.id DESC
	`,
		nestQuerySingle(queryBranch),
		nestQuerySingle(queryUnit),
		nestQuerySingle(queryBranch),
		nestQuerySingle(queryUnit),
	)

	rs, err := Sql().Query(sqlStatement, finance_id, invoice_id)

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
			&invoice.Ppn,
			&invoice.Nominal,
			&invoice.Subtotal,
			&invoice.UserName,
			&invoice.VerifiedBy,
			&invoice.ValidatedBy,
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
	var querFinance = `SELECT f.id, f.name, f.short_name "shortName", f.street, f.city, f.phone, f.cell, f.zip, f.email FROM finances f WHERE f.id = v.finance_id`
	var queryAccount = `SELECT c.id, c.name, c.type_id AS "typeId", c.descriptions, c.is_active AS "isActive", c.receivable_option AS "receivableOption", c.is_auto_debet AS "isAutoDebet" FROM acc_code c WHERE c.id = v.account_id`

	var sqlStatement = fmt.Sprintf(`SELECT v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman, v.finance_id, memo, total, tax, account_id, %s AS finance, %s AS account 
		FROM invoices AS v
		WHERE token @@ to_tsquery('indonesian', $1) 
		ORDER BY v.id DESC`,
		nestQuerySingle(querFinance),
		nestQuerySingle(queryAccount),
	)

	rs, err := Sql().Query(sqlStatement, txt)

	if err != nil {
		log.Fatalf("Unable to execute merks query %v", err)
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
			&item.Memo,
			&item.Total,
			&item.Tax,
			&item.AccountId,
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
	var querFinance = `SELECT f.id, f.name, f.short_name "shortName", f.street, f.city, f.phone, f.cell, f.zip, f.email FROM finances f WHERE f.id = v.finance_id`
	var queryAccount = `SELECT c.id, c.name, c.type_id AS "typeId", c.descriptions, c.is_active AS "isActive", c.receivable_option AS "receivableOption", c.is_auto_debet AS "isAutoDebet" FROM acc_code c WHERE c.id = v.account_id`

	var sqlStatement = fmt.Sprintf(`SELECT v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman, v.finance_id, memo, total, tax, account_id, %s AS finance, %s AS account 
		FROM invoices AS v
		WHERE EXTRACT(MONTH FROM v.invoice_at)=$1 AND EXTRACT(YEAR FROM v.invoice_at)=$2 OR 0=$1
		ORDER BY v.id DESC`,
		nestQuerySingle(querFinance),
		nestQuerySingle(queryAccount),
	)

	rs, err := Sql().Query(sqlStatement, month, year)

	if err != nil {
		log.Fatalf("Unable to execute merks query %v", err)
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
			&item.Memo,
			&item.Total,
			&item.Tax,
			&item.AccountId,
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
	var querFinance = `SELECT f.id, f.name, f.short_name "shortName", f.street, f.city, f.phone, f.cell, f.zip, f.email FROM finances f WHERE f.id = v.finance_id`
	var queryAccount = `SELECT c.id, c.name, c.type_id AS "typeId", c.descriptions, c.is_active AS "isActive", c.receivable_option AS "receivableOption", c.is_auto_debet AS "isAutoDebet" FROM acc_code c WHERE c.id = v.account_id`

	var sqlStatement = fmt.Sprintf(`SELECT v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman, v.finance_id, memo, total, tax, account_id, %s AS finance, %s AS account 
		FROM invoices AS v
		WHERE v.finance_id=$1 OR 0=$1
		ORDER BY v.id DESC`,
		nestQuerySingle(querFinance),
		nestQuerySingle(queryAccount),
	)

	rs, err := Sql().Query(sqlStatement, finance_id)

	if err != nil {
		log.Fatalf("Unable to execute merks query %v", err)
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
			&item.Memo,
			&item.Total,
			&item.Tax,
			&item.AccountId,
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
