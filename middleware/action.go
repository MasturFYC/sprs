package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"fyc.com/sprs/models"
	"github.com/gorilla/mux"
)

// get all action by order
func GetActions(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	actions, err := getAllActions(&id)

	if err != nil {
		log.Fatalf("Unable to get all actions. %v", err)
	}

	json.NewEncoder(w).Encode(&actions)
}

func GetAction(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	act, err := getAction(&id)

	if err != nil {
		log.Fatalf("Unable to get category. %v", err)
	}

	json.NewEncoder(w).Encode(&act)
}

func DeleteAction(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteAction(&id)

	msg := fmt.Sprintf("Action deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

func CreateAction(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var act models.Action

	err := json.NewDecoder(r.Body).Decode(&act)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	id, err := createAction(&act)

	if err == nil {
		act.ID = id
	}

	json.NewEncoder(w).Encode(&act)
}

func UpdateAction(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var act models.Action

	err := json.NewDecoder(r.Body).Decode(&act)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateCategory(&id, &act)

	msg := fmt.Sprintf("Action updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

func getAllActions(OrderID *int64) ([]models.Action, error) {
	// defer Sql().Close()
	var actions []models.Action

	sqlStatement := `SELECT
		id, action_at, code, pic, descriptions, order_id
	FROM actions
	WHERE order_id=$1`

	rows, err := Sql().Query(sqlStatement, OrderID)

	if err != nil {
		log.Fatalf("Unable to execute actions query %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var act models.Action
		err := rows.Scan(&act.ID, &act.ActionAt, &act.Pic, &act.Descriptions, &act.OrderId)

		if err != nil {
			log.Fatalf("ACTION unable to scan the row. %v", err)
		}

		actions = append(actions, act)
	}

	return actions, err
}

func getAction(id *int64) (models.Action, error) {
	var act models.Action

	sqlStatement := `SELECT 
		id, action_at, code, pic, descriptions, order_id
	FROM actions 
	WHERE id=$1`
	//stmt, _ := Sql().Prepare(sqlStatement)

	//defer stmt.Close()
	row := Sql().QueryRow(sqlStatement, id)

	err := row.Scan(&act.ID, &act.ActionAt, &act.Pic, &act.Descriptions, &act.OrderId)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return act, nil
	case nil:
		return act, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return act, err
}

// func getProductsByCategory(id *int) ([]models.Product, error) {
// 	// defer Sql().Close()

// 	var products []models.Product

// 	sqlStatement := `SELECT
// 		p.id, p.name, p.spec, p.base_unit,
// 		p.base_weight, p.base_price, p.first_stock,
// 		p.stock, p.is_active, p.is_sale, p.category_id
// 	FROM products AS p
// 	WHERE p.category_id=$1
// 	ORDER BY p.name`

// 	rows, err := Sql().Query(sqlStatement, id)

// 	if err != nil {
// 		log.Fatalf("Unable to execute product query %v", err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {

// 		var product models.Product

// 		err := rows.Scan(
// 			&product.ID,
// 			&product.Name,
// 			&product.Spec,
// 			&product.BaseUnit,
// 			&product.BaseWeight,
// 			&product.BasePrice,
// 			&product.FirstStock,
// 			&product.Stock,
// 			&product.IsActive,
// 			&product.IsSale,
// 			&product.CategoryID,
// 		)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		products = append(products, product)
// 	}

// 	return products, err
// }

func deleteAction(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM actions WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete action. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createAction(act *models.Action) (int64, error) {

	sqlStatement := `INSERT INTO actions
		(action_at, code, pic, descriptions, order_id)
	VALUES
		($1, $2, $3, $4, $5)
	RETURNING id`

	var id int64

	err := Sql().QueryRow(sqlStatement,
		act.ActionAt,
		act.Code,
		act.Pic,
		act.Descriptions,
		act.OrderId).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to create action. %v", err)
	}

	return id, err
}

func updateCategory(id *int64, act *models.Action) int64 {

	sqlStatement := `UPDATE actions SET 
		action_at=$2, code=$3, pic=$4, descriptions=$5, order_id=$6
	WHERE id=$1`

	res, err := Sql().Exec(sqlStatement, id, act.ActionAt,
		act.Code, act.Pic, act.Descriptions, act.OrderId)

	if err != nil {
		log.Fatalf("Unable to update action. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating action. %v", err)
	}

	return rowsAffected
}
