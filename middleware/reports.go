package middleware

import (
	"encoding/json"
	"log"
	"strings"

	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

type reportTrxBase struct {
	ID    int32   `json:"id"`
	Name  string  `json:"name"`
	Debt  float64 `json:"debt"`
	Cred  float64 `json:"cred"`
	Saldo float64 `json:"saldo"`
}

type reportMonth struct {
	GroupID int32 `json:"groupId"`
	reportTrxBase
	Types []reportType `json:"types,omitempty"`
}

type reportType struct {
	reportTrxBase
	Accounts []reportAccount `json:"accounts,omitempty"`
}

type reportAccount struct {
	reportTrxBase
	TrxDate string `json:"trxDate"`
}

/*
// func GetRepotTrxByTypeMonth(w http.ResponseWriter, r *http.Request) {
// 	EnableCors(&w)

// 	params := mux.Vars(r)

// 	m, _ := strconv.Atoi(params["month"])
// 	//m2, err := strconv.Atoi(params["month2"])
// 	y, _ := strconv.Atoi(params["year"])

// 	result, _ := strconv.ParseInt(params["type"], 10, 32)
// 	acc_type := int32(result)

// 	rpt, err := get_report_trx_by_type_month(&acc_type, &m, &y)

// 	if err != nil || len(rpt) == 0 {
// 		//log.Printf("Unable to get all account codes. %v", err)
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 		//var test []models.AccCode
// 		//json.NewEncoder(w).Encode(test)
// 		//return
// 	}

// 	json.NewEncoder(w).Encode(&rpt)
// }

// func GetRepotTrxByAccountMonth(w http.ResponseWriter, r *http.Request) {
// 	EnableCors(&w)

// 	params := mux.Vars(r)

// 	m, _ := strconv.Atoi(params["month"])
// 	//m2, err := strconv.Atoi(params["month2"])
// 	y, _ := strconv.Atoi(params["year"])

// 	acc_type, _ := strconv.Atoi(params["type"])
// 	acc_id, _ := strconv.Atoi(params["acc"])

// 	rpt, err := get_trx_details_by_acc(&acc_id, &acc_type, &m, &y)

// 	if err != nil || len(rpt) == 0 {
// 		//log.Printf("Unable to get all account codes. %v", err)
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 		//var test []models.AccCode
// 		//json.NewEncoder(w).Encode(test)
// 		//return
// 	}

// 	json.NewEncoder(w).Encode(&rpt)
// }
*/

func GetRepotTrxByMonth(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	m, _ := strconv.Atoi(params["month"])
	//m2, err := strconv.Atoi(params["month2"])
	y, err := strconv.Atoi(params["year"])

	if err != nil {
		//log.Printf("Unable to convert the string into int.  %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rpt, err := get_report_trx_by_month(&m, &y)

	if err != nil || len(rpt) == 0 {
		//log.Printf("Unable to get all account codes. %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
		//var test []models.AccCode
		//json.NewEncoder(w).Encode(test)
		//return
	}

	json.NewEncoder(w).Encode(&rpt)
}

func get_report_trx_by_month(m *int, y *int) ([]reportMonth, error) {

	var reports []reportMonth

	b := strings.Builder{}

	b.WriteString("WITH RECURSIVE rs AS (\n")
	b.WriteString("SELECT 0 AS group, 0 AS id, 'Saldo awal' AS name,")
	// -- COALESCE(SUM(d.debt), 0) AS debt,
	b.WriteString(" 0 AS debt,")
	b.WriteString(" COALESCE(SUM(d.debt - d.cred), 0) AS cred")
	b.WriteString(" FROM trx_detail d")
	b.WriteString(" INNER JOIN trx x on x.id = d.trx_id")
	b.WriteString(" INNER JOIN acc_code c on c.id = d.code_id")
	//-- INNER JOIN acc_type t on t.id = c.type_id
	b.WriteString(" WHERE c.receivable_option = 1 AND")
	b.WriteString(" EXTRACT(MONTH FROM x.trx_date) < $1 AND")
	b.WriteString(" EXTRACT(YEAR  FROM x.trx_date) = $2")

	b.WriteString("\nUNION ALL\n")

	b.WriteString("SELECT 1 as group, t.id, t.name,")
	b.WriteString(" COALESCE(sum(d.debt), 0) AS debt,")
	b.WriteString(" COALESCE(sum(d.cred), 0) AS cred")
	b.WriteString(" FROM trx_detail d")
	b.WriteString(" INNER JOIN trx x on x.id = d.trx_id")
	b.WriteString(" INNER JOIN acc_code c on c.id = d.code_id")
	b.WriteString(" INNER JOIN acc_type t on t.id = c.type_id")
	b.WriteString(" WHERE c.receivable_option != 1 AND")
	b.WriteString(" extract(MONTH FROM x.trx_date) = $1")
	b.WriteString(" AND extract(YEAR  FROM x.trx_date) = $2")
	b.WriteString(" GROUP BY t.id)\n")
	b.WriteString("SELECT")
	b.WriteString(" t.group,")
	b.WriteString(" t.id,")
	b.WriteString(" t.name,")
	b.WriteString(" t.debt, t.cred,")
	b.WriteString(" SUM(t.cred - t.debt)")
	b.WriteString(" OVER (ORDER BY t.group, t.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) as saldo")
	b.WriteString(" FROM rs t")
	b.WriteString(" ORDER BY t.group, t.id;")

	log.Println(b.String())
	rs, err := Sql().Query(b.String(), m, y)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var o reportMonth

		err := rs.Scan(
			&o.GroupID,
			&o.ID,
			&o.Name,
			&o.Debt,
			&o.Cred,
			&o.Saldo,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		if o.ID > 0 {
			t, _ := get_report_trx_by_type_month(&o.ID, m, y)
			o.Types = t
		}

		reports = append(reports, o)
	}

	return reports, err
}

func get_report_trx_by_type_month(group_id *int32, m *int, y *int) ([]reportType, error) {

	var reports []reportType

	var sqlStatement = `WITH RECURSIVE rs AS (
		SELECT 
			c.id as id, c.name, 0, 0
			-- SUM(d.debt) debt, SUM(d.cred) cred
		FROM trx_detail d
		INNER JOIN trx x on x.id = d.trx_id
		INNER JOIN acc_code c on c.id = d.code_id
		INNER JOIN acc_type t on t.id = c.type_id
		WHERE c.receivable_option != 1 AND 
			t.id = $1 AND
			EXTRACT(MONTH FROM x.trx_date) = $2	AND
			EXTRACT(YEAR  FROM x.trx_date) = $3
		GROUP BY c.id, c.name
	)

	SELECT
		t.id,
		t.name,
		0,
		-- t.debt,
		0,
		--t.cred,
		0
		-- SUM(t.cred + t.debt) OVER (ORDER BY t.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS saldo
	FROM rs t
	ORDER BY t.id;`

	rs, err := Sql().Query(sqlStatement, group_id, m, y)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var o reportType

		err := rs.Scan(
			&o.ID,
			&o.Name,
			&o.Debt,
			&o.Cred,
			&o.Saldo,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		t, _ := get_trx_details_by_acc(&o.ID, group_id, m, y)
		o.Accounts = t

		reports = append(reports, o)
	}

	return reports, err
}

func get_trx_details_by_acc(acc *int32, group_id *int32, m *int, y *int) ([]reportAccount, error) {

	var reports []reportAccount

	var sqlStatement = `WITH RECURSIVE rs AS (
		SELECT x.id, x.trx_date, x.descriptions as name, SUM(d.debt) debt, SUM(d.cred) cred
		from trx_detail d
		INNER JOIN trx x on x.id = d.trx_id
		INNER JOIN acc_code c on c.id = d.code_id
		INNER JOIN acc_type t on t.id = c.type_id
		WHERE c.id = $1 AND
			t.id = $2 AND
			EXTRACT(MONTH FROM x.trx_date) = $3	AND
			EXTRACT(YEAR  FROM x.trx_date) = $4
		GROUP BY x.id
	)

	SELECT
		t.id,
		t.trx_date,
		t.name,
		t.debt,
		t.cred,
		SUM(t.debt - t.cred)
		OVER (ORDER BY t.id, t.trx_date ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS saldo
	FROM rs t
	ORDER BY t.id, t.trx_date;`

	rs, err := Sql().Query(sqlStatement, acc, group_id, m, y)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var o reportAccount

		err := rs.Scan(
			&o.ID,
			&o.TrxDate,
			&o.Name,
			&o.Debt,
			&o.Cred,
			&o.Saldo,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		reports = append(reports, o)
	}

	return reports, err
}
