package reports

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"fyc.com/sprs/middleware"
	"github.com/gorilla/mux"
)

type order_invoiced struct {
	ID int64 `json:"id"`

	Name      string  `json:"name"`
	OrderAt   string  `json:"orderAt"`
	BtFinance float64 `json:"btFinance"`
	BtPercent float32 `json:"btPercent"`
	BtMatel   float64 `json:"btMatel"`
	IsStnk    bool    `json:"isStnk"`
	StnkPrice float64 `json:"stnkPrice"`
	Status    int     `json:"status"`
	FinanceId int     `json:"financeId"`

	Branch  json.RawMessage  `json:"branch,omitempty"`
	Unit    *json.RawMessage `json:"unit,omitempty"`
	Finance json.RawMessage  `json:"finance,omitempty"`
}

func ReportOrder(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	params := mux.Vars(r)

	financeId, _ := strconv.Atoi(params["finance"])
	branchId, _ := strconv.Atoi(params["branch"])
	typeId, _ := strconv.Atoi(params["type"])
	month, _ := strconv.Atoi(params["month"])
	year, _ := strconv.Atoi(params["year"])

	dateFrom := params["from"]
	dateTo := params["to"]

	if dateFrom == "-" {
		orders, err := rptOrder1(&financeId, &branchId, &typeId, &month, &year)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(&orders)
		return
	}

	orders, err := rptOrder2(&financeId, &branchId, &typeId, &dateFrom, &dateTo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&orders)

}

func ReportOrderAllWaiting(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	params := mux.Vars(r)

	financeId, _ := strconv.Atoi(params["finance"])
	branchId, _ := strconv.Atoi(params["branch"])
	typeId, _ := strconv.Atoi(params["type"])

	orders, err := rptOrder3(&financeId, &branchId, &typeId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&orders)

}

