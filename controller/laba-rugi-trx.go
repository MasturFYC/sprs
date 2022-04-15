package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LabaRugi struct {
	Id       int32   `json:"id"`
	RowIndex int     `json:"index"`
	Title    string  `json:"title"`
	Branch   int     `json:"branch"`
	Division string  `json:"division"`
	Debt     float64 `json:"debt"`
	Cred     float64 `json:"cred"`
	Profit   float64 `json:"profit"`
}

//api/labarugi/bydate/:from/to
func LabaRugiGetByDate(c *gin.Context) {
	db := c.Keys["db"].(*sql.DB)
	fd := c.Param("from")
	td := c.Param("to")

	result, err := laba_rugi_getByDate(db, fd, td)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &result)
}

func laba_rugi_getByDate(db *sql.DB, fd string, td string) ([]LabaRugi, error) {

	var result = make([]LabaRugi, 0)
	/*
		//sb := strings.Builder{}

		// sb.WriteString("WITH RECURSIVE rs AS (")
		// sb.WriteString("SELECT 1 as row_index, 'Modal dikeluarkan' as title, x.division, sum(d.debt) debt, sum(d.cred) cred, 0 as profit ")
		// // sb.WriteString(", SUM(d.debt - d.cred) as saldo")
		// sb.WriteString(" FROM trx_detail d")
		// sb.WriteString(" INNER JOIN trx x ON x.id = d.trx_id")
		// sb.WriteString(" INNER JOIN acc_code c ON c.id = d.code_id")
		// sb.WriteString(" INNER JOIN acc_type e ON e.id = c.type_id")
		// sb.WriteString(" INNER JOIN acc_group g ON g.id = e.group_id")
		// //sb.WriteString(" WHERE x.division = 'trx-order'")
		// sb.WriteString(" WHERE c.type_id != 11")
		// //sb.WriteString(" AND g.id = 5")
		// sb.WriteString(" AND e.id = 55")
		// sb.WriteString(" AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD')")
		// sb.WriteString(" AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))")
		// sb.WriteString(" GROUP BY x.division")

		// sb.WriteString("\nUNION ALL\n")

		// sb.WriteString("SELECT 2 as row_index, 'Pendapatan Invoice' as title, x.division, sum(o.bt_matel) debt, sum(d.cred) cred, sum(d.cred - o.bt_matel) as profit ")
		// // sb.WriteString(", SUM(d.debt - d.cred) as saldo")
		// sb.WriteString(" FROM trx_detail d")
		// sb.WriteString(" INNER JOIN trx x ON x.id = d.trx_id")
		// sb.WriteString(" INNER JOIN acc_code c ON c.id = d.code_id")
		// sb.WriteString(" INNER JOIN acc_type e ON e.id = c.type_id")
		// sb.WriteString(" INNER JOIN acc_group g ON g.id = e.group_id")
		// sb.WriteString(" INNER JOIN invoices v ON v.id = x.ref_id")
		// sb.WriteString(" INNER JOIN invoice_details vd ON vd.invoice_id = v.id")
		// sb.WriteString(" INNER JOIN orders o ON o.id = vd.order_id")
		// //sb.WriteString(" WHERE x.division = 'trx-order'")
		// sb.WriteString(" WHERE c.type_id != 11")
		// sb.WriteString(" AND x.division = 'trx-invoice'")
		// sb.WriteString(" AND g.id = 4")
		// sb.WriteString(" AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD')")
		// sb.WriteString(" AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))")
		// sb.WriteString(" GROUP BY x.division")

		// sb.WriteString("\nUNION ALL\n")

		// sb.WriteString("SELECT 2 as row_index, 'Pendapatan cicilan unit' as title, x.division")
		// sb.WriteString(", sum(d.cred - (d.cred * ((o.bt_finance - o.bt_matel) / o.bt_finance))) debt")
		// sb.WriteString(", sum(d.cred) cred")
		// sb.WriteString(", sum(d.cred * ((o.bt_finance - o.bt_matel) / o.bt_finance)) as profit ")
		// // sb.WriteString(", SUM(d.debt - d.cred) as saldo")
		// sb.WriteString(" FROM trx_detail d")
		// sb.WriteString(" INNER JOIN trx x ON x.id = d.trx_id")
		// sb.WriteString(" INNER JOIN acc_code c ON c.id = d.code_id")
		// sb.WriteString(" INNER JOIN acc_type e ON e.id = c.type_id")
		// sb.WriteString(" INNER JOIN acc_group g ON g.id = e.group_id")
		// sb.WriteString(" INNER JOIN lents v ON v.order_id = x.ref_id")
		// sb.WriteString(" INNER JOIN orders o ON o.id = v.order_id")
		// //sb.WriteString(" WHERE x.division = 'trx-order'")
		// sb.WriteString(" WHERE c.type_id != 11")
		// sb.WriteString(" AND x.division = 'trx-cicilan'")
		// sb.WriteString(" AND g.id = 4")
		// sb.WriteString(" AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD')")
		// sb.WriteString(" AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))")
		// sb.WriteString(" GROUP BY x.division")

		// sb.WriteString("\nUNION ALL\n")

		// sb.WriteString("SELECT 2 as row_index, 'Pendapatan angsuran piutang' as title, x.division")
		// sb.WriteString(", sum(d.cred - (d.cred * v.persen)) debt")
		// sb.WriteString(", sum(d.cred) cred")
		// sb.WriteString(", sum(d.cred * v.persen) as profit ")
		// // sb.WriteString(", SUM(d.debt - d.cred) as saldo")
		// sb.WriteString(" FROM trx_detail d")
		// sb.WriteString(" INNER JOIN trx x ON x.id = d.trx_id")
		// sb.WriteString(" INNER JOIN acc_code c ON c.id = d.code_id")
		// sb.WriteString(" INNER JOIN acc_type e ON e.id = c.type_id")
		// sb.WriteString(" INNER JOIN acc_group g ON g.id = e.group_id")
		// sb.WriteString(" INNER JOIN loans v ON v.id = x.ref_id")
		// //sb.WriteString(" WHERE x.division = 'trx-order'")
		// sb.WriteString(" WHERE c.type_id != 11")
		// sb.WriteString(" AND x.division = 'trx-angsuran'")
		// sb.WriteString(" AND g.id = 4")
		// sb.WriteString(" AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD')")
		// sb.WriteString(" AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))")
		// sb.WriteString(" GROUP BY x.division")

		// sb.WriteString("\nUNION ALL\n")

		// sb.WriteString("SELECT 3 as row_index, 'Beban Biaya' as title, c.name as division, sum(d.debt) debt, sum(d.cred) cred, 0 as profit ")
		// // sb.WriteString(", SUM(d.debt - d.cred) as saldo")
		// sb.WriteString(" FROM trx_detail d")
		// sb.WriteString(" INNER JOIN trx x ON x.id = d.trx_id")
		// sb.WriteString(" INNER JOIN acc_code c ON c.id = d.code_id")
		// sb.WriteString(" INNER JOIN acc_type e ON e.id = c.type_id")
		// sb.WriteString(" INNER JOIN acc_group g ON g.id = e.group_id")
		// //sb.WriteString(" WHERE x.division = 'trx-order'")
		// sb.WriteString(" WHERE c.type_id != 11")
		// sb.WriteString(" AND g.id = 5")
		// sb.WriteString(" AND e.id != 55")
		// sb.WriteString(" AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD')")
		// sb.WriteString(" AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))")
		// sb.WriteString(" GROUP BY c.name")

		// sb.WriteString(")")
		// sb.WriteString("SELECT t.row_index, t.title, t.division, t.debt, t.cred, t.profit")
		// // sb.WriteString(" SUM(t.debt - t.cred) as saldo")
		// //	sb.WriteString(" OVER (ORDER BY t.division ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) as saldo")
		// sb.WriteString(" FROM rs t")
		// sb.WriteString(" ORDER BY t.row_index, t.division")
	*/
	sqlStmt := `WITH RECURSIVE rs AS (
		SELECT
												1 									as row_index,												
												'Modal Piutang'			as title,
												o.branch_id					as branch_id,										
												x.division					as division,
												sum(d.debt) 				as debt,
												sum(d.cred) 				as cred,
												0										as profit
		FROM trx_detail d
		INNER JOIN trx 				as x							ON x.id = d.trx_id
		INNER JOIN acc_code 	as c							ON c.id = d.code_id
		INNER JOIN orders 		as o							ON o.id = x.ref_id
		WHERE c.type_id != 11
			AND c.type_id = 55
			AND c.id != 5512
			AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD') AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))
		GROUP BY o.branch_id, x.division

		UNION ALL

		SELECT
												1 										as row_index,												
												'Modal Pinjaman'			as title,
												1											as branch_id,										
												x.division						as division,
												sum(d.debt) 					as debt,
												sum(d.cred) 					as cred,
												0											as profit
		FROM trx_detail d
		INNER JOIN trx 				as x								ON x.id = d.trx_id
		INNER JOIN acc_code 	as c								ON c.id = d.code_id
		WHERE c.type_id != 11
			AND c.id = 5512
			AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD') AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))
		GROUP BY x.division

		UNION ALL

		SELECT 
										2													as row_index,
										'Pendapatan Invoice' 			as title,
										o.branch_id								as branch_id,
										x.division								as division,
										sum(o.bt_matel)						as debt,
										sum(d.cred) 							as cred,
										sum(d.cred - o.bt_matel) 	as profit 

	 	FROM trx_detail 							as d
			INNER JOIN trx 							as x				ON x.id = d.trx_id
			INNER JOIN acc_code 				as c				ON c.id = d.code_id
			INNER JOIN acc_type 				as e				ON e.id = c.type_id
			INNER JOIN acc_group 				as g				ON g.id = e.group_id
			INNER JOIN invoices 				as v				ON v.id = x.ref_id
			INNER JOIN invoice_details 	as vd				ON vd.invoice_id = v.id
			INNER JOIN orders 					as o 				ON o.id = vd.order_id
		WHERE c.type_id != 11
			AND x.division = 'trx-invoice'
			AND g.id = 4
			AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD') AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))
		GROUP BY o.branch_id, x.division

	UNION ALL

	SELECT 
									2														as row_index,
									'Pendapatan cicilan unit' 	as title,
									o.branch_id									as branch_id,
									x.division									as division,
									sum(d.cred - (d.cred * ((o.bt_finance - o.bt_matel) / o.bt_finance))) as debt,
									sum(d.cred) 								as cred,
									sum(d.cred * ((o.bt_finance - o.bt_matel) / o.bt_finance)) as profit 
	FROM trx_detail 							as d
		INNER JOIN trx 							as x					ON x.id = d.trx_id
		INNER JOIN acc_code 				as c					ON c.id = d.code_id
		INNER JOIN acc_type 				as e 					ON e.id = c.type_id
		INNER JOIN acc_group 				as g 					ON g.id = e.group_id
		INNER JOIN lents 						as v 					ON v.order_id = x.ref_id
		INNER JOIN orders 					as o 					ON o.id = v.order_id
	WHERE c.type_id != 11
		AND x.division = 'trx-cicilan'
		AND g.id = 4
		AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD')
		AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))
	GROUP BY o.branch_id, x.division

	UNION ALL

	SELECT
									2																	as row_index,
									'Pendapatan angsuran piutang' 		as title,
									1																	as branch_id,
									x.division												as division,
									sum(d.cred - (d.cred * v.persen)) as debt,
									sum(d.cred) 											as cred,
									sum(d.cred * v.persen)						as profit 
	 FROM trx_detail							as d
		INNER JOIN trx 							as x 		ON x.id = d.trx_id
		INNER JOIN acc_code 				as c 		ON c.id = d.code_id
		INNER JOIN acc_type 				as e 		ON e.id = c.type_id
		INNER JOIN acc_group 				as g 		ON g.id = e.group_id
		INNER JOIN loans 						as v 		ON v.id = x.ref_id
	 WHERE c.type_id != 11
		AND x.division = 'trx-angsuran'
		AND g.id = 4
		AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD') AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))
	 GROUP BY x.division

	UNION ALL

	SELECT
									3							as row_index,
									'Beban Biaya' as title,
									1							as branch_id,
									c.name 				as division,
									sum(d.debt)		as debt,
									sum(d.cred)		as cred,
									0 						as profit 
	FROM trx_detail 				as d
		INNER JOIN trx 				as x	ON x.id = d.trx_id
		INNER JOIN acc_code 	as c	ON c.id = d.code_id
		INNER JOIN acc_type 	as e	ON e.id = c.type_id
		INNER JOIN acc_group	as g	ON g.id = e.group_id
	WHERE c.type_id != 11
		AND g.id = 5
		AND e.id != 55
		AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD') AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))
	GROUP BY c.name
	)
	SELECT 		ROW_NUMBER () OVER (ORDER BY t.row_index, t.branch_id, t.division) as id,
						t.row_index,
						t.title,
						t.branch_id,
						t.division,
						t.debt,
						t.cred,
						t.profit
	 FROM rs t
	 ORDER BY t.row_index, t.branch_id, t.division
	`

	rs, err := db.Query(sqlStmt, fd, td)

	if err != nil {
		log.Printf("Unable to execute laba rugi query %v", err)
		return result, err
	}

	defer rs.Close()

	for rs.Next() {
		var p LabaRugi

		err := rs.Scan(
			&p.Id,
			&p.RowIndex,
			&p.Title,
			&p.Branch,
			&p.Division,
			&p.Debt,
			&p.Cred,
			&p.Profit,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		result = append(result, p)
	}

	return result, err

}
