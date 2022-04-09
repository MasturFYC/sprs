package controller

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LabaRugi struct {
	RowIndex int     `json:"index"`
	Title    string  `json:"title"`
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

	sb := strings.Builder{}

	sb.WriteString("WITH RECURSIVE rs AS (")
	sb.WriteString("SELECT 1 as row_index, 'Modal dikeluarkan' as title, x.division, sum(d.debt) debt, sum(d.cred) cred, 0 as profit ")
	// sb.WriteString(", SUM(d.debt - d.cred) as saldo")
	sb.WriteString(" FROM trx_detail d")
	sb.WriteString(" INNER JOIN trx x ON x.id = d.trx_id")
	sb.WriteString(" INNER JOIN acc_code c ON c.id = d.code_id")
	sb.WriteString(" INNER JOIN acc_type e ON e.id = c.type_id")
	sb.WriteString(" INNER JOIN acc_group g ON g.id = e.group_id")
	//sb.WriteString(" WHERE x.division = 'trx-order'")
	sb.WriteString(" WHERE c.type_id != 11")
	//sb.WriteString(" AND g.id = 5")
	sb.WriteString(" AND e.id = 55")
	sb.WriteString(" AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD')")
	sb.WriteString(" AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))")
	sb.WriteString(" GROUP BY x.division")

	sb.WriteString("\nUNION ALL\n")

	sb.WriteString("SELECT 2 as row_index, 'Pendapatan Invoice' as title, x.division, sum(o.bt_matel) debt, sum(d.cred) cred, sum(d.cred - o.bt_matel) as profit ")
	// sb.WriteString(", SUM(d.debt - d.cred) as saldo")
	sb.WriteString(" FROM trx_detail d")
	sb.WriteString(" INNER JOIN trx x ON x.id = d.trx_id")
	sb.WriteString(" INNER JOIN acc_code c ON c.id = d.code_id")
	sb.WriteString(" INNER JOIN acc_type e ON e.id = c.type_id")
	sb.WriteString(" INNER JOIN acc_group g ON g.id = e.group_id")
	sb.WriteString(" INNER JOIN invoices v ON v.id = x.ref_id")
	sb.WriteString(" INNER JOIN invoice_details vd ON vd.invoice_id = v.id")
	sb.WriteString(" INNER JOIN orders o ON o.id = vd.order_id")
	//sb.WriteString(" WHERE x.division = 'trx-order'")
	sb.WriteString(" WHERE c.type_id != 11")
	sb.WriteString(" AND x.division = 'trx-invoice'")
	sb.WriteString(" AND g.id = 4")
	sb.WriteString(" AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD')")
	sb.WriteString(" AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))")
	sb.WriteString(" GROUP BY x.division")

	sb.WriteString("\nUNION ALL\n")

	sb.WriteString("SELECT 2 as row_index, 'Pendapatan cicilan unit' as title, x.division")
	sb.WriteString(", sum(d.cred - (d.cred * ((o.bt_finance - o.bt_matel) / o.bt_finance))) debt")
	sb.WriteString(", sum(d.cred) cred")
	sb.WriteString(", sum(d.cred * ((o.bt_finance - o.bt_matel) / o.bt_finance)) as profit ")
	// sb.WriteString(", SUM(d.debt - d.cred) as saldo")
	sb.WriteString(" FROM trx_detail d")
	sb.WriteString(" INNER JOIN trx x ON x.id = d.trx_id")
	sb.WriteString(" INNER JOIN acc_code c ON c.id = d.code_id")
	sb.WriteString(" INNER JOIN acc_type e ON e.id = c.type_id")
	sb.WriteString(" INNER JOIN acc_group g ON g.id = e.group_id")
	sb.WriteString(" INNER JOIN lents v ON v.order_id = x.ref_id")
	sb.WriteString(" INNER JOIN orders o ON o.id = v.order_id")
	//sb.WriteString(" WHERE x.division = 'trx-order'")
	sb.WriteString(" WHERE c.type_id != 11")
	sb.WriteString(" AND x.division = 'trx-cicilan'")
	sb.WriteString(" AND g.id = 4")
	sb.WriteString(" AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD')")
	sb.WriteString(" AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))")
	sb.WriteString(" GROUP BY x.division")

	sb.WriteString("\nUNION ALL\n")

	sb.WriteString("SELECT 2 as row_index, 'Pendapatan angsuran piutang' as title, x.division")
	sb.WriteString(", sum(d.cred - (d.cred * v.persen)) debt")
	sb.WriteString(", sum(d.cred) cred")
	sb.WriteString(", sum(d.cred * v.persen) as profit ")
	// sb.WriteString(", SUM(d.debt - d.cred) as saldo")
	sb.WriteString(" FROM trx_detail d")
	sb.WriteString(" INNER JOIN trx x ON x.id = d.trx_id")
	sb.WriteString(" INNER JOIN acc_code c ON c.id = d.code_id")
	sb.WriteString(" INNER JOIN acc_type e ON e.id = c.type_id")
	sb.WriteString(" INNER JOIN acc_group g ON g.id = e.group_id")
	sb.WriteString(" INNER JOIN loans v ON v.id = x.ref_id")
	//sb.WriteString(" WHERE x.division = 'trx-order'")
	sb.WriteString(" WHERE c.type_id != 11")
	sb.WriteString(" AND x.division = 'trx-angsuran'")
	sb.WriteString(" AND g.id = 4")
	sb.WriteString(" AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD')")
	sb.WriteString(" AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))")
	sb.WriteString(" GROUP BY x.division")

	sb.WriteString("\nUNION ALL\n")

	sb.WriteString("SELECT 3 as row_index, 'Beban Biaya' as title, c.name as division, sum(d.debt) debt, sum(d.cred) cred, 0 as profit ")
	// sb.WriteString(", SUM(d.debt - d.cred) as saldo")
	sb.WriteString(" FROM trx_detail d")
	sb.WriteString(" INNER JOIN trx x ON x.id = d.trx_id")
	sb.WriteString(" INNER JOIN acc_code c ON c.id = d.code_id")
	sb.WriteString(" INNER JOIN acc_type e ON e.id = c.type_id")
	sb.WriteString(" INNER JOIN acc_group g ON g.id = e.group_id")
	//sb.WriteString(" WHERE x.division = 'trx-order'")
	sb.WriteString(" WHERE c.type_id != 11")
	sb.WriteString(" AND g.id = 5")
	sb.WriteString(" AND e.id != 55")
	sb.WriteString(" AND (x.trx_date >= TO_DATE($1, 'YYYY-MM-DD')")
	sb.WriteString(" AND x.trx_date <= TO_DATE($2, 'YYYY-MM-DD'))")
	sb.WriteString(" GROUP BY c.name")

	sb.WriteString(")")
	sb.WriteString("SELECT t.row_index, t.title, t.division, t.debt, t.cred, t.profit")
	// sb.WriteString(" SUM(t.debt - t.cred) as saldo")
	//	sb.WriteString(" OVER (ORDER BY t.division ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) as saldo")
	sb.WriteString(" FROM rs t")
	sb.WriteString(" ORDER BY t.row_index, t.division")

	rs, err := db.Query(sb.String(), fd, td)

	if err != nil {
		log.Printf("Unable to execute laba rugi query %v", err)
		return result, err
	}

	defer rs.Close()

	for rs.Next() {
		var p LabaRugi

		err := rs.Scan(
			&p.RowIndex,
			&p.Title,
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
