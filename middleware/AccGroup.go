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

type all_accounts struct {
	models.AccCode
	IsGroup   bool `json:"isGroup"`
	IsType    bool `json:"isType"`
	IsAccount bool `json:"isAccount"`
}

func Group_GetAllAccount(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	accounts, err := get_all_accounts()

	if err != nil {
		//log.Printf("Unable to get all account groups. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&accounts)
}

func GetAccGroups(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	groups, err := getAllAccGroups()

	if err != nil {
		//log.Printf("Unable to get all account groups. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&groups)
}
func Group_GetTypes(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_group, err := group_get_types(&id)

	if err != nil {
		//log.Printf("Unable to get account group. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_group)
}

/*


type Acc_Type struct {
	GroupID      int32             `json:"groupId"`
	ID           int32             `json:"id"`
	Name         string            `json:"name"`
	Descriptions models.NullString `json:"descriptions"`
	Accounts     json.RawMessage   `json:"accounts"`
}


	var sqlStatement = `SELECT t.group_id, t.id, t.name, t.descriptions,
	coalesce((SELECT array_to_json(array_agg(row_to_json(x)))
        FROM (select c.id, c.name from acc_code c where c.type_id = t.id) x),
      '[]') AS accounts
	FROM acc_type t
	WHERE t.group_id=$1 OR 0=$1
	ORDER BY t.id`
*/

func group_get_types(id *int) ([]models.AccType, error) {
	var results []models.AccType

	var sqlStatement = `SELECT group_id, id, name, descriptions FROM acc_type
	WHERE group_id=$1 OR 0 = $1
	ORDER BY id`

	rs, err := Sql().Query(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to execute account type query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccType

		err := rs.Scan(&p.GroupID, &p.ID, &p.Name, &p.Descriptions)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func GetAccGroup(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	acc_group, err := getAccGroup(&id)

	if err != nil {
		//log.Printf("Unable to get account group. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&acc_group)
}

func CreateAccGroup(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var acc_group models.AccGroup

	err := json.NewDecoder(r.Body).Decode(&acc_group)

	if err != nil {
		//log.Printf("Unable to decode the request body to account group.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rowsAffected, err := createAccGroup(&acc_group)

	if err != nil {
		//log.Printf("(API) Unable to create account group.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	res := Response{
		ID:      rowsAffected,
		Message: "Account group created successfully",
	}

	json.NewEncoder(w).Encode(&res)

}

func UpdateAccGroup(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var acc_group models.AccGroup

	err := json.NewDecoder(r.Body).Decode(&acc_group)

	if err != nil {
		//log.Printf("Unable to decode the request body to account group.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedRows, err := updateAccGroup(&id, &acc_group)

	if err != nil {
		//log.Printf("Unable to update account group.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Account group updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func DeleteAccGroup(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows, err := deleteAccGroup(&id)

	if err != nil {
		//log.Printf("Unable to delete account group.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Account group deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      deletedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getAccGroup(id *int) (models.AccGroup, error) {

	var acc_group models.AccGroup

	var sqlStatement = `SELECT 
		id, name, descriptions
	FROM acc_group
	WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&acc_group.ID, &acc_group.Name, &acc_group.Descriptions)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return acc_group, nil
	case nil:
		return acc_group, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return acc_group, err
}

func getAllAccGroups() ([]models.AccGroup, error) {

	var results []models.AccGroup

	var sqlStatement = `SELECT 
		id, name, descriptions
	FROM acc_group
	ORDER BY id`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute account group query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccGroup

		err := rs.Scan(&p.ID, &p.Name, &p.Descriptions)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func createAccGroup(p *models.AccGroup) (int64, error) {

	sqlStatement := `INSERT INTO acc_group (id, name, descriptions) VALUES ($1, $2, $3)`

	res, err := Sql().Exec(sqlStatement, p.ID, p.Name, p.Descriptions)

	if err != nil {
		log.Printf("Unable to create account group. %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Unable to create account group. %v", err)
	}

	return rowsAffected, err
}

func updateAccGroup(id *int, p *models.AccGroup) (int64, error) {

	sqlStatement := `UPDATE acc_group SET
	name=$2, descriptions=$3
	WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		p.Name,
		p.Descriptions,
	)

	if err != nil {
		log.Printf("Unable to update account group. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating account group. %v", err)
		return 0, err
	}

	return rowsAffected, err
}

func deleteAccGroup(id *int) (int64, error) {

	sqlStatement := `DELETE FROM acc_group WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to delete account group. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected, err
}

func get_all_accounts() ([]all_accounts, error) {
	var accounts []all_accounts
	var sqlStatement = `WITH RECURSIVE rs AS (

		SELECT true as is_group, false is_type, false is_account,
		id, name, 0 as type_id, descriptions, 0 as receivable_option, false as is_active, false as is_auto_debet
		FROM acc_group

		union all

		SELECT false as is_group, true is_type, false is_account,
		id, name, group_id as type_id, descriptions, 0 as receivable_option, false is_active, false as is_auto_debet
		FROM acc_type

		union all

		SELECT false as is_group, false is_type, true is_account,
		id, name, type_id, descriptions, receivable_option, is_active, is_auto_debet
		FROM acc_code 
		ORDER BY name
	)

	select
		t.is_group, t.is_type, t.is_account,
		t.id, t.name, t.type_id, t.descriptions, t.receivable_option, t.is_active, t.is_auto_debet
	from rs t
	order by t.is_group, t.is_account, t.id;

	`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute saldo query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p all_accounts

		err := rs.Scan(
			&p.IsGroup,
			&p.IsType,
			&p.IsAccount,
			&p.ID,
			&p.Name,
			&p.TypeID,
			&p.Descriptions,
			&p.ReceivableOption,
			&p.IsActive,
			&p.IsAutoDebet,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		accounts = append(accounts, p)
	}

	return accounts, err
}
