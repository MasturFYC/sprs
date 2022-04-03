package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type remain_saldo struct {
	ID    float64 `json:"id"`
	Name  string  `json:"name"`
	Debt  float64 `json:"debt"`
	Cred  float64 `json:"cred"`
	Saldo float64 `json:"saldo"`
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
	/*

	   union all

	   	select g.id, g.name,
	   		coalesce(sum(d.debt),0) debt,
	   		coalesce(sum(d.cred),0) cred,
	   		coalesce(sum(d.debt-d.cred), 0) as saldo
	   		from trx_detail d
	   		inner join acc_code c on c.id = d.code_id
	   		inner join acc_type t on t.id = c.type_id
	   		inner join acc_group g on g.id = t.group_id
	   		WHERE g.id != 1
	   		group by g.id
	*/
	b := strings.Builder{}
	b.WriteString("WITH RECURSIVE rs AS (\n")
	b.WriteString("	SELECT c.id, c.name,")
	b.WriteString("	COALESCE(sum(d.debt), 0) AS debt,")
	b.WriteString("	COALESCE(sum(d.cred), 0) AS cred")
	b.WriteString("	FROM trx_detail AS d")
	b.WriteString("	RIGHT JOIN acc_code AS c on c.id = d.code_id")
	b.WriteString("	INNER JOIN acc_type AS t on t.id = c.type_id")
	b.WriteString("	WHERE t.id = 11")
	b.WriteString("	GROUP BY t.group_id, c.id)\n")
	b.WriteString("SELECT	t.id, t.name, t.debt, t.cred,")
	b.WriteString("	t.debt - t.cred AS saldo")
	b.WriteString("	FROM rs t;")

	rs, err := Sql().Query(b.String())

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
			&p.Saldo,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		saldos = append(saldos, p)
	}

	return saldos, err
}