/*
func rptOrder(filter *report_order_request) ([]order_invoiced, error) {
	var orders []order_invoiced

	var qWheel = middleware.NestQuerySingle(`SELECT name, short_name as "shortName" FROM wheels WHERE id = t.wheel_id`)
	var qMerk = middleware.NestQuerySingle("SELECT name FROM merks WHERE id = t.merk_id")

	var qTye = middleware.NestQuerySingle(fmt.Sprintf(`SELECT t.name, %s AS wheel, %s AS merk FROM types t WHERE t.id = u.type_id`,
		qWheel, qMerk,
	))

	var qUnit = middleware.NestQuerySingle(fmt.Sprintf(`SELECT u.nopol, u.year,
		%s AS type
		FROM units u
		WHERE u.order_id = t.id`,
		qTye))

	var qFinance = middleware.NestQuerySingle(`SELECT f.name, f.short_name "shortName" FROM finances f WHERE f.id = t.finance_id`)
	var qBranch = middleware.NestQuerySingle(`SELECT b.name FROM branchs AS b WHERE b.id = t.branch_id`)

	b := strings.Builder{}

	b.WriteString("WITH RECURSIVE rs AS(")

	// mencari order yg belum diverifikasi
	b.WriteString(" SELECT 0 as status, o.id, o.name, o.order_at, o.bt_finance,")
	b.WriteString(" o.bt_percent, o.bt_matel, o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")

	if filter.Tipe.Where {
		b.WriteString(" INNER JOIN units un ON un.order_id=o.id")
		b.WriteString(" INNER JOIN types tp ON tp.id=un.type_id")
	}

	b.WriteString(" WHERE o.verified_by IS NULL")

	if filter.MonthOrder.Where {
		b.WriteString(fmt.Sprintf(" AND (EXTRACT(MONTH from o.order_at)=%d AND EXTRACT(YEAR from o.order_at)=%d)",
			filter.MonthOrder.Month, filter.MonthOrder.Year))
	}

	if filter.DateOrder.Where {
		b.WriteString(fmt.Sprintf(" AND (o.order_at >= TO_DATE('%s', 'YYYY-MM-DD') AND o.order_at <= TO_DATE('%s', 'YYYY-MM-DD'))",
			filter.DateOrder.From, filter.DateOrder.To))
	}

	if filter.Finance.Where {
		b.WriteString(fmt.Sprintf(" AND (o.finance_id=%d)", filter.Finance.ID))
	}

	if filter.Branch.Where {
		b.WriteString(fmt.Sprintf(" AND (o.branch_id=%d)", filter.Branch.ID))
	}

	if filter.Tipe.Where {
		b.WriteString(fmt.Sprintf(" AND (tp.wheel_id=%d)", filter.Tipe.ID))
	}

	b.WriteString(" UNION ALL")

	// mencari order yg sudah dicairkan
	b.WriteString(" SELECT 1 as status, o.id, v.id::text as name, v.invoice_at,")
	b.WriteString(" o.bt_finance, v.ppn AS bt_percent,")
	b.WriteString(" o.bt_finance - (o.bt_finance * (v.ppn / 100.0)) AS bt_matel,")
	b.WriteString(" o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")

	if filter.Tipe.Where {
		b.WriteString(" INNER JOIN units un ON un.order_id=o.id")
		b.WriteString(" INNER JOIN types tp ON tp.id=un.type_id")
	}

	b.WriteString(" INNER JOIN invoice_details d ON d.order_id = o.id")
	b.WriteString(" INNER JOIN invoices v ON v.id = d.invoice_id")
	b.WriteString(" WHERE o.id IN (SELECT order_id FROM invoice_details)")

	if filter.MonthOrder.Where {
		b.WriteString(fmt.Sprintf(" AND (EXTRACT(MONTH from v.invoice_at)=%d AND EXTRACT(YEAR from v.invoice_at)=%d)",
			filter.MonthOrder.Month, filter.MonthOrder.Year))
	}

	if filter.DateOrder.Where {
		b.WriteString(fmt.Sprintf(" AND (v.invoice_at >= TO_DATE('%s', 'YYYY-MM-DD') AND v.invoice_at <= TO_DATE('%s', 'YYYY-MM-DD'))",
			filter.DateOrder.From, filter.DateOrder.To))
	}

	if filter.Finance.Where {
		b.WriteString(fmt.Sprintf(" AND (o.finance_id=%d)", filter.Finance.ID))
	}

	if filter.Branch.Where {
		b.WriteString(fmt.Sprintf(" AND (o.branch_id=%d)", filter.Branch.ID))
	}

	if filter.Tipe.Where {
		b.WriteString(fmt.Sprintf(" AND (tp.wheel_id=%d)", filter.Tipe.ID))
	}

	b.WriteString(" UNION ALL")

	// mencari orders yg belum dicairkan tapi sudah diverifikasi
	b.WriteString(" SELECT 2 as status, o.id, o.name, o.order_at, o.bt_finance,")
	b.WriteString(" o.bt_percent, o.bt_matel, o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")

	if filter.Tipe.Where {
		b.WriteString(" INNER JOIN units un ON un.order_id=o.id")
		b.WriteString(" INNER JOIN types tp ON tp.id=un.type_id")
	}

	b.WriteString(" WHERE o.id NOT IN (SELECT order_id FROM invoice_details)")
	b.WriteString(" AND o.verified_by IS NOT NULL")

	if filter.MonthOrder.Where {
		b.WriteString(fmt.Sprintf(" AND (EXTRACT(MONTH from o.order_at)=%d AND EXTRACT(YEAR from o.order_at)=%d)",
			filter.MonthOrder.Month, filter.MonthOrder.Year))
	}

	if filter.DateOrder.Where {
		b.WriteString(fmt.Sprintf(" AND (o.order_at >= TO_DATE('%s', 'YYYY-MM-DD') AND o.order_at <= TO_DATE('%s', 'YYYY-MM-DD'))",
			filter.DateOrder.From, filter.DateOrder.To))
	}

	if filter.Finance.Where {
		b.WriteString(fmt.Sprintf(" AND (o.finance_id=%d)", filter.Finance.ID))
	}

	if filter.Branch.Where {
		b.WriteString(fmt.Sprintf(" AND (o.branch_id=%d)", filter.Branch.ID))
	}

	if filter.Tipe.Where {
		b.WriteString(fmt.Sprintf(" AND (tp.wheel_id=%d)", filter.Tipe.ID))
	}

	b.WriteString(")")

	b.WriteString(" SELECT t.status, t.id, t.name, t.order_at, t.bt_finance,")
	b.WriteString(" t.bt_percent, t.bt_matel, t.is_stnk, t.stnk_price, t.finance_id,")
	b.WriteString(qBranch)
	b.WriteString(" AS branch, ")
	b.WriteString(qUnit)
	b.WriteString(" AS unit, ")
	b.WriteString(qFinance)
	b.WriteString(" AS finance ")
	b.WriteString(" FROM rs AS t")
	b.WriteString(" ORDER BY t.status DESC, t.order_at")

	rs, err := middleware.Sql().Query(b.String())

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p order_invoiced

		err := rs.Scan(&p.Status, &p.ID, &p.Name, &p.OrderAt, &p.BtFinance,
			&p.BtPercent, &p.BtMatel,
			&p.IsStnk, &p.StnkPrice,
			&p.FinanceId,
			&p.Branch,
			&p.Unit,
			&p.Finance,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		orders = append(orders, p)
	}

	return orders, err
}
*/
func create_query() (string, string, string) {

	var qWheel = middleware.NestQuerySingle(`SELECT name, short_name as "shortName" FROM wheels WHERE id = t.wheel_id`)
	var qMerk = middleware.NestQuerySingle("SELECT name FROM merks WHERE id = t.merk_id")

	var qTye = middleware.NestQuerySingle(fmt.Sprintf(`SELECT t.name, %s AS wheel, %s AS merk FROM types t WHERE t.id = u.type_id`,
		qWheel, qMerk,
	))

	var qUnit = middleware.NestQuerySingle(fmt.Sprintf(`SELECT u.nopol, u.year,
		%s AS type
		FROM units u
		WHERE u.order_id = t.id`,
		qTye))

	var qFinance = middleware.NestQuerySingle(`SELECT f.name, f.short_name "shortName" FROM finances f WHERE f.id = t.finance_id`)
	var qBranch = middleware.NestQuerySingle(`SELECT b.name FROM branchs AS b WHERE b.id = t.branch_id`)

	return qBranch, qFinance, qUnit
}

