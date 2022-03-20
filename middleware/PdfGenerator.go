package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"fyc.com/sprs/models"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Pdf_GenInvoice(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	invoice, err := invoice_get_item(&id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var finance models.Finance
	source := (*json.RawMessage)(&invoice.Finance)
	err = json.Unmarshal(*source, &finance)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var orders []models.Order
	source = (*json.RawMessage)(&invoice.Details)
	err = json.Unmarshal(*source, &orders)

	//log.Println(orders)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var buf bytes.Buffer

	err = createInvoice(&buf, &id, &invoice, &finance, orders)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	//w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	buf.WriteTo(w)
	return
}

func isNullString(v models.NullString, prefix string, sufix string) string {
	s := string(v)

	if s == "" {
		return ""
	}

	return fmt.Sprintf("%s%s%s", prefix, string(s), sufix)
}

func format_long_date(d string) string {
	const (
		RFC3339  = "2006-01-02"
		layoutID = "02 Jan 2006"
	)

	t, _ := time.Parse(RFC3339, d[0:10])
	return t.Format(layoutID)
}

func format_short_date(d string) string {
	const (
		RFC3339  = "2006-01-02"
		layoutID = "02-01-2006"
	)

	t, _ := time.Parse(RFC3339, d[0:10])
	return t.Format(layoutID)
}

func TermString(v int) string {
	if v == 1 {
		return "Cash"
	}
	return "Transfer"
}

func format_number(f float64) string {
	p := message.NewPrinter(language.Indonesian)
	s := p.Sprintf("%0.f", f)
	return s
}

func createInvoice(w io.Writer, invoice_id *int64, inv *invoice_item, finance *models.Finance, orders []models.Order) (err error) {

	const (
		unit        = "mm"
		size        = "A4"
		orientation = "P"
		font        = "Helvetica"
		col1        = "No. SPK"
		col2        = "TANGGAL"
		col3        = "MERK"
		col4        = "TYPE"
		col5        = "NOPOL"
		col6        = "TAHUN"
		col7        = "BT FINANCE"
		col8        = "PAJAK"

		cw1 = 35
		cw2 = 18
		cw3 = 20
		cw4 = 20
		cw5 = 20
		cw6 = 12
		cw7 = 25
		cw8 = 15
	)
	var lh float64 = 5.5
	var mt float64 = 25
	var ml float64 = 30
	var mr float64 = 15

	pdf := gofpdf.New(orientation, unit, size, "")
	pdf.SetMargins(ml, mt, mr)
	pw, _ := pdf.GetPageSize()
	pw = pw - ml - mr

	pdf.AddPage()
	pdf.SetFont(font, "B", 16)
	pdf.CellFormat(pw, 15, fmt.Sprintf("INVOICE #%d", (*invoice_id)), "", 1, "L", false, 0, "")

	pdf.SetFont(font, "B", 10)
	pdf.CellFormat(pw, lh, finance.Name+" - "+finance.ShortName, "", 1, "L", false, 0, "")

	pdf.SetFont(font, "", 10)

	address := fmt.Sprintf("%s%s%s\n%s%s\n%s",
		finance.Street,
		isNullString(finance.City, ", ", " "),
		finance.Zip,
		isNullString(finance.Phone, "Telp. ", ""),
		isNullString(finance.Cell, " / ", ""),
		isNullString(finance.Email, "e-mail: ", ""),
	)

	pdf.MultiCell(pw/2, lh, address, "", "L", false)

	x := (pw+ml)/2 + 25
	w2 := pw - x
	y := mt + 15
	koma := w2/2 + 5

	pdf.SetXY(x, y)
	pdf.Cell(koma, lh, "Tanggal:")
	pdf.Cell(w2/2, lh, format_long_date(inv.InvoiceAt))

	y += lh
	pdf.SetXY(x, y)
	pdf.Cell(koma, lh, "Salesman:")
	pdf.Cell(w2/2, lh, inv.Salesman)

	y += lh
	pdf.SetXY(x, y)
	pdf.Cell(koma, lh, "Payment term:")
	pdf.Cell(w2/2, lh, TermString(int(inv.PaymentTerm)))

	y += lh
	pdf.SetXY(x, y)
	pdf.Cell(koma, lh, "Jatuh tempo:")
	pdf.Cell(w2/2, lh, format_long_date(inv.DueAt))

	y += lh + 5
	x = ml
	pdf.SetXY(x, y)

	pdf.SetFont(font, "B", 10)
	pdf.CellFormat(30, lh, "Invoice details", "", 0, "L", false, 0, "")

	y += lh + 2
	pdf.SetXY(x, y)

	r, g, b := pdf.GetFillColor()

	pdf.SetFillColor(225, 225, 225)
	pdf.SetLineWidth(0.1)

	pdf.SetFont(font, "", 8)
	pdf.CellFormat(cw1, lh, col1, "1", 0, "L", true, 0, "")
	pdf.CellFormat(cw2, lh, col2, "1", 0, "C", true, 0, "")
	pdf.CellFormat(cw3, lh, col3, "1", 0, "L", true, 0, "")
	pdf.CellFormat(cw4, lh, col4, "1", 0, "L", true, 0, "")
	pdf.CellFormat(cw5, lh, col5, "1", 0, "L", true, 0, "")
	pdf.CellFormat(cw6, lh, col6, "1", 0, "C", true, 0, "")
	pdf.CellFormat(cw7, lh, col7, "1", 0, "R", true, 0, "")
	pdf.CellFormat(cw8, lh, col8, "1", 1, "R", true, 0, "")

	// y += lh + 2
	pdf.SetX(x)
	pdf.SetY(pdf.GetY() + 0.4)

	pdf.SetFillColor(r, g, b)

	for i := 0; i < len(orders); i++ {
		o := orders[i]
		pdf.CellFormat(cw1, lh, o.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(cw2, lh, format_short_date(o.OrderAt), "1", 0, "C", false, 0, "")
		pdf.CellFormat(cw3, lh, o.Unit.Type.Merk.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(cw4, lh, o.Unit.Type.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(cw5, lh, o.Unit.Nopol, "1", 0, "L", false, 0, "")
		pdf.CellFormat(cw6, lh, fmt.Sprintf("%d", o.Unit.Year), "1", 0, "C", false, 0, "")
		pdf.CellFormat(cw7, lh, format_number(o.BtFinance), "1", 0, "R", false, 0, "")
		pdf.CellFormat(cw8, lh, format_number(o.Nominal), "1", 1, "R", false, 0, "")
	}

	pdf.SetY(pdf.GetY() + 0.4)

	x = ml + cw1 + cw2 + cw3 + cw4
	pdf.SetX(x)
	pdf.SetFont(font, "", 8)
	pdf.CellFormat(cw5+cw6, lh, "Total invoice:", "", 0, "L", false, 0, "")
	pdf.SetFont(font, "B", 8)
	pdf.CellFormat(cw7+cw8, lh, format_number(inv.Total), "1", 1, "R", false, 0, "")
	pdf.SetX(x)
	pdf.SetFont(font, "", 8)
	pdf.CellFormat(cw5+cw6, lh, "Pajak:", "", 0, "L", false, 0, "")
	pdf.SetFont(font, "B", 8)
	pdf.CellFormat(cw7+cw8, lh, format_number(inv.Tax), "1", 1, "R", false, 0, "")
	pdf.SetX(x)
	pdf.SetFont(font, "", 8)
	pdf.CellFormat(cw5+cw6, lh, "Grand Total:", "", 0, "L", false, 0, "")
	pdf.SetFont(font, "B", 8)
	pdf.CellFormat(cw7+cw8, lh, format_number(inv.Tax+inv.Total), "1", 1, "R", false, 0, "")

	pdf.SetFont(font, "", 10)

	pw = (cw1 + cw2 + cw3 + cw4 + cw5 + cw6 + cw7 + cw8) / 2
	pdf.SetY(pdf.GetY() + 10)
	pdf.CellFormat(pw*2, lh, "Mengetahui:", "0", 1, "C", false, 0, "")
	pdf.SetY(pdf.GetY() + 15)
	pdf.SetFont(font, "B", 10)
	pdf.CellFormat(pw-20, lh, "Finance,", "0", 0, "C", false, 0, "")
	pdf.SetX(pdf.GetX() + 40)
	pdf.CellFormat(pw-20, lh, "SPRS,", "0", 1, "C", false, 0, "")

	err = pdf.Output(w)
	pdf.Close()
	//_ = pdf.OutputFileAndClose("hello.pdf")
	return err
}
