package middleware

import (
	"encoding/json"
	"log"

	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

type reportTrxByMonth struct {
	Group int32   `json:"group"`
	ID    int32   `json:"id"`
	Name  string  `json:"name"`
	Debt  float64 `json:"debt"`
	Cred  float64 `json:"cred"`
	Saldo float64 `json:"saldo"`
}

func GetRepotTrxByAccMonth(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	m, _ := strconv.Atoi(params["month"])
	//m2, err := strconv.Atoi(params["month2"])
	y, _ := strconv.Atoi(params["year"])

	acc, _ := strconv.Atoi(params["acc"])

	rpt, err := get_report_trx_by_acc_month(&acc, &m, &y)

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

func GetRepotTrxByMonth(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	m, err := strconv.Atoi(params["month"])
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

func get_report_trx_by_month(m *int, y *int) ([]reportTrxByMonth, error) {

	var reports []reportTrxByMonth

	var sqlStatement = `with recursive rs as (
		select 0 as group, t.id, 'Saldo awal' as name, 0 debt, sum(d.debt - d.cred) cred
		from trx_detail d
		inner join trx x on x.id = d.trx_id
		inner join acc_code a on a.id = d.acc_code_id
		inner join acc_type t on t.id = a.acc_type_id
		WHERE t.id = 11 AND
		extract(MONTH FROM x.trx_date) < $1 AND
		extract(YEAR  FROM x.trx_date) = $2
		group by t.id
		
		UNION ALL
		
		select 1 as group, t.id, t.name, sum(d.debt) debt, sum(d.cred) cred
		from trx_detail d
		inner join trx x on x.id = d.trx_id
		inner join acc_code a on a.id = d.acc_code_id
		inner join acc_type t on t.id = a.acc_type_id
		where t.id != 11 AND
			extract(MONTH FROM x.trx_date) = $1 
			AND extract(YEAR  FROM x.trx_date) = $2
		group by t.id
		)
		
		select t.group,
			t.id,
			t.name,
			t.debt, t.cred,
			sum(t.cred - t.debt)
			over (order by t.group, t.id rows between unbounded preceding and current row) as saldo
		from rs t    
		order by t.group, t.id;`

	rs, err := Sql().Query(sqlStatement, m, y)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var o reportTrxByMonth

		err := rs.Scan(
			&o.Group,
			&o.ID,
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

func get_report_trx_by_acc_month(acc *int, m *int, y *int) ([]reportTrxByMonth, error) {

	var reports []reportTrxByMonth

	var sqlStatement = `WITH RECURSIVE rs AS (
		SELECT 1 as group, a.id, a.name, SUM(d.debt) debt, SUM(d.cred) cred
		from trx_detail d
		INNER JOIN trx x on x.id = d.trx_id
		INNER JOIN acc_code a on a.id = d.acc_code_id
		INNER JOIN acc_type t on t.id = a.acc_type_id
		WHERE t.id != 11 AND 
			t.id = $1 AND
			EXTRACT(MONTH FROM x.trx_date) = $2	AND
			EXTRACT(YEAR  FROM x.trx_date) = $3
		GROUP BY a.id
	)

	SELECT
		t.group,
		t.id,
		t.name,
		t.debt, t.cred,
		SUM(t.cred + t.debt)
		OVER (ORDER BY t.group, t.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS saldo
	FROM rs t
	ORDER BY t.group, t.id;`

	rs, err := Sql().Query(sqlStatement, acc, m, y)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var o reportTrxByMonth

		err := rs.Scan(
			&o.Group,
			&o.ID,
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