func rptOrder1(financeId *int, branchId *int, typeId *int, month *int, year *int) ([]order_invoiced, error) {
	var orders []order_invoiced

	qBranch, qFinance, qUnit := create_query()

	b := strings.Builder{}

	b.WriteString("WITH RECURSIVE rs AS(")

	// mencari order yg belum diverifikasi
	b.WriteString(" SELECT 0 as status, o.id, o.name, o.order_at, o.bt_finance,")
	b.WriteString(" o.bt_percent, o.bt_matel, o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")
	b.WriteString(" INNER JOIN units un ON un.order_id=o.id")
	b.WriteString(" INNER JOIN types tp ON tp.id=un.type_id")
	b.WriteString(" WHERE o.verified_by IS NULL")
	b.WriteString(" AND (o.finance_id=$1 OR 0=$1)")
	b.WriteString(" AND (o.branch_id=$2 OR 0=$2)")
	b.WriteString(" AND (tp.wheel_id=$3 OR 0=$3)")
	b.WriteString(" AND (EXTRACT(MONTH from o.order_at)=$4 AND EXTRACT(YEAR from o.order_at)=$5 OR 0=$4)")

	b.WriteString("\n\nUNION ALL\n\n")

	// mencari order yg sudah dicairkan
	b.WriteString(" SELECT 1 as status, o.id, o.name, o.order_at")
	b.WriteString(", o.bt_finance, o.bt_percent, o.bt_matel")
	b.WriteString(", o.branch_id, o.finance_id")
	b.WriteString(", o.is_stnk, o.stnk_price")
	b.WriteString(" FROM orders AS o")
	b.WriteString(" INNER JOIN invoice_details d ON d.order_id = o.id")
	b.WriteString(" INNER JOIN invoices v ON v.id = d.invoice_id")
	b.WriteString(" INNER JOIN units un ON un.order_id=o.id")
	b.WriteString(" INNER JOIN types tp ON tp.id=un.type_id")
	b.WriteString(" WHERE (o.finance_id=$1 OR 0=$1)")
	b.WriteString(" AND (o.branch_id=$2 OR 0=$2)")
	b.WriteString(" AND (tp.wheel_id=$3 OR 0=$3)")
	b.WriteString(" AND o.id IN (SELECT order_id FROM invoice_details)")
	//	b.WriteString(" AND ((v.invoice_at>=TO_DATE($4, 'YYYY-MM-DD') AND v.invoice_at<=TO_DATE($5, 'YYYY-MM-DD')) OR (''=$4 AND ''=$5))")
	b.WriteString(" AND (EXTRACT(MONTH from o.order_at)=$4 AND EXTRACT(YEAR from o.order_at)=$5 OR 0=$4)")

	b.WriteString("\n\nUNION ALL\n\n")

	// mencari orders yg belum dicairkan tapi sudah diverifikasi
	b.WriteString(" SELECT 2 as status, o.id, o.name, o.order_at, o.bt_finance,")
	b.WriteString(" o.bt_percent, o.bt_matel, o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")
	b.WriteString(" INNER JOIN units un ON un.order_id=o.id")
	b.WriteString(" INNER JOIN types tp ON tp.id=un.type_id")
	b.WriteString(" WHERE (o.finance_id=$1 OR 0=$1)")
	b.WriteString(" AND (o.branch_id=$2 OR 0=$2)")
	b.WriteString(" AND (tp.wheel_id=$3 OR 0=$3)")
	b.WriteString(" AND o.verified_by IS NOT NULL")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM invoice_details)")
	b.WriteString(" AND (EXTRACT(MONTH from o.order_at)=$4 AND EXTRACT(YEAR from o.order_at)=$5 OR 0=$4)")

	b.WriteString(")")

	b.WriteString(" SELECT t.status, t.id, t.name, t.order_at, t.bt_finance,")
	b.WriteString(" t.bt_percent, t.bt_matel, t.is_stnk, t.stnk_price, t.finance_id,")
	b.WriteString(qBranch)
	b.WriteString(" AS branch, ")
	b.WriteString(qFinance)
	b.WriteString(" AS finance, ")
	b.WriteString(qUnit)
	b.WriteString(" AS unit")
	b.WriteString(" FROM rs AS t")
	b.WriteString(" ORDER BY t.order_at, t.id")

	rs, err := middleware.Sql().Query(b.String(), financeId, branchId, typeId, month, year)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p order_invoiced

		err := rs.Scan(&p.Status, &p.ID, &p.Name, &p.OrderAt, &p.BtFinance,
			&p.BtPercent, &p.BtMatel,
			&p.IsStnk, &p.StnkPrice,
			&p.FinanceId,
			&p.Branch,
			&p.Finance,
			&p.Unit,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		orders = append(orders, p)
	}

	return orders, err
}

