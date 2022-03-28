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

type loan_details struct {
	models.Loan
	Details *json.RawMessage `json:"details,omitempty"`
}

type loan_all struct {
	models.Loan
	Debt   float64 `json:"debt"`
	Cred   float64 `json:"cred"`
	Remain float64 `json:"remain"`
}

func Loan_GetAll(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	loans, err := loan_get_all()

	if err != nil {
		//		log.Fatalf("Unable to get all finances. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&loans)
}

func Loan_GetItem(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	loan, err := loan_get_item(&id)

	if err != nil {
		//log.Fatalf("Unable to get finance. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&loan)
}

func Loan_Delete(w http.ResponseWriter, r *http.Request) {
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

	deletedRows, err := loan_delete(&id)

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

func Loan_Create(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var loan models.Loan

	err := json.NewDecoder(r.Body).Decode(&loan)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := loan_create(&loan)

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

func Loan_Update(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var loan models.Loan

	err := json.NewDecoder(r.Body).Decode(&loan)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := loan_update(&id, &loan)

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

func loan_get_item(id *int64) (loan_details, error) {

	var p loan_details
	sb := strings.Builder{}
	sb2 := strings.Builder{}

	sb2.WriteString("WITH RECURSIVE rs AS (")
	sb2.WriteString(" SELECT loan_id, payment_at, id, descripts, debt, cred FROM loan_details WHERE loan_id=$1")
	sb2.WriteString(")")
	sb2.WriteString(" SELECT")
	sb2.WriteString(" rs.loan_id, rs.payment_at, rs.id, rs.descripts, rs.debt, rs.cred, ")
	sb2.WriteString(" sum(rs.debt - rs.cred) OVER (ORDER BY rs.payment_at, rs.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) as saldo")
	sb2.WriteString(" FROM rs")

	sb.WriteString("SELECT")
	sb.WriteString(" t.id, t.name, t.street, t.city, t.phone, t.cell, t.zip, t.descripts, ")
	sb.WriteString(NestQuery(sb2.String()))
	sb.WriteString(" AS details")
	sb.WriteString(" FROM loans AS t")
	sb.WriteString(" WHERE t.id=$1")

	rs := Sql().QueryRow(sb.String(), id)

	err := rs.Scan(
		&p.ID,
		&p.Name,
		&p.Street,
		&p.City,
		&p.Phone,
		&p.Cell,
		&p.Zip,
		&p.Descripts,
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

func loan_get_all() ([]loan_all, error) {

	var loans []loan_all

	sb := strings.Builder{}
	sb.WriteString("WITH RECURSIVE rs AS (")
	sb.WriteString(" select loan_id, sum(debt) debt, sum(cred) as cred FROM loan_details GROUP BY loan_id")
	sb.WriteString(")\n")
	sb.WriteString("SELECT t.id, t.name, t.street, t.city, t.phone, t.cell, t.zip, t.descripts, ")
	sb.WriteString(" r.debt, r.cred, r.debt - r.cred AS remain")
	sb.WriteString(" FROM loans AS t")
	sb.WriteString(" LEFT JOIN rs AS r ON r.loan_id = t.id")
	sb.WriteString(" ORDER BY t.name")

	rs, err := Sql().Query(sb.String())

	if err != nil {
		// log.Fatalf("Unable to execute finances query %v", err)
		return loans, err
	}

	defer rs.Close()

	for rs.Next() {
		var p loan_all

		err := rs.Scan(
			&p.ID,
			&p.Name,
			&p.Street,
			&p.City,
			&p.Phone,
			&p.Cell,
			&p.Zip,
			&p.Descripts,
			&p.Debt,
			&p.Cred,
			&p.Remain,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		loans = append(loans, p)
	}

	return loans, err
}

func loan_delete(id *int64) (int64, error) {
	sqlStatement := `DELETE FROM loans WHERE id=$1`
	res, err := Sql().Exec(sqlStatement, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}

func loan_create(loan *models.Loan) (int64, error) {

	sb := strings.Builder{}
	sb.WriteString("INSERT INTO loans")
	sb.WriteString(" (name, street, city, phone, cell, zip, descripts)")
	sb.WriteString(" VALUES")
	sb.WriteString(" ($1, $2, $3, $4, $5, $6, $7)")
	sb.WriteString(" RETURNING id")

	var id int64

	err := Sql().QueryRow(sb.String(),
		loan.Name,
		loan.Street,
		loan.City,
		loan.Phone,
		loan.Cell,
		loan.Zip,
		loan.Descripts,
	).Scan(&id)

	return id, err
}

func loan_update(id *int64, loan *models.Loan) (int64, error) {
	sb := strings.Builder{}
	sb.WriteString("UPDATE INTO SET")
	sb.WriteString(" name=$2, street=$3, city=$4, phone=$5, cell=$6, zip=$7, descripts=$8")
	sb.WriteString(" WHERE id=$1")

	res, err := Sql().Exec(sb.String(),
		id,
		loan.Name,
		loan.Street,
		loan.City,
		loan.Phone,
		loan.Cell,
		loan.Zip,
		loan.Descripts,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}
