package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type remain_saldo struct {
	ID   float64 `json:"id"`
	Name string  `json:"name"`
	Debt float64 `json:"debt"`
	Cred float64 `json:"cred"`
}

func GetRemainSaldo(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	rv, err := get_remain_saldo()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&rv)
}

func get_remain_saldo() ([]remain_saldo, error) {

	var saldos []remain_saldo

	var sqlStatement = `SELECT 
		t.id, t.name, SUM(COALESCE(d.debt,0)) AS debt, SUM(COALESCE(d.cred,0)) as cred
	FROM trx_detail d
	INNER JOIN trx o ON o.id = d.trx_id
	RIGHT JOIN trx_type t ON t.id = o.trx_type_id
	GROUP BY t.id
	ORDER BY t.id`

	rs, err := Sql().Query(sqlStatement)

	if err != nil {
		log.Printf("Unable to execute saldo query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p remain_saldo

		err := rs.Scan(
			&p.ID,
			&p.Name,
			&p.Debt,
			&p.Cred,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		saldos = append(saldos, p)
	}

	return saldos, err
}
