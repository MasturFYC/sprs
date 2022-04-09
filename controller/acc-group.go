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

type all_accounts struct {
	models.AccCode
	IsGroup   bool `json:"isGroup"`
	IsType    bool `json:"isType"`
	IsAccount bool `json:"isAccount"`
}

func AccGroupGetAllAccount(c *gin.Context) {
	db := c.Keys["db"].(*sql.DB)
	accounts, err := get_all_accounts(db)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &accounts)
}

func AccGroupGetAll(c *gin.Context) {
	db := c.Keys["db"].(*sql.DB)
	groups, err := getAllAccGroups(db)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &groups)
}

func AccGroupGetTypes(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	acc_group, err := group_get_types(db, &id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &acc_group)
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

func GetAccGroup(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	acc_group, err := getAccGroup(db, &id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &acc_group)
}

func AccGroupCreate(c *gin.Context) {

	var acc_group models.AccGroup

	if err := c.BindJSON(&acc_group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	rowsAffected, err := createAccGroup(db, &acc_group)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"rowsAffected": rowsAffected, "message": "Group was created"})

}

func AccGroupUpdate(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var acc_group models.AccGroup

	if err := c.BindJSON(&acc_group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateAccGroup(db, &id, &acc_group)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rowsAffected": updatedRows, "message": "Group was updated"})
}

func AccGroupDelete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	deletedRows, err := deleteAccGroup(db, &id)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rowsDeleted": deletedRows, "message": "Account group deleted successfully."})

}

func group_get_types(db *sql.DB, id *int) ([]models.AccType, error) {
	var results = make([]models.AccType, 0)

	var sqlStatement = `SELECT group_id, id, name, descriptions FROM acc_type WHERE group_id=$1 OR 0 = $1 ORDER BY id`

	rs, err := db.Query(sqlStatement, id)

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

func getAccGroup(db *sql.DB, id *int) (models.AccGroup, error) {

	var acc_group models.AccGroup

	var sqlStatement = "SELECT id, name, descriptions FROM acc_group WHERE id=$1"

	rs := db.QueryRow(sqlStatement, id)

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

func getAllAccGroups(db *sql.DB) ([]models.AccGroup, error) {

	var results = make([]models.AccGroup, 0)

	var sqlStatement = `SELECT id, name, descriptions FROM acc_group ORDER BY id`

	rs, err := db.Query(sqlStatement)

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

func createAccGroup(db *sql.DB, p *models.AccGroup) (int64, error) {

	sqlStatement := `INSERT INTO acc_group (id, name, descriptions) VALUES ($1, $2, $3)`

	res, err := db.Exec(sqlStatement, p.ID, p.Name, p.Descriptions)

	if err != nil {
		log.Printf("Unable to create account group. %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func updateAccGroup(db *sql.DB, id *int, p *models.AccGroup) (int64, error) {

	sqlStatement := `UPDATE acc_group SET
	name=$2, descriptions=$3
	WHERE id=$1`

	res, err := db.Exec(sqlStatement,
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

	return rowsAffected, err
}

func deleteAccGroup(db *sql.DB, id *int) (int64, error) {

	sqlStatement := `DELETE FROM acc_group WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to delete account group. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func get_all_accounts(db *sql.DB) ([]all_accounts, error) {
	var accounts = make([]all_accounts, 0)

	sb := strings.Builder{}
	sb.WriteString("WITH RECURSIVE rs AS (")
	sb.WriteString("SELECT true as is_group, false is_type, false is_account,")
	sb.WriteString(" id, name, 0 as type_id, descriptions, 0 as receivable_option, false as is_active, false as is_auto_debet")
	sb.WriteString(" FROM acc_group")
	sb.WriteString("\n\nUNION ALL\n\n")
	sb.WriteString("SELECT false as is_group, true is_type, false is_account,")
	sb.WriteString(" id, name, group_id as type_id, descriptions, 0 as receivable_option, false is_active, false as is_auto_debet")
	sb.WriteString(" FROM acc_type")
	sb.WriteString("\n\nUNION ALL\n\n")
	sb.WriteString("SELECT false as is_group, false is_type, true is_account,")
	sb.WriteString(" id, name, type_id, descriptions, receivable_option, is_active, is_auto_debet")
	sb.WriteString(" FROM acc_code")
	sb.WriteString(" ORDER BY name")
	sb.WriteString(")\n")
	sb.WriteString("SELECT")
	sb.WriteString(" t.is_group, t.is_type, t.is_account,")
	sb.WriteString(" t.id, t.name, t.type_id, t.descriptions, t.receivable_option, t.is_active, t.is_auto_debet")
	sb.WriteString(" FROM rs t")
	sb.WriteString(" ORDER BY t.is_group, t.is_account, t.id;")

	rs, err := db.Query(sb.String())

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
