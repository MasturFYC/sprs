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

func GetWarehouses(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	warehouses, err := getAllWarehouses()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
		//log.Fatalf("Unable to get all warehouses. %v", err)
	}

	json.NewEncoder(w).Encode(&warehouses)
}

func GetWarehouse(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//		log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	warehouses, err := getWarehouse(&id)

	if err != nil {
		//log.Fatalf("Unable to get warehouse. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&warehouses)
}

func DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	deletedRows, err := deleteWarehouse(&id)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Warehouse deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var warehouse models.Warehouse

	err := json.NewDecoder(r.Body).Decode(&warehouse)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := createWarehouse(&warehouse)

	if err != nil {
		//log.Fatalf("Nama warehouse tidak boleh sama.  %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	warehouse.ID = id

	json.NewEncoder(w).Encode(&warehouse)

}

func UpdateWarehouse(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var warehouse models.Warehouse

	err := json.NewDecoder(r.Body).Decode(&warehouse)

	if err != nil {
		//		log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}

	updatedRows, err := updateWarehouse(&id, &warehouse)

	if err != nil {
		//		log.Fatalf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}

	msg := fmt.Sprintf("Warehouse updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getWarehouse(id *int) (models.Warehouse, error) {

	var warehouse models.Warehouse

	var sqlStatement = `SELECT id, name, descriptions FROM warehouses WHERE id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Descriptions)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return warehouse, nil
	case nil:
		return warehouse, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return warehouse, err
}

func getAllWarehouses() ([]models.Warehouse, error) {

	var warehouses []models.Warehouse

	var sqlStatement = `SELECT id, name, descriptions FROM warehouses`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute warehouses query %v", err)
	}

	defer rs.Close()

	for rs.Next() {
		var warehouse models.Warehouse

		err := rs.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Descriptions)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		warehouses = append(warehouses, warehouse)
	}

	return warehouses, err
}

func deleteWarehouse(id *int) (int64, error) {
	// create the delete sql query
	sqlStatement := `DELETE FROM warehouses WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		//log.Fatalf("Unable to delete warehouse. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Fatalf("Error while checking the affected rows. %v", err)
	// }

	return rowsAffected, err
}

func createWarehouse(warehouse *models.Warehouse) (int, error) {

	sqlStatement := `INSERT INTO warehouses (name, descriptions) VALUES ($1, $2) RETURNING id`

	var id int

	err := Sql().QueryRow(sqlStatement,
		warehouse.Name,
		warehouse.Descriptions,
	).Scan(&id)

	if err != nil {
		//log.Fatalf("Unable to create warehouse. %v", err)
		return 0, err
	}

	warehouse.ID = id

	return id, nil
}

func updateWarehouse(id *int, warehouse *models.Warehouse) (int64, error) {

	sqlStatement := `UPDATE warehouses SET name=$2, descriptions=$3 WHERE id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		warehouse.Name,
		warehouse.Descriptions,
	)

	if err != nil {
		//log.Fatalf("Unable to update warehouse. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	// if err != nil {
	// 	log.Fatalf("Error while updating warehouse. %v", err)
	// }

	return rowsAffected, err
}
