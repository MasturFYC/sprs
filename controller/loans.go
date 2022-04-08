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

func LoanGetAll(c *gin.Context) {
	db := c.Keys["db"].(*sql.DB)

	loans, err := loan_get_all(db)

	if err != nil {
		//		log.Fatalf("Unable to get all finances. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &loans)
}

func LoanGetItem(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)

	loan, err := loan_get_item(db, &id)

	if err != nil {
		//log.Fatalf("Unable to get finance. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &loan)
}

func LoanDelete(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.Keys["db"].(*sql.DB)

	deletedRows, err := loan_delete(db, &id)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("loan deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      id,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

type ts_loan_create struct {
	Loan  models.Loan `json:"loan"`
	Trx   models.Trx  `json:"trx"`
	Token string      `json:"token"`
}

func LoanCreate(c *gin.Context) {

	var loan ts_loan_create

	err := c.BindJSON(&loan)

	if err != nil {
		log.Fatalf("Fatal %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	newLoanId, err := loan_create(db, &loan.Loan)
	if err != nil {
		log.Fatalf("Fatal %v", err)
		//log.Fatalf("Nama finance tidak boleh sama.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	loan.Loan.ID = newLoanId
	loan.Trx.RefID = newLoanId
	trxid, err := createTransaction(db, &loan.Trx, loan.Token)

	if err != nil {
		log.Fatalf("Fatal %v", err)
		//log.Printf("(API) Unable to create transaction.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	if len(loan.Trx.Details) > 0 {

		err = bulkInsertDetails(db, loan.Trx.Details, &trxid)

		if err != nil {
			log.Fatalf("Fatal %v", err)
			//log.Printf("Unable to insert transaction details.  %v", err)
			//c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
	}

	loan.Trx.ID = trxid
	c.JSON(http.StatusOK, &loan)

}

type ts_loan_payment struct {
	Trx   models.Trx `json:"trx"`
	Token string     `json:"token"`
}

func LoanPayment(c *gin.Context) {

	trxid, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var data ts_loan_payment

	err := c.BindJSON(&data)

	if err != nil {
		log.Fatalf("Unable to decode trx from body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedRows int64 = 0
	db := c.Keys["db"].(*sql.DB)

	if trxid == 0 {
		trxid, err = createTransaction(db, &data.Trx, data.Token)
	} else {
		_, err = updateTransaction(db, &trxid, &data.Trx, data.Token)

	}

	if err != nil {
		//log.Printf("Unable to update transaction.  %v", err)
		log.Fatalf("Unable to update transaction %v", err)
		//c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	if len(data.Trx.Details) > 0 {

		_, err = deleteDetailsByOrder(db, &trxid)
		if err != nil {
			log.Fatalf("Unable to delete trx detail query %v", err)
			//c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
		//}

		// 	var newId int64 = 0

		err = bulkInsertDetails(db, data.Trx.Details, &trxid)

		if err != nil {
			log.Fatalf("Unable to execute finances query %v", err)
			//log.Printf("Unable to insert transaction details (message from command).  %v", err)
			//c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
	}

	msg := fmt.Sprintf("Loan updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      trxid,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)

}

func LoanUpdate(c *gin.Context) {

	// create the postgres db connection

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var loan ts_loan_create

	err := c.BindJSON(&loan)

	if err != nil {
		log.Fatalf("Unable to decode loan from body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	_, err = loan_update(db, &id, &loan.Loan)

	if err != nil {
		log.Fatalf("Loan update error.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	updatedRows, err := updateTransaction(db, &loan.Trx.ID, &loan.Trx, loan.Token)

	if err != nil {
		//log.Printf("Unable to update transaction.  %v", err)
		log.Fatalf("Unable to update transaction %v", err)
		//c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	if len(loan.Trx.Details) > 0 {

		_, err = deleteDetailsByOrder(db, &id)
		if err != nil {
			log.Fatalf("Unable to delete trx detail query %v", err)
			//c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
		//}

		// 	var newId int64 = 0

		err = bulkInsertDetails(db, loan.Trx.Details, &loan.Trx.ID)

		if err != nil {
			log.Fatalf("Unable to execute finances query %v", err)
			//log.Printf("Unable to insert transaction details (message from command).  %v", err)
			//c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
	}

	msg := fmt.Sprintf("Loan updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

type loan_item struct {
	models.Loan
	Trxs *json.RawMessage `json:"trxs,omitempty"`
}

func loan_get_item(db *sql.DB, id *int64) (loan_item, error) {

	var p loan_item
	sb := strings.Builder{}
	sbTrxDetail := strings.Builder{}
	sbTrx := strings.Builder{}

	sbTrxDetail.WriteString("WITH RECURSIVE rs AS (")
	sbTrxDetail.WriteString(" SELECT 1 as group, d.trx_id, d.id, d.code_id, d.debt, d.cred")
	sbTrxDetail.WriteString(" FROM trx_detail AS d")
	sbTrxDetail.WriteString(" INNER JOIN trx ON trx.id = d.trx_id")
	sbTrxDetail.WriteString(" INNER JOIN acc_code AS c ON c.id = d.code_id")
	//	sbTrxDetail.WriteString(" INNER JOIN acc_type AS e ON e.id = c.type_id")
	sbTrxDetail.WriteString(" WHERE c.type_id = 11")
	sbTrxDetail.WriteString(")\n")

	sbTrxDetail.WriteString(`SELECT rs.group AS "groupId", rs.id, rs.trx_id AS "trxId", rs.code_id AS "codeId", rs.debt, rs.cred`)
	sbTrxDetail.WriteString(", sum(rs.debt - rs.cred) OVER (ORDER BY rs.group, rs.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) as saldo")
	sbTrxDetail.WriteString(" FROM rs")
	sbTrxDetail.WriteString(" INNER JOIN trx ON trx.id = rs.trx_id")
	sbTrxDetail.WriteString(" WHERE rs.trx_id = x.id")
	sbTrxDetail.WriteString(" AND (trx.division ='trx-loan' OR trx.division ='trx-angsuran')")

	//log.Printf("%s", sbTrxDetail.String())

	sbTrx.WriteString(`SELECT x.id, x.ref_id AS "refId", x.division, x.descriptions, x.trx_date AS "trxDate", x.memo`)
	sbTrx.WriteString(", ")
	sbTrx.WriteString(fyc.NestQuerySingle(sbTrxDetail.String()))
	sbTrx.WriteString(" AS detail")
	sbTrx.WriteString(" FROM trx AS x")
	sbTrx.WriteString(" WHERE x.ref_id = t.id AND (x.division ='trx-loan' OR x.division ='trx-angsuran')")
	sbTrx.WriteString(" ORDER BY x.id")

	sb.WriteString("SELECT")
	sb.WriteString(" t.id, t.name, t.street, t.city, t.phone, t.cell, t.zip, t.persen")
	sb.WriteString(",")
	sb.WriteString(fyc.NestQuery(sbTrx.String()))
	sb.WriteString(" AS trxs")
	sb.WriteString(" FROM loans AS t")
	sb.WriteString(" WHERE t.id=$1")

	rs := db.QueryRow(sb.String(), id)

	err := rs.Scan(
		&p.ID,
		&p.Name,
		&p.Street,
		&p.City,
		&p.Phone,
		&p.Cell,
		&p.Zip,
		&p.Persen,
		&p.Trxs,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return p, err
	case nil:
		return p, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty on error
	return p, err
}

type loan_all struct {
	ID     int64             `json:"id"`
	Name   string            `json:"name"`
	Street models.NullString `json:"street,omitempty"`
	City   models.NullString `json:"city,omitempty"`
	Phone  models.NullString `json:"phone,omitempty"`
	Cell   models.NullString `json:"cell,omitempty"`
	Zip    models.NullString `json:"zip,omitempty"`
	Persen float32           `json:"persen"`

	TrxID        int64             `json:"trxId"`
	Division     string            `json:"division"`
	Descriptions models.NullString `json:"descriptions"`
	TrxDate      string            `json:"trxDate"`
	Memo         models.NullString `json:"memo"`

	Loan *json.RawMessage `json:"loan,omitempty"`
}

func loan_get_all(db *sql.DB) ([]loan_all, error) {

	var loans []loan_all

	sb := strings.Builder{}
	sb2 := strings.Builder{}

	sb2.WriteString("SELECT ln.id")
	sb2.WriteString(", sum(d.debt) as debt")
	sb2.WriteString(", sum(d.cred) as cred")
	sb2.WriteString(", sum(d.debt + (d.debt * (ln.persen / 100))) as piutang")
	sb2.WriteString(", sum((d.debt + (d.debt * (ln.persen / 100))) - d.cred) as saldo")
	sb2.WriteString(" FROM trx_detail AS d")
	sb2.WriteString(" INNER JOIN acc_code AS c ON c.id = d.code_id")
	//sb2.WriteString(" INNER JOIN acc_type AS e ON e.id = c.type_id")
	sb2.WriteString(" INNER JOIN trx r ON r.id = d.trx_id")
	sb2.WriteString(" INNER JOIN loans ln ON ln.id = r.ref_id")
	sb2.WriteString(" WHERE c.type_id != 11 and ln.id = t.id AND (r.division = 'trx-loan' OR r.division='trx-angsuran')")
	sb2.WriteString(" GROUP BY ln.id")

	sb.WriteString("SELECT t.id id, t.name, t.street, t.city, t.phone, t.cell, t.zip, t.persen,")
	sb.WriteString(" x.id as trx_id, x.division, x.descriptions, x.trx_date, x.memo, ")
	sb.WriteString(fyc.NestQuerySingle(sb2.String()))
	sb.WriteString(" AS details")
	sb.WriteString(" FROM loans AS t")
	sb.WriteString(" INNER JOIN trx AS x on x.ref_id = t.id AND x.division = 'trx-loan'")
	sb.WriteString(" ORDER BY t.serial_num")

	//	log.Println(sb.String())

	rs, err := db.Query(sb.String())

	if err != nil {
		log.Fatalf("Unable to execute finances query %v", err)
		return loans, err
	}

	defer rs.Close()

	for rs.Next() {
		var p loan_all

		err := rs.Scan(
			&p.ID,
			&p.Name,
			&p.Street,
			&p.City,
			&p.Phone,
			&p.Cell,
			&p.Zip,
			&p.Persen,
			&p.TrxID,
			&p.Division,
			&p.Descriptions,
			&p.TrxDate,
			&p.Memo,
			&p.Loan,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		loans = append(loans, p)
	}

	return loans, err
}

func loan_delete(db *sql.DB, id *int64) (int64, error) {

	//log.Printf("%d", id)
	sqlStatement := `DELETE FROM loans WHERE id=$1`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	sqlStatement = `DELETE FROM trx WHERE ref_id=$1 AND (division='trx-loan' OR division='trx-angsuran')`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}

func loan_create(db *sql.DB, loan *models.Loan) (int64, error) {

	sb := strings.Builder{}
	sb.WriteString("INSERT INTO loans")
	sb.WriteString(" (name, street, city, phone, cell, zip, persen)")
	sb.WriteString(" VALUES")
	sb.WriteString(" ($1, $2, $3, $4, $5, $6, $7)")
	sb.WriteString(" RETURNING id")

	var id int64

	err := db.QueryRow(sb.String(),
		loan.Name,
		loan.Street,
		loan.City,
		loan.Phone,
		loan.Cell,
		loan.Zip,
		loan.Persen,
	).Scan(&id)

	return id, err
}

func loan_update(db *sql.DB, id *int64, loan *models.Loan) (int64, error) {
	sb := strings.Builder{}
	sb.WriteString("UPDATE loans SET")
	sb.WriteString(" name=$2, street=$3, city=$4, phone=$5, cell=$6, zip=$7, persen=$8")
	sb.WriteString(" WHERE id=$1")

	res, err := db.Exec(sb.String(),
		id,
		loan.Name,
		loan.Street,
		loan.City,
		loan.Phone,
		loan.Cell,
		loan.Zip,
		loan.Persen,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}
