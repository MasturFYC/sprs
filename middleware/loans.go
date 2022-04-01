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

type ts_loan_create struct {
	Loan  models.Loan `json:"loan"`
	Trx   models.Trx  `json:"trx"`
	Token string      `json:"token"`
}

func Loan_Create(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var loan ts_loan_create

	err := json.NewDecoder(r.Body).Decode(&loan)

	if err != nil {
		log.Fatalf("Fatal %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	newLoanId, err := loan_create(&loan.Loan)
	if err != nil {
		log.Fatalf("Fatal %v", err)
		//log.Fatalf("Nama finance tidak boleh sama.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	loan.Loan.ID = newLoanId
	loan.Trx.RefID = newLoanId
	trxid, err := createTransaction(&loan.Trx, loan.Token)

	if err != nil {
		log.Fatalf("Fatal %v", err)
		//log.Printf("(API) Unable to create transaction.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(loan.Trx.Details) > 0 {

		err = bulkInsertDetails(loan.Trx.Details, &trxid)

		if err != nil {
			log.Fatalf("Fatal %v", err)
			//log.Printf("Unable to insert transaction details.  %v", err)
			//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}

	loan.Trx.ID = trxid
	json.NewEncoder(w).Encode(&loan)

}

type ts_loan_payment struct {
	Trx   models.Trx `json:"trx"`
	Token string     `json:"token"`
}

func Loan_Payment(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	trxid, _ := strconv.ParseInt(params["id"], 10, 64)

	var data ts_loan_payment

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		log.Fatalf("Unable to decode trx from body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var updatedRows int64 = 0

	if trxid == 0 {
		trxid, err = createTransaction(&data.Trx, data.Token)
	} else {
		_, err = updateTransaction(&trxid, &data.Trx, data.Token)

	}

	if err != nil {
		//log.Printf("Unable to update transaction.  %v", err)
		log.Fatalf("Unable to update transaction %v", err)
		//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(data.Trx.Details) > 0 {

		_, err = deleteDetailsByOrder(&trxid)
		if err != nil {
			log.Fatalf("Unable to delete trx detail query %v", err)
			//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		//}

		// 	var newId int64 = 0

		err = bulkInsertDetails(data.Trx.Details, &trxid)

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
		ID:      trxid,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

func Loan_Update(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var loan ts_loan_create

	err := json.NewDecoder(r.Body).Decode(&loan)

	if err != nil {
		log.Fatalf("Unable to decode loan from body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = loan_update(&id, &loan.Loan)

	if err != nil {
		log.Fatalf("Loan update error.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	updatedRows, err := updateTransaction(&loan.Trx.ID, &loan.Trx, loan.Token)

	if err != nil {
		//log.Printf("Unable to update transaction.  %v", err)
		log.Fatalf("Unable to update transaction %v", err)
		//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if len(loan.Trx.Details) > 0 {

		_, err = deleteDetailsByOrder(&id)
		if err != nil {
			log.Fatalf("Unable to delete trx detail query %v", err)
			//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		//}

		// 	var newId int64 = 0

		err = bulkInsertDetails(loan.Trx.Details, &loan.Trx.ID)

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

type loan_item struct {
	models.Loan
	Trxs *json.RawMessage `json:"trxs,omitempty"`
}

func loan_get_item(id *int64) (loan_item, error) {

	var p loan_item
	sb := strings.Builder{}
	sbTrxDetail := strings.Builder{}
	sbTrx := strings.Builder{}

	sbTrxDetail.WriteString("WITH RECURSIVE rs AS (")
	sbTrxDetail.WriteString(" SELECT 1 as group, d.trx_id, d.id, d.code_id, d.debt, d.cred")
	sbTrxDetail.WriteString(" FROM trx_detail AS d")
	sbTrxDetail.WriteString(" INNER JOIN acc_code AS c ON c.id = d.code_id")
	sbTrxDetail.WriteString(" INNER JOIN acc_type AS e ON e.id = c.type_id")
	sbTrxDetail.WriteString(" WHERE e.group_id = 1")
	sbTrxDetail.WriteString(")\n")

	sbTrxDetail.WriteString(`SELECT rs.group AS "groupId", rs.id, rs.trx_id AS "trxId", rs.code_id AS "codeId", rs.debt, rs.cred`)
	sbTrxDetail.WriteString(", sum(rs.debt - rs.cred) OVER (ORDER BY rs.group, rs.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) as saldo")
	sbTrxDetail.WriteString(" FROM rs")
	sbTrxDetail.WriteString(" WHERE rs.trx_id = x.id")

	sbTrx.WriteString(`SELECT x.id, x.ref_id AS "refId", x.division, x.descriptions, x.trx_date AS "trxDate", x.memo`)
	sbTrx.WriteString(", ")
	sbTrx.WriteString(fyc.NestQuerySingle(sbTrxDetail.String()))
	sbTrx.WriteString(" AS detail")
	sbTrx.WriteString(" FROM trx AS x")
	sbTrx.WriteString(" WHERE x.ref_id = t.id AND (x.division ='trx-loan' OR x.division ='trx-angsuran')")
	sbTrx.WriteString(" ORDER BY x.id")

	sb.WriteString("SELECT")
	sb.WriteString(" t.id, t.name, t.street, t.city, t.phone, t.cell, t.zip, t.persen")
	sb.WriteString(",")
	sb.WriteString(fyc.NestQuery(sbTrx.String()))
	sb.WriteString(" AS trx")
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
		&p.Persen,
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

	// return empty on error
	return p, err
}

type loan_all struct {
	ID     int64             `json:"id"`
	Name   string            `json:"name"`
	Street models.NullString `json:"street,omitempty"`
	City   models.NullString `json:"city,omitempty"`
	Phone  models.NullString `json:"phone,omitempty"`
	Cell   models.NullString `json:"cell,omitempty"`
	Zip    models.NullString `json:"zip,omitempty"`
	Persen float32           `json:"persen"`

	TrxID        int64             `json:"trxId"`
	Division     string            `json:"division"`
	Descriptions models.NullString `json:"descriptions"`
	TrxDate      string            `json:"trxDate"`
	Memo         models.NullString `json:"memo"`

	Loan *json.RawMessage `json:"loan,omitempty"`
}

func loan_get_all() ([]loan_all, error) {

	var loans []loan_all

	sb := strings.Builder{}
	sb2 := strings.Builder{}

	sb2.WriteString("SELECT ln.id")
	sb2.WriteString(", sum(d.debt) as debt")
	sb2.WriteString(", sum(d.cred) as cred")
	sb2.WriteString(", sum(d.debt + (d.debt * (ln.persen / 100))) as piutang")
	sb2.WriteString(", sum((d.debt + (d.debt * (ln.persen / 100))) - d.cred) as saldo")
	sb2.WriteString(" FROM trx_detail AS d")
	sb2.WriteString(" INNER JOIN trx r ON r.id = d.trx_id")
	sb2.WriteString(" INNER JOIN loans ln ON ln.id = r.ref_id")
	sb2.WriteString(" INNER JOIN acc_code AS c ON c.id = d.code_id")
	sb2.WriteString(" INNER JOIN acc_type AS e ON e.id = c.type_id")
	sb2.WriteString(" WHERE e.group_id != 1 and ln.id = t.id AND (r.division = 'trx-loan' or r.division = 'trx-angsuran')")
	sb2.WriteString(" GROUP BY ln.id")

	sb.WriteString("SELECT t.id id, t.name, t.street, t.city, t.phone, t.cell, t.zip, t.persen,")
	sb.WriteString(" x.id as trx_id, x.division, x.descriptions, x.trx_date, x.memo, ")
	sb.WriteString(fyc.NestQuerySingle(sb2.String()))
	sb.WriteString(" AS details")
	sb.WriteString(" FROM loans AS t")
	sb.WriteString(" INNER JOIN trx AS x on x.ref_id = t.id AND x.division = 'trx-loan'")
	sb.WriteString(" ORDER BY x.trx_date, t.id")

	//log.Println(sb.String())

	rs, err := Sql().Query(sb.String())

	if err != nil {
		log.Fatalf("Unable to execute finances query %v", err)
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
			&p.Persen,
			&p.TrxID,
			&p.Division,
			&p.Descriptions,
			&p.TrxDate,
			&p.Memo,
			&p.Loan,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		loans = append(loans, p)
	}

	return loans, err
}

func loan_delete(id *int64) (int64, error) {

	//log.Printf("%d", id)
	sqlStatement := `DELETE FROM loans WHERE id=$1`
	_, err := Sql().Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	sqlStatement = `DELETE FROM trx WHERE ref_id=$1 AND (division='trx-loan' OR division='trx-angsuran')`
	res, err := Sql().Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}

func loan_create(loan *models.Loan) (int64, error) {

	sb := strings.Builder{}
	sb.WriteString("INSERT INTO loans")
	sb.WriteString(" (name, street, city, phone, cell, zip, persen)")
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
		loan.Persen,
	).Scan(&id)

	return id, err
}

func loan_update(id *int64, loan *models.Loan) (int64, error) {
	sb := strings.Builder{}
	sb.WriteString("UPDATE loans SET")
	sb.WriteString(" name=$2, street=$3, city=$4, phone=$5, cell=$6, zip=$7, persen=$8")
	sb.WriteString(" WHERE id=$1")

	res, err := Sql().Exec(sb.String(),
		id,
		loan.Name,
		loan.Street,
		loan.City,
		loan.Phone,
		loan.Cell,
		loan.Zip,
		loan.Persen,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}
