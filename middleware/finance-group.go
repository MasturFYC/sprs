package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"fyc.com/sprs/models"
	"github.com/MasturFYC/fyc"

	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func FinanceGroup_GetFinances(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	fgs, err := fg_get_finances()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&fgs)
}

func FinanceGroup_GetAll(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	fgs, err := get_all_finance_groups()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&fgs)
}

func FinanceGroup_GetItem(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	finances, err := get_finance_group(&id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&finances)
}

func FinanceGroup_Delete(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows, err := delete_finance_group(&id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Finance group deleted successfully. Total rows/record affected %v", deletedRows)

	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func FinanceGroup_Create(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var fg models.FinanceGroup

	err := json.NewDecoder(r.Body).Decode(&fg)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := create_finance_group(&fg)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fg.ID = id

	json.NewEncoder(w).Encode(&fg)

}

func FinanceGroup_Update(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var fg models.FinanceGroup

	err := json.NewDecoder(r.Body).Decode(&fg)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := update_finance_group(&id, &fg)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Finance group updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func get_finance_group(id *int) (models.FinanceGroup, error) {

	var fg models.FinanceGroup

	var sqlStatement = `SELECT id, name	FROM finance_groups	WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&fg.ID, &fg.Name)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return fg, nil
	case nil:
		return fg, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return fg, err
}

type FgFinances struct {
	models.FinanceGroup
	Finances json.RawMessage `json:"finances,omitempty"`
}

func fg_get_finances() ([]FgFinances, error) {
	var fgs []FgFinances
	var querFinance = `SELECT id, name, short_name AS "shortName" FROM finances WHERE group_id = g.id ORDER BY name`

	var sqlStatement = fmt.Sprintf("SELECT g.id, g.name, %s AS finances	FROM finance_groups AS g ORDER BY g.name",
		fyc.NestQuery(querFinance),
	)

	//	log.Print(sqlStatement)
	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		//log.Fatalf("Unable to execute merks query %v", err)
		return fgs, err
	}

	defer rs.Close()

	for rs.Next() {
		var item FgFinances

		err := rs.Scan(
			&item.ID,
			&item.Name,
			&item.Finances,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		fgs = append(fgs, item)
	}

	return fgs, err
}

func get_all_finance_groups() ([]models.FinanceGroup, error) {

	var fgs []models.FinanceGroup

	var sqlStatement = `SELECT id, name FROM finance_groups ORDER BY name`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		return fgs, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.FinanceGroup

		err := rs.Scan(&p.ID, &p.Name)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		fgs = append(fgs, p)
	}

	return fgs, err
}

func delete_finance_group(id *int) (int64, error) {

	sqlStatement := `DELETE FROM finance_groups WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func create_finance_group(fg *models.FinanceGroup) (int, error) {

	sqlStatement := `INSERT INTO finance_groups (name) VALUES ($1) RETURNING id`

	var id int

	err := Sql().QueryRow(sqlStatement, fg.Name).Scan(&id)

	return id, err
}

func update_finance_group(id *int, fg *models.FinanceGroup) (int64, error) {

	sqlStatement := `UPDATE finance_groups SET name=$2 WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		fg.Name,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}
