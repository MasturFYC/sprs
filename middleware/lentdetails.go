package middleware

import (
	"encoding/json"
	"fmt"
	"strings"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func LentDetail_Delete(w http.ResponseWriter, r *http.Request) {
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

	deletedRows, err := lent_detail_delete(&id)

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

func LentDetail_Create(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var loan models.LentDetail

	err := json.NewDecoder(r.Body).Decode(&loan)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := lent_detail_create(&loan)

	if err != nil {
		//log.Fatalf("Nama finance tidak boleh sama.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	res := Response{
		ID:      id,
		Message: "Lent was created succesfully",
	}

	json.NewEncoder(w).Encode(&res)

}

func LentDetail_Update(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var loan models.LentDetail

	err := json.NewDecoder(r.Body).Decode(&loan)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := lent_detail_update(&id, &loan)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Lent details updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func lent_detail_delete(id *int64) (int64, error) {
	sqlStatement := `DELETE FROM lent_details WHERE id=$1`
	res, err := Sql().Exec(sqlStatement, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}

func lent_detail_create(loan *models.LentDetail) (int64, error) {

	sb := strings.Builder{}
	sb.WriteString("INSERT INTO lent_details")
	sb.WriteString(" (order_id, payment_at, descripts, debt, cred, cash_id)")
	sb.WriteString(" VALUES")
	sb.WriteString(" ($1, $2, $3, $4, $5, $6)")
	sb.WriteString(" RETURNING id")

	var id int64

	err := Sql().QueryRow(sb.String(),
		loan.OrderID,
		loan.PaymentAt,
		loan.Descripts,
		loan.Debt,
		loan.Cred,
		loan.CashId,
	).Scan(&id)

	return id, err
}

func lent_detail_update(id *int64, loan *models.LentDetail) (int64, error) {
	sb := strings.Builder{}
	sb.WriteString("UPDATE lent_details SET")
	sb.WriteString(" order_id=$2, payment_at=$3, descripts=$4, debt=$5, cred=$6, cash_id=$7")
	sb.WriteString(" WHERE id=$1")

	res, err := Sql().Exec(sb.String(),
		id,
		loan.OrderID,
		loan.PaymentAt,
		loan.Descripts,
		loan.Debt,
		loan.Cred,
		loan.CashId,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}
