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

func Lent_GetUnits(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	units, err := lent_get_units()

	if err != nil {
		//log.Fatalf("Unable to get all lent. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&units)
}

func Lent_GetAll(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	lents, err := lent_get_all()

	if err != nil {
		//log.Fatalf("Unable to get all lent. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&lents)
}

func Lent_GetItem(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	loan, err := lent_get_item(&id)

	if err != nil {
		//log.Fatalf("Unable to get finance. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&loan)
}

func Lent_Delete(w http.ResponseWriter, r *http.Request) {
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

	deletedRows, err := lent_delete(&id)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("loan deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      id,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

type ts_lent_create struct {
	Lent  models.Lent `json:"lent"`
	Trx   models.Trx  `json:"trx"`
	Token string      `json:"token"`
}

func Lent_Create(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var data ts_lent_create

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = lent_create(&data.Lent)

	if err != nil {
		log.Fatalf("Pinjaman tidak bisa disimpan.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	trxid, err := createTransaction(&data.Trx, data.Token)

	if err != nil {
		log.Fatalf("Tidak bsia membuat transaksi baru %v", err)
		//log.Printf("(API) Unable to create transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(data.Trx.Details) > 0 {

		err = bulkInsertDetails(data.Trx.Details, &trxid)

		if err != nil {
			log.Fatalf("Fatal %v", err)
			return
		}
	}

	data.Trx.ID = trxid
	json.NewEncoder(w).Encode(&data)

}

func Lent_Update(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var data models.Lent

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		log.Fatalf("Unable to decode lent from body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := lent_update(&id, &data)

	log.Printf("\n\n%v\n\n", data)

	if err != nil {
		log.Fatalf("Loan update error.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// updatedRows, err := updateTransaction(&data.Trx.ID, &data.Trx, data.Token)

	// if err != nil {
	// 	//log.Printf("Unable to update transaction.  %v", err)
	// 	log.Fatalf("Unable to update transaction %v", err)
	// 	//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 	return
	// }

	// if len(data.Trx.Details) > 0 {

	// 	_, err = deleteDetailsByOrder(&id)
	// 	if err != nil {
	// 		log.Fatalf("Unable to delete trx detail query %v", err)
	// 		//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 		return
	// 	}
	// 	//}

	// 	// 	var newId int64 = 0

	// 	err = bulkInsertDetails(data.Trx.Details, &data.Trx.ID)

	// 	if err != nil {
	// 		log.Fatalf("Unable to insert trx details %v", err)
	// 		//log.Printf("Unable to insert transaction details (message from command).  %v", err)
	// 		//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 		return
	// 	}
	// }

	msg := fmt.Sprintf("Lent updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func lent_create_unit_query() *strings.Builder {
	b := strings.Builder{}

	b.WriteString("SELECT")
	b.WriteString(` o.id, o.name, o.order_at as "orderAt", o.bt_finance AS "btFinance"`)
	b.WriteString(`, o.bt_percent AS "btPercent", o.bt_matel AS "btMatel"`)
	b.WriteString(", u.nopol, u.year")
	b.WriteString(`, e.name AS "type"`)
	b.WriteString(", w.short_name AS wheel")
	b.WriteString(", m.name AS merk")
	b.WriteString(" FROM orders as o")
	b.WriteString(" INNER JOIN units AS u on u.order_id = o.id")
	b.WriteString(" INNER JOIN types AS e on e.id = u.type_id")
	b.WriteString(" INNER JOIN wheels AS w on w.id = e.wheel_id")
	b.WriteString(" INNER JOIN merks AS m on m.id = e.merk_id")

	return &b
}

func lent_create_trx_detail_query() *strings.Builder {
	b := strings.Builder{}

	b.WriteString(`SELECT d.trx_id AS "trxId", d.id, d.code_id AS "codeId", d.debt, d.cred`)
	b.WriteString(", sum(d.debt - d.cred) OVER (ORDER BY d.trx_id, d.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) as saldo")
	b.WriteString(" FROM trx_detail AS d")
	b.WriteString(" INNER JOIN acc_code AS c ON c.id = d.code_id")
	b.WriteString(" INNER JOIN acc_type AS e ON e.id = c.type_id")
	b.WriteString(" WHERE d.trx_id=x.id")
	b.WriteString(" AND e.group_id = 1 ")
	//	b.WriteString(" AND (t1.division='trx-lent' AND t1.division='trx-cicilan')")
	b.WriteString(" ORDER BY d.trx_id, d.id")

	sb := strings.Builder{}

	sb.WriteString(`SELECT x.id, x.ref_id AS "refId", x.division, x.trx_date AS "trxDate", x.descriptions, x.memo,`)
	sb.WriteString(fyc.NestQuerySingle(b.String()))
	sb.WriteString(" AS detail ")
	sb.WriteString(" FROM trx x")
	sb.WriteString(" WHERE x.ref_id=$1")

	return &sb
}

type ts_lent_item struct {
	models.Lent
	Unit *json.RawMessage `json:"unit,omitempty"`
	Trxs *json.RawMessage `json:"trxs,omitempty"`
}

func lent_get_item(order_id *int64) (ts_lent_item, error) {
	var p ts_lent_item

	qunit := lent_create_unit_query()
	qunit.WriteString(" WHERE o.id = $1")

	qtrx := lent_create_trx_detail_query()

	sb := strings.Builder{}
	sb.WriteString("SELECT")
	sb.WriteString(" t.order_id, t.name, t.street, t.city, t.phone, t.cell, t.zip, ")
	sb.WriteString(fyc.NestQuerySingle(qunit.String()))
	sb.WriteString(" AS unit, ")
	sb.WriteString(fyc.NestQuery(qtrx.String()))
	sb.WriteString(" AS trxs ")
	sb.WriteString(" FROM lents AS t")
	sb.WriteString(" WHERE t.order_id=$1")

	rs := Sql().QueryRow(sb.String(), order_id)

	err := rs.Scan(
		&p.OrderID,
		&p.Name,
		&p.Street,
		&p.City,
		&p.Phone,
		&p.Cell,
		&p.Zip,
		&p.Unit,
		&p.Trxs,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return p, err
	case nil:
		return p, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return p, err
}

type ts_lent_all struct {
	models.Lent
	Payment *json.RawMessage `json:"payment"`
	Unit    *json.RawMessage `json:"unit"`
}

type lent_all_unit struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	OrderAt   string  `json:"orderAt"`
	BtFinance float64 `json:"btFinance"`
	BtPercent float32 `json:"btPercent"`
	BtMatel   float64 `json:"btMatel"`
	Nopol     string  `json:"nopol"`
	Year      int     `json:"year"`
	Type      string  `json:"type"`
	Wheel     string  `json:"wheel"`
	Merk      string  `json:"merk"`
}

func lent_get_units() ([]lent_all_unit, error) {

	var units []lent_all_unit

	qunit := lent_create_unit_query()
	qunit.WriteString(" WHERE o.verified_by IS NULL")
	qunit.WriteString(" AND o.id NOT IN (SELECT order_id FROM invoice_details)")

	rs, err := Sql().Query(qunit.String())

	if err != nil {
		log.Fatalf("Unable to execute finances query %v", err)
		return units, err
	}

	defer rs.Close()

	for rs.Next() {
		var p lent_all_unit

		err := rs.Scan(
			&p.ID,
			&p.Name,
			&p.OrderAt,
			&p.BtFinance,
			&p.BtPercent,
			&p.BtMatel,
			&p.Nopol,
			&p.Year,
			&p.Type,
			&p.Wheel,
			&p.Merk,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		units = append(units, p)
	}

	return units, err

}

func lent_get_all() ([]ts_lent_all, error) {

	var lents []ts_lent_all

	qunit := lent_create_unit_query()
	qunit.WriteString(" WHERE o.id = t.order_id")
	sb := strings.Builder{}
	sb2 := strings.Builder{}

	sb2.WriteString(`SELECT ln.order_id as "orderId"`)
	sb2.WriteString(", sum(d.debt) as debt")
	sb2.WriteString(", sum(d.cred) as cred")
	sb2.WriteString(", sum(t3.bt_finance) as piutang")
	sb2.WriteString(", sum(t3.bt_finance - d.cred) as saldo")
	sb2.WriteString(" FROM trx_detail AS d")
	sb2.WriteString(" INNER JOIN trx r ON r.id = d.trx_id")
	sb2.WriteString(" INNER JOIN lents ln ON ln.order_id = r.ref_id")
	sb2.WriteString(" INNER JOIN orders t3 ON t3.id = ln.order_id")
	sb2.WriteString(" INNER JOIN acc_code AS c ON c.id = d.code_id")
	sb2.WriteString(" INNER JOIN acc_type AS e ON e.id = c.type_id")
	sb2.WriteString(" WHERE e.group_id != 1 and ln.order_id = o.id AND (r.division = 'trx-lent' or r.division = 'trx-cicilan')")
	sb2.WriteString(" GROUP BY ln.order_id")

	sb.WriteString("SELECT t.order_id, t.name, t.street, t.city, t.phone, t.cell, t.zip, ")
	sb.WriteString(fyc.NestQuerySingle(sb2.String()))
	sb.WriteString(" AS payment, ")
	sb.WriteString(fyc.NestQuerySingle(qunit.String()))
	sb.WriteString(" AS unit ")
	sb.WriteString(" FROM lents AS t")
	sb.WriteString(" INNER JOIN orders AS o ON o.id = t.order_id")
	sb.WriteString(" ORDER BY o.order_at")

	rs, err := Sql().Query(sb.String())

	if err != nil {
		log.Fatalf("Unable to execute lents query %v", err)
		return lents, err
	}

	defer rs.Close()

	for rs.Next() {
		var p ts_lent_all

		err := rs.Scan(
			&p.OrderID,
			&p.Name,
			&p.Street,
			&p.City,
			&p.Phone,
			&p.Cell,
			&p.Zip,
			&p.Payment,
			&p.Unit,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		lents = append(lents, p)
	}

	return lents, err
}

func lent_delete(id *int64) (int64, error) {

	//log.Printf("%d", id)
	sqlStatement := `DELETE FROM lents WHERE id=$1`
	_, err := Sql().Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	sqlStatement = `DELETE FROM trx WHERE ref_id=$1 AND (division='trx-lent' OR division='trx-cicilan')`
	res, err := Sql().Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err

}

func lent_create(lent *models.Lent) (int64, error) {

	sb := strings.Builder{}
	sb.WriteString("INSERT INTO lents")
	sb.WriteString(" (order_id, name, street, city, phone, cell, zip)")
	sb.WriteString(" VALUES")
	sb.WriteString(" ($1, $2, $3, $4, $5, $6, $7)")

	rs, err := Sql().Exec(sb.String(),
		lent.OrderID,
		lent.Name,
		lent.Street,
		lent.City,
		lent.Phone,
		lent.Cell,
		lent.Zip,
	)

	if err != nil {
		log.Fatalf("Error %v", err)
		return 0, err
	}

	rowsAffected, err := rs.RowsAffected()
	return rowsAffected, err
}

func lent_update(id *int64, lent *models.Lent) (int64, error) {
	sb := strings.Builder{}
	sb.WriteString("UPDATE lents SET")
	sb.WriteString(" name=$2, street=$3, city=$4, phone=$5, cell=$6, zip=$7")
	sb.WriteString(" WHERE order_id=$1")

	rs, err := Sql().Exec(sb.String(),
		id,
		lent.Name,
		lent.Street,
		lent.City,
		lent.Phone,
		lent.Cell,
		lent.Zip,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := rs.RowsAffected()
	return rowsAffected, err
}