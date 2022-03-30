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

type lent_details struct {
	models.Lent
	Details *json.RawMessage `json:"details"`
}

type lent_all struct {
	models.Lent
	TotalPayment  float64 `json:"totalPayment"`
	RemainPayment float64 `json:"remainPayment"`
}

func Lent_GetAll(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	lents, err := lent_get_all()

	if err != nil {
		//		log.Fatalf("Unable to get all finances. %v", err)
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

func Lent_Create(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var loan models.Lent

	err := json.NewDecoder(r.Body).Decode(&loan)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := lent_create(&loan)

	if err != nil {
		//log.Fatalf("Nama finance tidak boleh sama.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	res := Response{
		ID:      id,
		Message: "Loan was created succesfully",
	}

	json.NewEncoder(w).Encode(&res)

}

func Lent_Update(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var loan models.Lent

	err := json.NewDecoder(r.Body).Decode(&loan)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := lent_update(&id, &loan)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Loand updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func lent_get_item(order_id *int64) (lent_details, error) {

	var p lent_details
	sb := strings.Builder{}

	sb.WriteString("WITH RECURSIVE rs AS (")
	sb.WriteString(" select order_id, id, payment_at, debt, descripts, cash_id FROM lent_details WHERE order_id=$1")
	sb.WriteString(")\n")
	sb.WriteString("SELECT")
	sb.WriteString(` t2.order_id as "OrderId", t2.id, t2.payment_at as "paymentAt", t2.debt, t2.descripts, t2.cash_id As "cashId"`)
	sb.WriteString(" t3.bt_matel - t2.debt")
	sb.WriteString(" OVER (ORDER BY t.payment_at, t.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) as saldo")
	sb.WriteString(" FROM rs as t2")
	sb.WriteString(" INNER JOIN orders AS t3 ON t3.id = t2.order_id")

	sb.WriteString("SELECT")
	sb.WriteString(" t.order_id, t.name, t.descripts, t.street, t.city, t.phone, t.cell, t.zip, ")
	sb.WriteString(fyc.NestQuery(`SELECT order_id AS "orderId", id, payment_at AS "paymentAt", debt, descripts, cash_id As "cashId" FROM lent_details WHERE order_id = t.order_id`))
	sb.WriteString(" AS details")
	sb.WriteString(" FROM lents AS t")
	sb.WriteString(" LEFT JOIN rs AS r ON r.order_id = t.order_id")
	sb.WriteString(" WHERE t.order_id=$1")

	rs := Sql().QueryRow(sb.String(), order_id)

	err := rs.Scan(
		&p.OrderID,
		&p.Name,
		&p.Descripts,
		&p.Street,
		&p.City,
		&p.Phone,
		&p.Cell,
		&p.Zip,
		&p.Details,
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

func lent_get_all() ([]lent_all, error) {

	var lents []lent_all

	sb := strings.Builder{}
	sb.WriteString("WITH RECURSIVE rs AS (")
	sb.WriteString(" select order_id, sum(debt) debt FROM lent_details GROUP BY order_id")
	sb.WriteString(")\n")
	sb.WriteString("SELECT t.order_id, t.name, t.descripts, t.street, t.city, t.phone, t.cell, t.zip,")
	sb.WriteString(" COALESCE(r.debt,0) AS total_payment, o.bt_matel - COALESCE(r.debt,0) AS remain_payment")
	sb.WriteString(" FROM lents AS t")
	sb.WriteString(" INNER JOIN orders AS o ON o.id = t.order_id")
	sb.WriteString(" LEFT JOIN rs AS r ON r.order_id = t.order_id")
	sb.WriteString(" ORDER BY t.name")

	rs, err := Sql().Query(sb.String())

	if err != nil {
		// log.Fatalf("Unable to execute finances query %v", err)
		return lents, err
	}

	defer rs.Close()

	for rs.Next() {
		var p lent_all

		err := rs.Scan(
			&p.OrderID,
			&p.Name,
			&p.Descripts,
			&p.Street,
			&p.City,
			&p.Phone,
			&p.Cell,
			&p.Zip,
			&p.TotalPayment,
			&p.RemainPayment,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		lents = append(lents, p)
	}

	return lents, err
}

func lent_delete(id *int64) (int64, error) {
	sqlStatement := `DELETE FROM lents WHERE id=$1`
	res, err := Sql().Exec(sqlStatement, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}

func lent_create(lent *models.Lent) (int64, error) {

	sb := strings.Builder{}
	sb.WriteString("INSERT INTO lents")
	sb.WriteString(" (order_id, name, descripts, street, city, phone, cell, zip)")
	sb.WriteString(" VALUES")
	sb.WriteString(" ($1, $2, $3, $4, $5, $6, $7, $8)")

	rs, err := Sql().Exec(sb.String(),
		lent.Name,
		lent.Descripts,
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

func lent_update(id *int64, lent *models.Lent) (int64, error) {
	sb := strings.Builder{}
	sb.WriteString("UPDATE lents SET")
	sb.WriteString(" name=$2, descripts=$3, street=$4, city=$5, phone=$6, cell=$7, zip=$8")
	sb.WriteString(" WHERE order_id=$1")

	rs, err := Sql().Exec(sb.String(),
		id,
		lent.Name,
		lent.Descripts,
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
