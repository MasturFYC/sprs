package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/MasturFYC/fyc"
	"github.com/gin-gonic/gin"
)

// type local_detail struct {
// 	ID        int64   `json:"id"`
// 	Name      string  `json:"name"`
// 	AccCodeID int32   `json:"accCodeId"`
// 	Debt      float64 `json:"debt"`
// 	Cred      float64 `json:"cred"`
// }

type local_trx struct {
	models.Trx
	Details json.RawMessage `json:"details,omitempty"`
}

func TransactionGetByMonth(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	acc_codes, err := get_trx_by_month(db, &id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &acc_codes)
}

func TransactionGetByGroup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	acc_codes, err := getTransactionsByGroup(db, &id)

	if err != nil {
		//log.Printf("Unable to get all transactions. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &acc_codes)
}

func TransactionSearch(c *gin.Context) {
	var t models.SearchGroup

	err := c.BindJSON(&t)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	trxs, err := searchTransactions(db, &t.Txt)

	if err != nil {
		//log.Printf("Unable to get all account codes. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &trxs)
}

func TransactionGetAll(c *gin.Context) {

	db := c.Keys["db"].(*sql.DB)
	trxs, err := get_all_transactions(db)

	if err != nil {
		//log.Printf("Unable to get all transaction. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &trxs)
}

func GetTransaction(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	trx, err := getTransaction(db, &id)

	if err != nil {
		//log.Printf("Unable to get transaction. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &trx)
}

func TransactionCreate(c *gin.Context) {

	var param models.TrxDetailsToken

	err := c.BindJSON(&param)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	id, err := createTransaction(db, &param.Trx, param.Token)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	if len(param.Details) > 0 {

		err = bulkInsertDetails(db, param.Details, &id)

		if err != nil {
			//log.Printf("Unable to insert transaction details.  %v", err)
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
	}

	res := Response{
		ID:      id,
		Message: "Transaction was succesfully inserted.",
	}

	c.JSON(http.StatusOK, &res)

}

func TransactionUpdate(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Printf("Unable to decode the request body to transaction.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var trx models.TrxDetailsToken

	err = c.BindJSON(&trx)

	if err != nil {
		//	log.Printf("Unable to decode the request body to transaction.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateTransaction(db, &id, &trx.Trx, trx.Token)

	if err != nil {
		//log.Printf("Unable to update transaction.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	if len(trx.Details) > 0 {

		_, err = deleteDetailsByOrder(db, &id)
		if err != nil {
			//log.Printf("Unable to delete all details by transaction.  %v", err)
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
		//}

		// 	var newId int64 = 0

		err = bulkInsertDetails(db, trx.Details, &id)

		if err != nil {
			//log.Printf("Unable to insert transaction details (message from command).  %v", err)
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
	}

	msg := fmt.Sprintf("Transaction updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func deleteDetailsByOrder(db *sql.DB, id *int64) (int64, error) {
	sqlStatement := `DELETE FROM trx_detail WHERE trx_id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		//log.Printf("Unable to delete transaction. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err

}

func bulkInsertDetails(db *sql.DB, rows []models.TrxDetail, id *int64) error {
	valueStrings := make([]string, 0, len(rows))
	valueArgs := make([]interface{}, 0, len(rows)*5)
	i := 0
	for _, post := range rows {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		valueArgs = append(valueArgs, post.ID)
		valueArgs = append(valueArgs, post.CodeID)
		valueArgs = append(valueArgs, setTrxID(id, &post.TrxID))
		valueArgs = append(valueArgs, post.Debt)
		valueArgs = append(valueArgs, post.Cred)
		i++
	}
	stmt := fmt.Sprintf("INSERT INTO trx_detail (id, code_id, trx_id, debt, cred) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt+" ON CONFLICT (trx_id, id) DO UPDATE SET code_id=EXCLUDED.code_id, trx_id=EXCLUDED.trx_id, debt=EXCLUDED.debt, cred=EXCLUDED.cred", valueArgs...)

	//log.Printf("%s %s", strings.Join(valueStrings, ","), valueArgs)
	return err
}

func setTrxID(id *int64, id2 *int64) int64 {
	if (*id) == 0 {
		return (*id2)
	}
	return (*id)
}

func TransactionDelete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)

	deletedRows, err := deleteTransaction(db, &id)

	if err != nil {
		log.Printf("Unable to delete transaction.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Transaction deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      deletedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getTransaction(db *sql.DB, id *int64) (local_trx, error) {

	var p local_trx
	builder := strings.Builder{}

	builder.WriteString("SELECT t.id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,")
	//builder.WriteString(" (SELECT COALESCE(sum(s.debt),0) AS debt FROM trx_detail s WHERE s.trx_id = t.id)")
	//builder.WriteString(" AS saldo, ")
	builder.WriteString(fyc.NestQuery(get_query_details))
	builder.WriteString(" AS details ")
	builder.WriteString(" FROM trx t")
	builder.WriteString(" WHERE t.id=$1")

	rs := db.QueryRow(builder.String(), id)

	err := rs.Scan(
		&p.ID,
		&p.RefID,
		&p.Division,
		&p.TrxDate,
		&p.Descriptions,
		&p.Memo,
		//&p.Saldo,
		&p.Details,
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

	// d, _ := get_details(&p.ID)
	// p.Details = d

	// return empty user on error
	return p, err
}

const (
	get_query_details = `SELECT c.id, c.name, d.code_id AS "codeId", d.debt, d.cred
	FROM trx_detail d
	INNER JOIN acc_code c ON c.id = d.code_id
	WHERE d.trx_id=t.id 
	ORDER BY d.id`
)

func get_all_transactions(db *sql.DB) ([]local_trx, error) {

	var results = make([]local_trx, 0)

	builder := strings.Builder{}

	builder.WriteString("SELECT t.id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,")
	//builder.WriteString(" (SELECT COALESCE(sum(s.debt),0) AS debt FROM trx_detail s WHERE s.trx_id = t.id)")
	//builder.WriteString(" AS saldo, ")
	builder.WriteString(fyc.NestQuery(get_query_details))
	builder.WriteString(" AS details ")
	builder.WriteString(" FROM trx t")
	builder.WriteString(" ORDER BY t.id DESC")

	// 	var sqlStatement = `SELECT t.id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,
	// (SELECT COALESCE(sum(s.debt),0) AS debt FROM trx_detail s WHERE s.trx_id = t.id) saldo,
	// COALESCE
	// (
	// (
	// SELECT array_to_json(array_agg(row_to_json(x))) FROM
	// (
	// SELECT c.id, c.name, d.code_id, d.debt, d.cred
	// FROM trx_detail d
	// INNER JOIN acc_code c ON c.id = d.code_id
	// WHERE d.trx_id=t.id
	// ORDER BY d.id
	// ) x
	// ),
	// '[]'
	// ) AS details

	// FROM trx t
	// ORDER BY t.id DESC`

	rs, err := db.Query(builder.String())

	if err != nil {
		log.Printf("Unable to execute transaction query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p local_trx

		err := rs.Scan(
			&p.ID,
			&p.RefID,
			&p.Division,
			&p.TrxDate,
			&p.Descriptions,
			&p.Memo,
			//&p.Saldo,
			&p.Details,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// d, _ := get_details(&p.ID)
		// p.Details = d

		results = append(results, p)
	}

	return results, err
}

func createTransaction(db *sql.DB, p *models.Trx, token string) (int64, error) {

	sqlStatement := `INSERT INTO trx 
	(ref_id, division, trx_date, descriptions, memo, trx_token)
	VALUES ($1, $2, $3, $4, $5, to_tsvector('indonesian', $6))
	RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement,
		p.RefID,
		p.Division,
		p.TrxDate,
		p.Descriptions,
		p.Memo,
		token,
	).Scan(&id)

	if err != nil {
		log.Printf("Unable to create transaction. %v", err)
		return 0, err
	}

	return id, err
}

func trxGetOrder(db *sql.DB, orderId *int64) (int64, error) {

	var id int64

	rs := db.QueryRow("SELECT t.id FROM trx t WHERE t.ref_id=$1 AND division='trx-order'", orderId)

	err := rs.Scan(&id)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return id, err
	case nil:
		return id, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return id, err
}

func trxMoveToLent(db *sql.DB, id *int64, p *models.Trx, token string) (int64, error) {

	sb := strings.Builder{}
	sb.WriteString("UPDATE trx SET")
	sb.WriteString(" division='trx-lent'")
	sb.WriteString(", descriptions=$2, memo=$3")
	sb.WriteString(", trx_token=to_tsvector('indonesian', $4)")
	sb.WriteString(" WHERE id=$1")

	//var trxid int64

	_, err := db.Exec(sb.String(),
		id, p.Descriptions, p.Memo, token,
	)

	if err != nil {
		log.Printf("Unable move trx to lent. %v", err)
		return 0, err
	}

	res, err := db.Exec("UPDATE trx_detail SET code_id = 5513 WHERE trx_id = $1 AND code_id = 5511", id)

	if err != nil {
		log.Printf("Unable move trx to lent. %v", err)
		return 0, err
	}
	// // check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func updateTransaction(db *sql.DB, id *int64, p *models.Trx, token string) (int64, error) {

	sqlStatement := `UPDATE trx SET 
		ref_id=$2,
		division=$3,
		trx_date=$4,
		descriptions=$5,
		memo=$6,
		trx_token=to_tsvector('indonesian', $7)
	WHERE id=$1`

	res, err := db.Exec(sqlStatement,
		id,
		p.RefID,
		p.Division,
		p.TrxDate,
		p.Descriptions,
		p.Memo,
		token,
	)

	if err != nil {
		log.Printf("Unable to update transaction. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func deleteTransaction(db *sql.DB, id *int) (int64, error) {

	sqlStatement := `DELETE FROM trx WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Printf("Unable to delete transaction. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func searchTransactions(db *sql.DB, txt *string) ([]local_trx, error) {

	var results = make([]local_trx, 0)

	builder := strings.Builder{}

	builder.WriteString("SELECT t.id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,")
	//builder.WriteString(" (SELECT COALESCE(sum(s.debt),0) AS debt FROM trx_detail s WHERE s.trx_id = t.id)")
	//builder.WriteString(" AS saldo, ")
	builder.WriteString(fyc.NestQuery(get_query_details))
	builder.WriteString(" AS details ")
	builder.WriteString(" FROM trx t")
	builder.WriteString(" WHERE t.trx_token @@ to_tsquery('indonesian', $1)")
	builder.WriteString(" ORDER BY t.id DESC")

	rs, err := db.Query(builder.String(), txt)

	if err != nil {
		log.Printf("Unable to execute transactions code query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p local_trx

		err := rs.Scan(
			&p.ID,
			&p.RefID,
			&p.Division,
			&p.TrxDate,
			&p.Descriptions,
			&p.Memo,
			//&p.Saldo,
			&p.Details,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// d, _ := get_details(&p.ID)
		// p.Details = d

		results = append(results, p)
	}

	return results, err
}

func getTransactionsByGroup(db *sql.DB, id *int64) ([]local_trx, error) {

	var results = make([]local_trx, 0)

	builder := strings.Builder{}

	builder.WriteString("SELECT t.id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,")
	//builder.WriteString(" (SELECT COALESCE(sum(s.debt),0) AS debt FROM trx_detail s WHERE s.trx_id = t.id)")
	//builder.WriteString(" AS saldo, ")
	builder.WriteString(fyc.NestQuery(get_query_details))
	builder.WriteString(" AS details ")
	builder.WriteString(" FROM trx t")
	builder.WriteString(" INNER JOIN trx_detail d ON d.trx_id = d.id")
	builder.WriteString(" INNER JOIN acc_type e ON e.id = d.code_id")
	builder.WriteString(" INNER JOIN acc_code c ON c.id = e.group_id")
	builder.WriteString(" WHERE c.group_id=$1")
	builder.WriteString(" ORDER BY t.id DESC")

	rs, err := db.Query(builder.String(), id)

	if err != nil {
		log.Printf("Unable to execute transactions query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p local_trx

		err := rs.Scan(
			&p.ID,
			&p.RefID,
			&p.Division,
			&p.TrxDate,
			&p.Descriptions,
			&p.Memo,
			//&p.Saldo,
			&p.Details,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		// d, _ := get_details(&p.ID)
		// p.Details = d

		results = append(results, p)
	}

	return results, err
}

/*
func get_details(trxID *int64) ([]local_detail, error) {

	var details []local_detail

	var sqlStatement = `SELECT
	c.id, c.name, d.code_id, d.debt, d.cred
	FROM trx_detail d
	INNER JOIN acc_code c ON c.id = d.code_id
	WHERE d.trx_id=$1
	-- AND c.receivable_option != 1
	ORDER BY d.id`

	rs, err := db.Query(sqlStatement, trxID)

	if err != nil {
		log.Printf("Unable to execute transaction details query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p local_detail

		err := rs.Scan(
			&p.ID,
			&p.Name,
			&p.AccCodeID,
			&p.Debt,
			&p.Cred,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		details = append(details, p)
	}

	return details, err
}
*/
func get_trx_by_month(db *sql.DB, id *int) ([]local_trx, error) {

	var results = make([]local_trx, 0)

	builder := strings.Builder{}

	builder.WriteString("SELECT t.id, t.ref_id, t.division, t.trx_date, t.descriptions, t.memo,")
	//builder.WriteString(" (SELECT COALESCE(sum(s.debt),0) AS debt FROM trx_detail s WHERE s.trx_id = t.id)")
	//builder.WriteString(" AS saldo, ")
	builder.WriteString(fyc.NestQuery(get_query_details))
	builder.WriteString(" AS details ")
	builder.WriteString(" FROM trx t")
	builder.WriteString(" WHERE EXTRACT(MONTH from t.trx_date)=$1")
	builder.WriteString(" ORDER BY t.id DESC")

	rs, err := db.Query(builder.String(), id)

	if err != nil {
		log.Printf("Unable to execute transactions query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p local_trx

		err := rs.Scan(
			&p.ID,
			&p.RefID,
			&p.Division,
			&p.TrxDate,
			&p.Descriptions,
			&p.Memo,
			//&p.Saldo,
			&p.Details,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		// d, _ := get_details(&p.ID)
		// p.Details = d

		results = append(results, p)
	}

	return results, err
}
