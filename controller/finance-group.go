package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"fyc.com/sprs/models"
	"github.com/MasturFYC/fyc"
	"github.com/gin-gonic/gin"

	"net/http"

	"strconv"
)

func FinanceGroup_GetFinances(c *gin.Context) {

	fgs, err := fg_get_finances()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, &fgs)
}

func FinanceGroup_GetAll(c *gin.Context) {

	fgs, err := get_all_finance_groups()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, &fgs)
}

func FinanceGroup_GetItem(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	finances, err := get_finance_group(&id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, &finances)
}

func FinanceGroup_Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedRows, err := delete_finance_group(&id)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Finance group deleted successfully. Total rows/record affected %v", deletedRows)

	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func FinanceGroup_Create(c *gin.Context) {

	var fg models.FinanceGroup

	err := c.BindJSON(&fg)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := create_finance_group(&fg)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	fg.ID = id

	c.JSON(http.StatusOK, &fg)

}

func FinanceGroup_Update(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var fg models.FinanceGroup

	err := c.BindJSON(&fg)

	if err != nil {
		//log.Fatalf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRows, err := update_finance_group(&id, &fg)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Finance group updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
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
