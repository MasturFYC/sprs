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

// func GetUnits(w http.ResponseWriter, r *http.Request) {
// 	EnableCors(&w)

// 	units, err := getAllUnits()

// 	if err != nil {
// 		log.Fatalf("Unable to get all units. %v", err)
// 	}

// 	json.NewEncoder(w).Encode(&units)
// }

func GetUnit(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	units, err := getUnit(&id)

	if err != nil {
		log.Fatalf("Unable to get unit. %v", err)
	}

	json.NewEncoder(w).Encode(&units)
}

func DeleteUnit(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteUnit(&id)

	msg := fmt.Sprintf("Unit deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreateUnit(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var unit models.Unit

	err := json.NewDecoder(r.Body).Decode(&unit)

	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusRequestedRangeNotSatisfiable), http.StatusRequestedRangeNotSatisfiable)
		return
	}

	rowAffected, err := createUnit(&unit)

	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Unit created successfully. Total rows/record affected %v", rowAffected)

	// format the reponse message
	res := Response{
		ID:      rowAffected,
		Message: msg,
	}

	json.NewEncoder(w).Encode(&res)
}

func UpdateUnit(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var unit models.Unit

	err := json.NewDecoder(r.Body).Decode(&unit)

	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusRequestedRangeNotSatisfiable), http.StatusRequestedRangeNotSatisfiable)
		return
	}

	updatedRows, err := updateUnit(&id, &unit)

	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	msg := fmt.Sprintf("Unit updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(&res)
}

func getUnit(id *int64) (models.Unit, error) {

	var unit models.Unit

	var sqlStatement = `SELECT 
		order_id, nopol, year, frame_number, machine_number, bpkb_name,
		color, dealer, surveyor, type_id, warehouse_id 
	FROM units
	WHERE order_id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(&unit.OrderID,
		&unit.Nopol,
		&unit.Year,
		&unit.FrameNumber,
		&unit.MachineNumber,
		&unit.BpkbName,
		&unit.Color,
		&unit.Dealer,
		&unit.Surveyor,
		&unit.TypeID,
		&unit.WarehouseID,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return unit, nil
	case nil:
		t, err := getType(&unit.TypeID)
		if err == nil {
			unit.Type = t
		}

		w, err := getWarehouse(&unit.WarehouseID)

		if err == nil {
			unit.Warehouse = w
		}

		return unit, nil
	default:
		log.Printf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return unit, err
}

// func getAllUnits() ([]models.Unit, error) {

// 	var units []models.Unit

// 	var sqlStatement = `SELECT
// 		order_id, nopol, year, frame_number, machine_number, bpkb_name,
// 		color, dealer, surveyor, type_id, warehouse_id
// 	FROM units`

// 	rs, err := Sql().Query(sqlStatement)

// 	if err != nil {
// 		log.Fatalf("Unable to execute units query %v", err)
// 	}

// 	defer rs.Close()

// 	for rs.Next() {
// 		var unit models.Unit

// 		err := rs.Scan(&unit.OrderID,
// 			&unit.Nopol,
// 			&unit.Year,
// 			&unit.FrameNumber,
// 			&unit.MachineNumber,
// 			&unit.BpkbName,
// 			&unit.Color,
// 			&unit.Dealer,
// 			&unit.Surveyor,
// 			&unit.TypeID,
// 			&unit.WarehouseID,
// 		)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		units = append(units, unit)
// 	}

// 	return units, err
// }

func deleteUnit(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM units WHERE id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete unit. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createUnit(t *models.Unit) (int64, error) {

	sqlStatement := `INSERT INTO units 
	(order_id, nopol, year, frame_number, machine_number, bpkb_name,
		color, dealer, surveyor, type_id, warehouse_id)
	VALUES ($1, $2, $3, $4, $5, $6,
		$7, $8, $9, $10, $11)`

	res, err := Sql().Exec(sqlStatement,
		t.OrderID,
		t.Nopol,
		t.Year,
		t.FrameNumber,
		t.MachineNumber,
		t.BpkbName,
		t.Color,
		t.Dealer,
		t.Surveyor,
		t.TypeID,
		t.WarehouseID,
	)

	if err != nil {
		log.Fatalf("Unable to create unit. %v", err)
	}

	if err != nil {
		log.Fatalf("Unable to create customer. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Unable to create customer. %v", err)
	}

	return rowsAffected, err
}

func updateUnit(id *int64, t *models.Unit) (int64, error) {

	sqlStatement := `UPDATE units SET
		nopol=$2, year=$3, frame_number=$4, machine_number=$5, bpkb_name=$6,
		color=$7, dealer=$8, surveyor=$9, type_id=$10, warehouse_id=$11
	WHERE order_id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		t.Nopol,
		t.Year,
		t.FrameNumber,
		t.MachineNumber,
		t.BpkbName,
		t.Color,
		t.Dealer,
		t.Surveyor,
		t.TypeID,
		t.WarehouseID,
	)

	if err != nil {
		log.Printf("Unable to update unit. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating unit. %v", err)
	}

	return rowsAffected, err
}
