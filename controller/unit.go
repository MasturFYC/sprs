package controller

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

// func GetUnits(c *gin.Context) {
//

// 	units, err := getAllUnits()

// 	if err != nil {
// 		log.Fatalf("Unable to get all units. %v", err)
// 	}

// 	c.JSON(http.StatusOK, &units)
// }

func GetUnit(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	units, err := getUnit(db, &id)

	if err != nil {
		//log.Fatalf("Unable to get unit. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &units)
}

func DeleteUnit(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedRows, err := deleteUnit(c.Keys["db"].(*sql.DB), &id)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Unit deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func CreateUnit(c *gin.Context) {

	var unit models.Unit

	err := c.BindJSON(&unit)

	//log.Printf("%v --- ", unit)

	if err != nil {
		//log.Printf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	rowAffected, err := createUnit(db, &unit)

	if err != nil {
		//log.Printf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Unit created successfully. Total rows/record affected %v", rowAffected)

	order_update_unit_token(db, &unit)
	// format the reponse message
	res := Response{
		ID:      rowAffected,
		Message: msg,
	}

	c.JSON(http.StatusOK, &res)
}

func UpdateUnit(c *gin.Context) {

	// create the postgres db connection

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var unit models.Unit

	err := c.BindJSON(&unit)

	if err != nil {
		//log.Printf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateUnit(db, &id, &unit)

	if err != nil {
		//	log.Printf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	order_update_unit_token(db, &unit)

	msg := fmt.Sprintf("Unit updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, &res)
}

func getUnit(db *sql.DB, id *int64) (models.Unit, error) {

	var unit models.Unit

	var sqlStatement = `SELECT 
		order_id, nopol, year, frame_number, machine_number,
		color, type_id, warehouse_id 
	FROM units
	WHERE order_id=$1`

	rs := db.QueryRow(sqlStatement, id)

	err := rs.Scan(&unit.OrderID,
		&unit.Nopol,
		&unit.Year,
		&unit.FrameNumber,
		&unit.MachineNumber,
		//		&unit.BpkbName,
		&unit.Color,
		//		&unit.Dealer,
		//		&unit.Surveyor,
		&unit.TypeID,
		&unit.WarehouseID,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return unit, nil
	case nil:
		t, err := getType(db, &unit.TypeID)
		if err == nil {
			unit.Type = t
		}

		w, err := getWarehouse(db, &unit.WarehouseID)

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

// 	rs, err := db.Query(sqlStatement)

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

func deleteUnit(db *sql.DB, id *int64) (int64, error) {
	// create the delete sql query
	sqlStatement := `DELETE FROM units WHERE order_id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete unit. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func create_unit_token(p *models.Unit) string {
	sb := strings.Builder{}

	sb.WriteString(" ")
	sb.WriteString(strings.Replace(p.Nopol, " ", "-", -1))
	sb.WriteString(" ")
	sb.WriteString(p.Type.Name)

	sb.WriteString(" gudang ")
	sb.WriteString(p.Warehouse.Name)

	if p.FrameNumber != "" {
		sb.WriteString(" ")
		sb.WriteString(string(p.FrameNumber))
	}

	if p.MachineNumber != "" {
		sb.WriteString(" ")
		sb.WriteString(string(p.MachineNumber))
	}
	if p.Color != "" {
		sb.WriteString(" ")
		sb.WriteString(string(p.Color))
	}

	sb.WriteString(" tahun ")
	sb.WriteString(strconv.FormatInt(p.Year, 10))

	if p.Type.MerkID > 0 {
		sb.WriteString(" ")
		sb.WriteString(p.Type.Merk.Name)
	}
	if p.Type.WheelID > 0 {
		sb.WriteString(" ")
		sb.WriteString(p.Type.Wheel.Name)
		sb.WriteString(" ")
		sb.WriteString(p.Type.Wheel.ShortName)
	}

	return sb.String()
}

func order_update_unit_token(db *sql.DB, p *models.Unit) {

	token := create_unit_token(p)

	sqlStatement := `UPDATE orders SET token=token || ' ' || to_tsvector('indonesian', $2)	WHERE id=$1`

	db.Exec(sqlStatement,
		p.OrderID,
		token,
	)

}

func createUnit(db *sql.DB, t *models.Unit) (int64, error) {

	sqlStatement := `INSERT INTO units 
	(order_id, nopol, year, frame_number, machine_number, color, type_id, warehouse_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	res, err := db.Exec(sqlStatement,
		t.OrderID,
		t.Nopol,
		t.Year,
		t.FrameNumber,
		t.MachineNumber,
		//		t.BpkbName,
		t.Color,
		// t.Dealer,
		// t.Surveyor,
		t.TypeID,
		t.WarehouseID,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func updateUnit(db *sql.DB, id *int64, t *models.Unit) (int64, error) {

	sqlStatement := `UPDATE units SET
		nopol=$2, year=$3, frame_number=$4, machine_number=$5,
		color=$6, type_id=$7, warehouse_id=$8
	WHERE order_id=$1`

	res, err := db.Exec(sqlStatement,
		id,
		t.Nopol,
		t.Year,
		t.FrameNumber,
		t.MachineNumber,
		//t.BpkbName,
		t.Color,
		//t.Dealer,
		//t.Surveyor,
		t.TypeID,
		t.WarehouseID,
	)

	if err != nil {
		//log.Printf("Unable to update unit. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}
