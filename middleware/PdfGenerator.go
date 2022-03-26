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

type order_unit struct {
	models.Order
	Unit models.Unit `json:"unit,omitempty"`
}

// //will call the init() function of the package
// //thus enabling working with jpeg file
// func OpenImage(path string) (image.Image, error) {
// 	f, err := os.Open(path)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	defer f.Close()
// 	img, format, err := image.Decode(f)
// 	if err != nil {
// 		e := fmt.Errorf("error in decoding: %w", err)
// 		return nil, e
// 	}

// 	if format != "jpeg" && format != "png" {
// 		e := fmt.Errorf("error in image format - not jpeg")
// 		return nil, e
// 	}
// 	return img, nil
// }

func Pdf_GetInvoice(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

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

	var orders []order_unit
	source = (*json.RawMessage)(&invoice.Details)
	err = json.Unmarshal(*source, &orders)

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
	//return
}

func Pdf_GetInvoiceClipan(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

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

	var orders []order_unit
	source = (*json.RawMessage)(&invoice.Details)
	err = json.Unmarshal(*source, &orders)

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
	//return
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

func createInvoice(w io.Writer, invoice_id *int64, inv *invoice_item, finance *models.Finance, orders []order_unit) (err error) {

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
		//		col8        = "PAJAK"

		cw1 = 35
		cw2 = 18
		cw3 = 20
		cw4 = 20
		cw5 = 20
		cw6 = 17
		cw7 = 35
		//		cw8 = 15
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

	pdf.SetY(10)
	pdf.SetFont(font, "B", 16)
	pdf.CellFormat(pw, lh, "PT. SARANA PADMA RIDHO SEPUH", "", 1, "C", false, 0, "")
	pdf.SetFont(font, "", 10)
	pdf.SetY(pdf.GetY() + 2)
	pdf.CellFormat(pw, lh, "GENERAL SUPPLIER, CONTRACTOR, COLLECTION", "", 1, "C", false, 0, "")
	pdf.CellFormat(pw, lh, "Jl. Gator Subroto Villa Gatsu No. 01 - Indramayu", "", 1, "C", false, 0, "")

	pdf.SetLineWidth(0.75)
	pdf.Line(ml, pdf.GetY()+2, ml+pw, pdf.GetY()+2)
	pdf.SetLineWidth(0.25)
	pdf.Line(ml, pdf.GetY()+3, ml+pw, pdf.GetY()+3)

	lastY := pdf.GetY() + 5

	x := (pw+ml)/2 + 40
	w2 := pw - x
	y := lastY - 2
	koma := w2/2 + 12

	pdf.SetXY(x, y)
	pdf.SetFont(font, "B", 16)
	pdf.CellFormat(pw, 15, fmt.Sprintf("INVOICE #%d", (*invoice_id)), "", 1, "L", false, 0, "")

	y += 12
	pdf.SetFont(font, "", 10)
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

	pdf.SetY(lastY + 2)
	pdf.CellFormat(pw, lh, "Customer / Mitra kerja:", "", 1, "L", false, 0, "")
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

	if pdf.GetY() < lastY {
		y = lastY + 5
	} else {
		y = pdf.GetY() + 5
	}
	x = ml
	pdf.SetXY(x, y)

	pdf.SetFont(font, "B", 10)
	pdf.CellFormat(pw, lh, "Invoice details", "", 1, "L", false, 0, "")

	y = pdf.GetY() + 2
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
	pdf.CellFormat(cw7, lh, col7, "1", 1, "R", true, 0, "")
	//pdf.CellFormat(cw8, lh, col8, "1", 1, "R", true, 0, "")

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
		pdf.CellFormat(cw7, lh, format_number(o.BtFinance), "1", 1, "R", false, 0, "")
		//		pdf.CellFormat(cw8, lh, format_number(o.Nominal), "1", 1, "R", false, 0, "")
	}

	pdf.SetY(pdf.GetY() + 0.4)

	x = ml + cw1 + cw2 + cw3 + cw4
	pdf.SetX(x)
	pdf.CellFormat(cw5+cw6, lh, "Subtotal:", "", 0, "L", false, 0, "")
	pdf.CellFormat(cw7, lh, format_number(inv.Subtotal), "1", 1, "R", false, 0, "")
	pdf.SetX(x)
	pdf.CellFormat(cw5+cw6, lh, fmt.Sprintf("PPN: %s%%", format_number(float64(inv.Ppn))), "", 0, "L", false, 0, "")
	pdf.CellFormat(cw7, lh, format_number(float64(inv.Tax)), "1", 1, "R", false, 0, "")
	pdf.SetX(x)
	pdf.CellFormat(cw5+cw6, lh, "Total invoice:", "", 0, "L", false, 0, "")
	pdf.SetFont(font, "B", 8)
	pdf.CellFormat(cw7, lh, format_number(inv.Total), "1", 1, "R", false, 0, "")

	pdf.SetFont(font, "", 10)

	pw = (cw1 + cw2 + cw3 + cw4 + cw5 + cw6 + cw7) / 2
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

func create_indonesian_date(date string) string {
	months := [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "Nopember", "Desember"}

	t, err := time.Parse("2006-01-02", date[0:10])

	if err != nil {
		return date[0:10]
	}
	year, month, day := t.Date()
	return fmt.Sprintf(" %02d %s %d", day, months[month-1], year)
}

func create_invoice_number(id int64, date string) string {
	rom := [12]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII"}

	t, err := time.Parse("2006-01-02", date[0:10])

	if err != nil {
		return date[0:10]
	}
	year, month, _ := t.Date()
	return fmt.Sprintf(" %d/INV/SPRS/%s/%d", id, rom[month-1], year)
	//return fmt.Sprintf(" %05d/INV/SPRS/%s/%d", id, rom[month-1], year)
}
