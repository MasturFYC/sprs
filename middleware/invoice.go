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

func Invoice_GetAll(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	merks, err := invoice_get_all()

	if err != nil {
		log.Fatalf("Unable to get all merks. %v", err)
	}

	json.NewEncoder(w).Encode(&merks)
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

type invoice_item struct {
	models.Invoice
	Finance json.RawMessage `json:"finance,ommitempty"`
	Account json.RawMessage `json:"account,ommitempty"`
	Details json.RawMessage `json:"details,ommitempty"`
}

func invoice_get_item(id *int64) (invoice_item, error) {
	var item invoice_item

	var sqlStatement = `SELECT v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman, v.finance_id, memo, total, account_id,
(SELECT row_to_json(x) FROM (SELECT f.id, f.name, f.short_name, f.street, f.city, f.phone, f.cell, f.zip, f.email FROM finances f WHERE f.id = v.finance_id) x) AS finance,
(SELECT row_to_json(x) FROM (SELECT c.* FROM acc_code c WHERE c.id = v.account_id) x) AS account,
COALESCE((SELECT array_to_json(array_agg(row_to_json(x))) FROM (
SELECT d.invoice_id, d.id, d.order_id, d.price, d.tax, d.price - d.tax AS subtotal, (
SELECT row_to_json(x) FROM (
SELECT o.*, (
SELECT row_to_json(x) FROM (
SELECT u.*, (
SELECT row_to_json(x) FROM (
SELECT t.*,
(SELECT row_to_json(x) FROM (SELECT w.* FROM wheels w WHERE w.id = t.wheel_id) x) AS wheel,
(SELECT row_to_json(x) FROM (SELECT m.* FROM merks m WHERE m.id = t.merk_id) x) AS merk
FROM types t WHERE t.id = u.type_id)
x) AS type
FROM units u
WHERE u.order_id = o.id
) x) AS unit
FROM orders o
WHERE o.id = d.order_id
) x) AS spk	
FROM invoice_details d	
WHERE d.invoice_id = v.id) x), '[]') AS details
FROM invoices v
WHERE v.id=$1`

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
		&item.AccountId,
		&item.Finance,
		&item.Account,
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

func invoice_delete(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM invoice WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete finance. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func invoice_create(inv *models.Invoice) (int64, error) {

	sqlStatement := `INSERT INTO invoices 
	(invoice_at, payment_term, due_at, salesman, finance_id, memo, total, account_id) 
	VALUES 
	($1, $2, $3, $4, $5, $6, $7, $8)
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
		inv.AccountId,
	).Scan(&id)

	if err != nil {
		log.Printf("Unable to create finance. %v", err)
	}

	return id, err
}

func invoice_update(id *int64, inv *models.Invoice) (int64, error) {

	sqlStatement := `UPDATE invoice SET
	invoice_at=$1,
	payment_term=$2,
	due_at=$3,
	salesman=$4,
	finance_id=$5,
	memo=$6,
	total=$7,
	account_id=$8
	WHERE id=$9`

	res, err := Sql().Exec(sqlStatement,
		inv.InvoiceAt,
		inv.PaymentTerm,
		inv.DueAt,
		inv.Salesman,
		inv.FinanceID,
		inv.Memo,
		inv.Total,
		inv.AccountId,
		inv.AccountId,
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

func invoice_get_all() ([]invoice_item, error) {
	var invoices []invoice_item

	var sqlStatement = `SELECT v.id, v.invoice_at, v.payment_term, v.due_at, v.salesman, v.finance_id, memo, total, account_id,
(SELECT row_to_json(x) FROM (SELECT f.id, f.name, f.short_name, f.street, f.city, f.phone, f.cell, f.zip, f.email FROM finances f WHERE f.id = v.finance_id) x) AS finance,
(SELECT row_to_json(x) FROM (SELECT c.* FROM acc_code c WHERE c.id = v.account_id) x) AS account,
COALESCE((SELECT array_to_json(array_agg(row_to_json(x))) FROM (
SELECT d.invoice_id, d.id, d.order_id, d.price, d.tax, d.price - d.tax AS subtotal, (
SELECT row_to_json(x) FROM (
SELECT o.*, (
SELECT row_to_json(x) FROM (
SELECT u.*, (
SELECT row_to_json(x) FROM (
SELECT t.*,
(SELECT row_to_json(x) FROM (SELECT w.* FROM wheels w WHERE w.id = t.wheel_id) x) AS wheel,
(SELECT row_to_json(x) FROM (SELECT m.* FROM merks m WHERE m.id = t.merk_id) x) AS merk
FROM types t WHERE t.id = u.type_id)
x) AS type
FROM units u
WHERE u.order_id = o.id
) x) AS unit
FROM orders o
WHERE o.id = d.order_id
) x) AS spk	
FROM invoice_details d	
WHERE d.invoice_id = v.id) x), '[]') AS details
FROM invoices v
ORDER BY v.id`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute merks query %v", err)
	}

	defer rs.Close()

	for rs.Next() {
		var item invoice_item

		err := rs.Scan(
			&item.ID,
			&item.InvoiceAt,
			&item.PaymentTerm,
			&item.DueAt,
			&item.Salesman,
			&item.FinanceID,
			&item.Memo,
			&item.Total,
			&item.AccountId,
			&item.Finance,
			&item.Account,
			&item.Details,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		invoices = append(invoices, item)
	}

	return invoices, err
}
