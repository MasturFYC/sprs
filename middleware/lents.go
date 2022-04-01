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
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = lent_create(&data.Lent)

	if err != nil {
		log.Fatalf("Nama peminjam tidak boleh sama.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	trxid, err := createTransaction(&data.Trx, data.Token)

	if err != nil {
		log.Fatalf("Fatal %v", err)
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

	var data ts_lent_create

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		log.Fatalf("Unable to decode loan from body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = lent_update(&id, &data.Lent)

	if err != nil {
		log.Fatalf("Loan update error.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	updatedRows, err := updateTransaction(&data.Trx.ID, &data.Trx, data.Token)

	if err != nil {
		//log.Printf("Unable to update transaction.  %v", err)
		log.Fatalf("Unable to update transaction %v", err)
		//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(data.Trx.Details) > 0 {

		_, err = deleteDetailsByOrder(&id)
		if err != nil {
			log.Fatalf("Unable to delete trx detail query %v", err)
			//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		//}

		// 	var newId int64 = 0

		err = bulkInsertDetails(data.Trx.Details, &data.Trx.ID)

		if err != nil {
			log.Fatalf("Unable to execute finances query %v", err)
			//log.Printf("Unable to insert transaction details (message from command).  %v", err)
			//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}

	msg := fmt.Sprintf("Loan updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func create_unit_query() *strings.Builder {
	b := strings.Builder{}

	b.WriteString("SELECT")
	b.WriteString(` o.id, o.name, o.order_at as "orderAt", o.bt_finance AS "btFinance"`)
	b.WriteString(`, o.bt_percent AS "btPercent", o.bt_matel AS "btMatel",`)
	b.WriteString(`, u.nopol, u.year`)
	b.WriteString(`, e.name type"`)
	b.WriteString(`, w.short_name wheel`)
	b.WriteString(`, m.name merk"`)
	b.WriteString(" FROM orders as o")
	b.WriteString(" INNER JOIN units AS u on u.order_id = o.id")
	b.WriteString(" INNER JOIN types AS e on e.id = u.type_id")
	b.WriteString(" INNER JOIN wheels AS w on w.id = e.wheel_id")
	b.WriteString(" INNER JOIN merks AS m on m.id = e.merk_id")

	return &b
}

type ts_lent_item struct {
	models.Lent
	Unit *json.RawMessage `json:"unit,omitempty"`
	Trxs *json.RawMessage `json:"trxs,omitempty"`
}

func lent_get_item(order_id *int64) (ts_lent_item, error) {
	var p ts_lent_item
	sb := strings.Builder{}

	sb2 := create_unit_query()
	sb2.WriteString(" WHERE o.id = $1")

	sbTrxDetail := strings.Builder{}
	sbTrx := strings.Builder{}

	sbTrxDetail.WriteString("WITH RECURSIVE rs AS (")
	sbTrxDetail.WriteString(" SELECT d.trx_id, d.id, d.code_id, d.debt, d.cred")
	sbTrxDetail.WriteString(" FROM trx_detail AS d")
	sbTrxDetail.WriteString(" INNER JOIN acc_code AS c ON c.id = d.code_id")
	sbTrxDetail.WriteString(" INNER JOIN acc_type AS e ON e.id = c.type_id")
	sbTrxDetail.WriteString(" WHERE e.group_id = 1")
	sbTrxDetail.WriteString(")\n")

	sbTrxDetail.WriteString(`SELECT rs.trx_id AS "trxId", rs.id, rs.code_id AS "codeId", rs.debt, rs.cred`)
	sbTrxDetail.WriteString(", sum(rs.debt - rs.cred) OVER (ORDER BY rs.trx_id, rs.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) as saldo")
	sbTrxDetail.WriteString(" FROM rs")
	sbTrxDetail.WriteString(" WHERE rs.trx_id = x.id")

	sbTrx.WriteString(`SELECT x.id, x.ref_id AS "refId", x.division, x.descriptions, x.trx_date AS "trxDate", x.memo`)
	sbTrx.WriteString(", ")
	sbTrx.WriteString(fyc.NestQuerySingle(sbTrxDetail.String()))
	sbTrx.WriteString(" AS detail")
	sbTrx.WriteString(" FROM trx AS x")
	sbTrx.WriteString(" WHERE x.ref_id = $1 AND (x.division ='trx-lent' OR x.division ='trx-cicilan')")
	sbTrx.WriteString(" ORDER BY x.id")

	sb.WriteString("SELECT")
	sb.WriteString(" t.order_id, t.name, t.street, t.city, t.phone, t.cell, t.zip, t.descripts, ")
	sb.WriteString(fyc.NestQuerySingle(sb2.String()))
	sb.WriteString(" AS unit, ")
	sb.WriteString(fyc.NestQuery(sbTrx.String()))
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
		&p.Descripts,
		&p.Unit,
		&p.Trxs,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return p, nil
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

func lent_get_all() ([]ts_lent_all, error) {

	var lents []ts_lent_all

	sbUnit := create_unit_query()
	sb := strings.Builder{}
	sb2 := strings.Builder{}

	sb2.WriteString(`SELECT ln.order_id as "orderId"`)
	sb2.WriteString(", sum(d.debt) as debt")
	sb2.WriteString(", sum(d.cred) as cred")
	sb2.WriteString(", sum(d.debt - d.cred) as saldo")
	sb2.WriteString(" FROM trx_detail AS d")
	sb2.WriteString(" INNER JOIN trx r ON r.id = d.trx_id")
	sb2.WriteString(" INNER JOIN lents ln ON ln.order_id = r.ref_id")
	sb2.WriteString(" INNER JOIN acc_code AS c ON c.id = d.code_id")
	sb2.WriteString(" INNER JOIN acc_type AS e ON e.id = c.type_id")
	sb2.WriteString(" WHERE e.group_id != 1 and ln.order_id = o.id AND (r.division = 'trx-lent' or r.division = 'trx-cicilan')")
	sb2.WriteString(" GROUP BY ln.order_id")

	sb.WriteString("SELECT t.order_id, t.name, t.street, t.city, t.phone, t.cell, t.zip, t.descripts, ")
	sb.WriteString(fyc.NestQuerySingle(sb2.String()))
	sb.WriteString(" AS payment, ")
	sb.WriteString(fyc.NestQuerySingle(sbUnit.String()))
	sb.WriteString(" AS unit ")
	sb.WriteString(" FROM lents AS t")
	sb.WriteString(" INNER JOIN orders AS o ON o.id = t.order_id")
	sb.WriteString(" ORDER BY o.order_at")

	rs, err := Sql().Query(sb.String())

	if err != nil {
		log.Fatalf("Unable to execute finances query %v", err)
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
			&p.Descripts,
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
	sb.WriteString(" (order_id, name, street, city, phone, cell, zip, descripts)")
	sb.WriteString(" VALUES")
	sb.WriteString(" ($1, $2, $3, $4, $5, $6, $7, $8)")

	rs, err := Sql().Exec(sb.String(),
		lent.Name,
		lent.Street,
		lent.City,
		lent.Phone,
		lent.Cell,
		lent.Zip,
		lent.Descripts,
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
	sb.WriteString(" name=$2, street=$4, city=$5, phone=$6, cell=$7, zip=$8, descripts=$9")
	sb.WriteString(" WHERE order_id=$1")

	rs, err := Sql().Exec(sb.String(),
		id,
		lent.Name,
		lent.Street,
		lent.City,
		lent.Phone,
		lent.Cell,
		lent.Zip,
		lent.Descripts,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := rs.RowsAffected()
	return rowsAffected, err
}