func rptOrder2(financeId *int, branchId *int, typeId *int, dateFrom *string, dateTo *string) ([]order_invoiced, error) {
	var orders []order_invoiced

	qBranch, qFinance, qUnit := create_query()

	b := strings.Builder{}

	b.WriteString("WITH RECURSIVE rs AS(")

	// mencari order yg belum diverifikasi
	b.WriteString(" SELECT 0 as status, o.id, o.name, o.order_at, o.bt_finance,")
	b.WriteString(" o.bt_percent, o.bt_matel, o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")
	b.WriteString(" INNER JOIN units un ON un.order_id=o.id")
	b.WriteString(" INNER JOIN types tp ON tp.id=un.type_id")
	b.WriteString(" WHERE o.verified_by IS NULL")
	b.WriteString(" AND (o.finance_id=$1 OR 0=$1)")
	b.WriteString(" AND (o.branch_id=$2 OR 0=$2)")
	b.WriteString(" AND (tp.wheel_id=$3 OR 0=$3)")
	b.WriteString(" AND (o.order_at>=TO_DATE($4, 'YYYY-MM-DD') AND o.order_at<=TO_DATE($5,'YYYY-MM-DD'))")

	b.WriteString("\n\nUNION ALL\n\n")

	// mencari order yg sudah dicairkan
	b.WriteString(" SELECT 1 as status, o.id, o.name, o.order_at")
	b.WriteString(", o.bt_finance, o.bt_percent, o.bt_matel")
	b.WriteString(", o.branch_id, o.finance_id")
	b.WriteString(", o.is_stnk, o.stnk_price")
	b.WriteString(" FROM orders AS o")
	b.WriteString(" INNER JOIN invoice_details d ON d.order_id = o.id")
	b.WriteString(" INNER JOIN invoices v ON v.id = d.invoice_id")
	b.WriteString(" INNER JOIN units un ON un.order_id=o.id")
	b.WriteString(" INNER JOIN types tp ON tp.id=un.type_id")
	b.WriteString(" WHERE (o.finance_id=$1 OR 0=$1)")
	b.WriteString(" AND (o.branch_id=$2 OR 0=$2)")
	b.WriteString(" AND (tp.wheel_id=$3 OR 0=$3)")
	b.WriteString(" AND o.id IN (SELECT order_id FROM invoice_details)")
	//	b.WriteString(" AND ((v.invoice_at>=TO_DATE($4, 'YYYY-MM-DD') AND v.invoice_at<=TO_DATE($5, 'YYYY-MM-DD')) OR (''=$4 AND ''=$5))")
	b.WriteString(" AND (o.order_at>=TO_DATE($4, 'YYYY-MM-DD') AND o.order_at<=TO_DATE($5,'YYYY-MM-DD'))")

	b.WriteString("\n\nUNION ALL\n\n")

	// mencari orders yg belum dicairkan tapi sudah diverifikasi
	b.WriteString(" SELECT 2 as status, o.id, o.name, o.order_at, o.bt_finance,")
	b.WriteString(" o.bt_percent, o.bt_matel, o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")
	b.WriteString(" INNER JOIN units un ON un.order_id=o.id")
	b.WriteString(" INNER JOIN types tp ON tp.id=un.type_id")
	b.WriteString(" WHERE (o.finance_id=$1 OR 0=$1)")
	b.WriteString(" AND (o.branch_id=$2 OR 0=$2)")
	b.WriteString(" AND (tp.wheel_id=$3 OR 0=$3)")
	b.WriteString(" AND o.verified_by IS NOT NULL")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM invoice_details)")
	b.WriteString(" AND (o.order_at>=TO_DATE($4, 'YYYY-MM-DD') AND o.order_at<=TO_DATE($5,'YYYY-MM-DD'))")

	b.WriteString(")")

	b.WriteString(" SELECT t.status, t.id, t.name, t.order_at, t.bt_finance,")
	b.WriteString(" t.bt_percent, t.bt_matel, t.is_stnk, t.stnk_price, t.finance_id,")
	b.WriteString(qBranch)
	b.WriteString(" AS branch, ")
	b.WriteString(qFinance)
	b.WriteString(" AS finance, ")
	b.WriteString(qUnit)
	b.WriteString(" AS unit ")
	b.WriteString(" FROM rs AS t")
	b.WriteString(" ORDER BY t.order_at, t.id")

	rs, err := middleware.Sql().Query(b.String(), financeId, branchId, typeId, dateFrom, dateTo)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p order_invoiced

		err := rs.Scan(&p.Status, &p.ID, &p.Name, &p.OrderAt, &p.BtFinance,
			&p.BtPercent, &p.BtMatel,
			&p.IsStnk, &p.StnkPrice,
			&p.FinanceId,
			&p.Branch,
			&p.Finance,
			&p.Unit,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		orders = append(orders, p)
	}

	return orders, err
}

