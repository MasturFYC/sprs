package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"fyc.com/sprs/models"
	"github.com/gin-gonic/gin"
)

type account_specific struct {
	ID           int32             `json:"id"`
	Name         string            `json:"name"`
	Descriptions models.NullString `json:"descriptions"`
}

func AccCodeGetSpec(c *gin.Context) {

	spec_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	accounts, err := get_accounts_spec(db, &spec_id)

	if err != nil || len(accounts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &accounts)
}

func AccCodeGetProps(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	acc_codes, err := getAllAccCodeProps(db)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &acc_codes)
}

func AccCodeGetAll(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	acc_codes, err := getAllAccCodes(db)

	if err != nil || len(acc_codes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &acc_codes)
}

func AccCodeSearchByName(c *gin.Context) {

	var txt = c.Param("txt")

	db := c.Keys["db"].(*sql.DB)
	acc_codes, err := searchAccCodeByName(db, &txt)

	if err != nil || len(acc_codes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &acc_codes)
}

func AccCodeGetByType(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	acc_codes, err := getAccCodeByType(db, &id)

	if err != nil || len(acc_codes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &acc_codes)
}

func AccCodeGetItem(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	acc_code, err := getAccCode(db, &id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &acc_code)
}

func AccCodeCreate(c *gin.Context) {

	var acc_code models.AccCode

	err := c.BindJSON(&acc_code)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	rowsAffected, err := createAccCode(db, &acc_code)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	res := Response{
		ID:      rowsAffected,
		Message: "Account type created successfully",
	}

	c.JSON(http.StatusOK, &res)

}

func AccCodeUpdate(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var acc_code models.AccCode

	err := c.BindJSON(&acc_code)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateAccCode(db, &id, &acc_code)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Account type updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func AccCodeDelete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	deletedRows, err := deleteAccCode(db, &id)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Account type deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      deletedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getAccCode(db *sql.DB, id *int) (models.AccInfo, error) {

	var p models.AccInfo

	b := strings.Builder{}

	b.WriteString("SELECT")
	b.WriteString(" c.type_id, c.id, c.name, c.descriptions, c.is_active, c.is_auto_debet, c.receivable_option,")
	b.WriteString(" t.name as type_name, g.name as group_name, t.descriptions as type_desc, g.descriptions as group_desc")
	b.WriteString(" FROM acc_code c")
	b.WriteString(" INNER JOIN acc_type t ON t.id = c.type_id")
	b.WriteString(" INNER JOIN acc_group g ON g.id = t.group_id")
	b.WriteString(" WHERE c.id=$1")

	rs := db.QueryRow(b.String(), id)

	err := rs.Scan(
		&p.TypeID,
		&p.ID,
		&p.Name,
		&p.Descriptions,
		&p.IsActive,
		&p.IsAutoDebet,
		&p.ReceivableOption,
		&p.TypeName,
		&p.GroupName,
		&p.TypeDesc,
		&p.GroupDesc,
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

func getAccCodeByType(db *sql.DB, id *int) ([]models.AccCode, error) {

	var results []models.AccCode

	var sqlStatement = `SELECT 
		type_id, id, name, descriptions, is_active, is_auto_debet, receivable_option
	FROM acc_code
	WHERE type_id=$1 OR 0=$1
	ORDER BY id`

	rs, err := db.Query(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to execute account code query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCode

		err := rs.Scan(
			&p.TypeID,
			&p.ID,
			&p.Name,
			&p.Descriptions,
			&p.IsActive,
			&p.IsAutoDebet,
			&p.ReceivableOption,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func searchAccCodeByName(db *sql.DB, txt *string) ([]models.AccCode, error) {

	var results []models.AccCode

	var sqlStatement = `SELECT 
		type_id, id, name, descriptions, is_active, is_auto_debet, receivable_option
	FROM acc_code
	WHERE token_name @@ to_tsquery('indonesian', $1)
	ORDER BY id`

	rs, err := db.Query(sqlStatement, txt)

	if err != nil {
		log.Printf("Unable to execute account code query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCode

		err := rs.Scan(
			&p.TypeID,
			&p.ID,
			&p.Name,
			&p.Descriptions,
			&p.IsActive,
			&p.IsAutoDebet,
			&p.ReceivableOption,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func getAllAccCodes(db *sql.DB) ([]models.AccCode, error) {

	var results []models.AccCode

	var sqlStatement = `SELECT 
		type_id, id, name, descriptions, is_active, is_auto_debet, receivable_option
	FROM acc_code
	ORDER BY id`

	rs, err := db.Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute account code query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCode

		err := rs.Scan(
			&p.TypeID,
			&p.ID,
			&p.Name,
			&p.Descriptions,
			&p.IsActive,
			&p.IsAutoDebet,
			&p.ReceivableOption,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func createAccCode(db *sql.DB, p *models.AccCode) (int64, error) {

	sqlStatement := `INSERT INTO 
	acc_code (type_id, id, name, descriptions, is_active, is_auto_debet, receivable_option, token_name)
	VALUES ($1, $2, $3, $4, $5, $6, $7, to_tsvector('indonesian', $8))`

	token := fmt.Sprintf("%s %s", p.Name, p.Descriptions)

	res, err := db.Exec(sqlStatement,
		p.TypeID,
		p.ID,
		p.Name,
		p.Descriptions,
		p.IsActive,
		p.IsAutoDebet,
		p.ReceivableOption,
		token,
	)

	if err != nil {
		log.Printf("Unable to create account code. %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Unable to create account code. %v", err)
	}

	return rowsAffected, err
}

func updateAccCode(db *sql.DB, id *int, p *models.AccCode) (int64, error) {

	sqlStatement := `UPDATE acc_code SET 
	type_id=$2, name=$3, descriptions=$4, is_active=$5, is_auto_debet=$6,
	receivable_option=$7,
	token_name=to_tsvector('indonesian', $8)	
	WHERE id=$1`

	token := fmt.Sprintf("%s %s", p.Name, p.Descriptions)

	res, err := db.Exec(sqlStatement,
		id,
		p.TypeID,
		p.Name,
		p.Descriptions,
		p.IsActive,
		p.IsAutoDebet,
		p.ReceivableOption,
		token,
	)

	if err != nil {
		log.Printf("Unable to update account code. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error while updating account code. %v", err)
		return 0, err
	}

	return rowsAffected, err
}

func deleteAccCode(db *sql.DB, id *int) (int64, error) {

	sqlStatement := `DELETE FROM acc_code WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to delete account code. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected, err
}

func getAllAccCodeProps(db *sql.DB) ([]models.AccCodeType, error) {

	var results []models.AccCodeType

	var sqlStatement = `SELECT 
		c.type_id, c.id, c.name, t.name AS type_name, c.descriptions, c.is_active
	FROM acc_code c
	INNER JOIN acc_type t ON t.id = c.type_id
	ORDER BY c.id`

	rs, err := db.Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute account code property query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p models.AccCodeType

		err := rs.Scan(
			&p.TypeID,
			&p.ID,
			&p.Name,
			&p.TypeName,
			&p.Descriptions,
			&p.IsActive,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}

func get_accounts_spec(db *sql.DB, specId *int) ([]account_specific, error) {

	var results []account_specific

	var sqlStatement = `SELECT 
		id, name, descriptions
	FROM acc_code
	WHERE receivable_option = $1
	ORDER BY id`

	rs, err := db.Query(sqlStatement, specId)

	if err != nil {
		log.Printf("Unable to execute account code property query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p account_specific

		err := rs.Scan(
			&p.ID,
			&p.Name,
			&p.Descriptions,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		results = append(results, p)
	}

	return results, err
}