/*
Mencari semua data order yang belum dicairkan
termasuk yang belum diverifikasi
*/

func rptOrder3(financeId *int, branchId *int, typeId *int) ([]order_invoiced, error) {
	var orders []order_invoiced

	qBranch, qFinance, qUnit := create_query()

	b := strings.Builder{}

	b.WriteString("WITH RECURSIVE rs AS(")

	// mencari order yg belum diverifikasi
	b.WriteString(" SELECT 0 as status, o.id, o.name, o.order_at, o.bt_finance,")
	b.WriteString(" o.bt_percent, o.bt_matel, o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")
	b.WriteString(" INNER JOIN units un ON un.order_id=o.id")
	b.WriteString(" INNER JOIN types tp ON tp.id=un.type_id")
	b.WriteString(" WHERE o.verified_by IS NULL")
	b.WriteString(" AND (o.finance_id=$1 OR 0=$1)")
	b.WriteString(" AND (o.branch_id=$2 OR 0=$2)")
	b.WriteString(" AND (tp.wheel_id=$3 OR 0=$3)")

	b.WriteString("\n\nUNION ALL\n\n")

	// mencari orders yg belum dicairkan tapi sudah diverifikasi
	b.WriteString(" SELECT 2 as status, o.id, o.name, o.order_at, o.bt_finance,")
	b.WriteString(" o.bt_percent, o.bt_matel, o.branch_id, o.finance_id,")
	b.WriteString(" o.is_stnk, o.stnk_price ")
	b.WriteString(" FROM orders AS o")
	b.WriteString(" INNER JOIN units un ON un.order_id=o.id")
	b.WriteString(" INNER JOIN types tp ON tp.id=un.type_id")
	b.WriteString(" WHERE (o.finance_id=$1 OR 0=$1)")
	b.WriteString(" AND (o.branch_id=$2 OR 0=$2)")
	b.WriteString(" AND (tp.wheel_id=$3 OR 0=$3)")
	b.WriteString(" AND o.verified_by IS NOT NULL")
	b.WriteString(" AND o.id NOT IN (SELECT order_id FROM invoice_details)")

	b.WriteString(")")

	b.WriteString(" SELECT t.status, t.id, t.name, t.order_at, t.bt_finance,")
	b.WriteString(" t.bt_percent, t.bt_matel, t.is_stnk, t.stnk_price, t.finance_id,")
	b.WriteString(qBranch)
	b.WriteString(" AS branch, ")
	b.WriteString(qFinance)
	b.WriteString(" AS finance, ")
	b.WriteString(qUnit)
	b.WriteString(" AS unit ")
	b.WriteString(" FROM rs AS t")
	b.WriteString(" ORDER BY t.order_at, t.id")

	rs, err := middleware.Sql().Query(b.String(), financeId, branchId, typeId)

	if err != nil {
		log.Printf("Unable to execute orderes query %v", err)
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {
		var p order_invoiced

		err := rs.Scan(&p.Status, &p.ID, &p.Name, &p.OrderAt, &p.BtFinance,
			&p.BtPercent, &p.BtMatel,
			&p.IsStnk, &p.StnkPrice,
			&p.FinanceId,
			&p.Branch,
			&p.Finance,
			&p.Unit,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		orders = append(orders, p)
	}

	return orders, err
}
