package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "image/jpeg"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"fyc.com/sprs/models"
	"github.com/MasturFYC/fyc"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
)

type order_unit_customer struct {
	models.Order
	Unit     *models.Unit     `json:"unit,omitempty"`
	Customer *models.Customer `json:"customer,omitempty"`
}

func Clipan_GetInvoice(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	invoice, err := invoice_get_item_customer(&id)

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

	var orders []order_unit_customer
	source = (*json.RawMessage)(&invoice.Details)
	err = json.Unmarshal(*source, &orders)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var buf bytes.Buffer

	err = clipan_create_invoice(&buf, &id, &invoice, &finance, orders)

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

func clipan_create_invoice(w io.Writer, invoice_id *int64, inv *invoice_item, finance *models.Finance, orders []order_unit_customer) (err error) {

	var account models.AccCode
	source := (*json.RawMessage)(&inv.Account)
	_ = json.Unmarshal(*source, &account)

	const (
		unit        = "mm"
		size        = "A4"
		orientation = "P"
		font        = "Helvetica"
		font2       = "Times"
		col1        = "NO"
		col2        = "KETERANGAN"
		col3        = "JUMLAH"
	)
	var lh float64 = 5.5
	var mt float64 = 25
	var ml float64 = 20
	var mr float64 = 15

	p := gofpdf.New(orientation, unit, size, "")
	p.SetMargins(ml, mt, mr)
	pw, _ := p.GetPageSize()

	box := pw - ml - mr
	box1 := 30.0
	box2 := box - box1
	x := ml + box1

	//img, _ := OpenImage(filepath.Join(os.Getenv("POSTGRES_URL"), "logo.jpg"))

	p.AddPage()

	p.SetY(10)

	//log.Println()
	p.Image(filepath.Join(os.Getenv("UPLOADFILE_LOCATION"), "logo.jpg"), ml, p.GetY(), box1, 0, false, "", 0, "")
	p.SetX(x)
	p.SetFont(font, "B", 18)
	p.CellFormat(box2, lh, "PT. SARANA PADMA RIDHO SEPUH", "", 1, "C", false, 0, "")
	p.SetFont(font2, "", 12)
	p.SetY(p.GetY() + 2)
	p.SetX(x)
	p.CellFormat(box2, lh, "GENERAL SUPPLIER, CONTRACTOR, COLLECTION", "", 1, "C", false, 0, "")
	p.SetX(x)
	p.CellFormat(box2, lh, "Jl. Gator Subroto Villa Gatsu No. 01 - Indramayu", "", 1, "C", false, 0, "")

	p.SetX(ml)
	p.SetLineWidth(0.75)
	p.Line(ml, p.GetY()+2, ml+box, p.GetY()+2)
	p.SetLineWidth(0.25)
	p.Line(ml, p.GetY()+3, ml+box, p.GetY()+3)

	p.SetY(p.GetY() + 10)
	p.SetFont(font2, "B", 16)
	p.CellFormat(box, lh, "INVOICE", "", 1, "C", false, 0, "")

	y := p.GetY() + 5
	p.SetY(y)
	p.SetFont(font2, "", 12)
	p.CellFormat(box, lh, "Kepada Yth", "", 1, "L", false, 0, "")
	p.CellFormat(box, lh, "Pimpinan", "", 1, "L", false, 0, "")
	p.CellFormat(box, lh, "PT. CLIPAN FINANCE INDONESIA", "", 1, "L", false, 0, "")

	lh = 7.5
	box1 = 25.0
	box2 = 50.0

	x = pw - box1 - box2 - mr

	p.SetLineWidth(0.15)
	p.SetXY(x, y)
	p.CellFormat(box1, lh, " No. Invoice", "1", 0, "LM", false, 0, "")
	p.CellFormat(box2, lh, fyc.CreateInvoiceNumber(inv.ID, inv.InvoiceAt), "1", 1, "L", false, 0, "")

	p.SetX(x)
	p.CellFormat(box1, lh, " Tanggal", "1", 0, "LM", false, 0, "")
	p.CellFormat(box2, lh, fyc.CreateIndonesianDate(inv.InvoiceAt, false), "1", 1, "L", false, 0, "")

	p.SetY(p.GetY() + 5)
	p.CellFormat(box, lh, "Dengan hormat,", "0", 1, "L", false, 0, "")
	lh = 5.5
	p.MultiCell(box, lh, "Bersama dengan ini kami mengajukan penagihan komisi atas pencapaian hasil kerja\nAdapun rincian sebagai berikut:",
		"", "L", false)

	p.SetY(p.GetY() + 2)
	lh = 7
	box1 = box * (10.0 / 100.0)
	box2 = box * (65.0 / 100.0)
	box3 := box * (25.0 / 100.0)
	p.SetFont(font2, "B", 12)
	p.CellFormat(box1, lh, col1, "1", 0, "CM", false, 0, "")
	p.CellFormat(box2, lh, col2, "1", 0, "CM", false, 0, "")
	p.CellFormat(box3, lh, col3, "1", 1, "CM", false, 0, "")

	//y1 := p.GetY()
	p.SetY(p.GetY() + 0.5)
	p.Line(ml, p.GetY(), ml+box1+box2+box3, p.GetY())
	total := 0.0
	p.SetFont(font2, "", 12)
	for i := 0; i < len(orders); i++ {
		o := orders[i]
		y = p.GetY()
		top := y + 1
		p.SetXY(ml+box1+2, top)

		if o.Customer != nil {
			p.MultiCell(box2-4, 5, fmt.Sprintf("Success fee atas nama %s\r\nNomor Perjanjian pembiayaan %s\r\nNomor Polisi %s",
				o.Customer.Name, o.Name, o.Unit.Nopol), "0", "LM", false)
		} else {
			p.MultiCell(box2-4, 5, fmt.Sprintf("Success fee atas nama %s\r\nNomor Perjanjian pembiayaan %s\r\nNomor Polisi %s",
				finance.Name, o.Name, o.Unit.Nopol), "0", "LM", false)
		}
		bottom := p.GetY() + 1

		p.Line(ml, bottom, ml+box1+box2+box3, bottom)
		p.Line(ml, y, ml, bottom)
		p.Line(ml+box1, y, ml+box1, bottom)
		p.Line(ml+box1+box2, y, ml+box1+box2, bottom)
		p.Line(ml+box1+box2+box3, y, ml+box1+box2+box3, bottom)

		lh = bottom - top + 1
		top = top - 1
		p.SetY(top)
		p.SetX(ml)
		p.CellFormat(box1, lh, fmt.Sprintf("%d.", i+1), "0", 1, "CM", false, 0, "")
		p.SetXY(ml+box1+box2, top)
		subtotal := o.BtFinance - (o.BtFinance * (float64(inv.Ppn) / 100.0))
		total = total + subtotal
		p.CellFormat(box3, lh, fmt.Sprintf("Rp %s   ", format_number(subtotal)), "0", 1, "RM", false, 0, "")
	}

	p.SetY(p.GetY() + 0.5)
	lh = 7.5
	p.CellFormat(box1+box2, lh, "  Total Komisi", "1", 0, "LM", false, 0, "")
	p.SetFont(font2, "B", 12)
	p.CellFormat(box3, lh, fmt.Sprintf("Rp %s   ", format_number(total)), "1", 1, "RM", false, 0, "")

	lh = 5.5
	p.SetFont(font2, "", 12)
	p.SetY(p.GetY() + 3)
	sw := p.GetStringWidth("Terbilang : ")
	p.CellFormat(sw, lh, "Terbilang :", "0", 0, "L", false, 0, "")
	p.SetFont(font2, "I", 12)
	p.CellFormat(box-sw, lh, fmt.Sprintf("( ** %s Rupiah )", strings.TrimSpace(fyc.Terbilang(total))), "0", 1, "L", false, 0, "")
	p.SetFont(font2, "", 12)
	p.CellFormat(box, lh, fmt.Sprintf("Pembayaran ke Rekening %s %s", account.Name, account.Descriptions), "0", 1, "L", false, 0, "")

	p.SetY(p.GetY() + 3)
	p.MultiCell(box, 5, "Demikian pengajuan penagihan ini kami buat, atas perhatian dan kerjasamanya kami Ucapkan terima kasih.",
		"0", "L", false)

	p.SetY(p.GetY() + 5)
	p.CellFormat(box3, lh, "Hormat kami", "0", 1, "L", false, 0, "")
	p.CellFormat(box3, lh, "PT SARANA PADMA RIDHO SEPUH", "0", 1, "L", false, 0, "")
	p.SetY(p.GetY() + 15)
	p.SetFont(font2, "BU", 12)
	p.CellFormat(box3, lh, "DEDDY PRANOTO", "0", 1, "L", false, 0, "")
	p.SetFont(font2, "", 12)
	p.CellFormat(box3, lh, "Direktur", "0", 0, "L", false, 0, "")

	err = p.Output(w)
	p.Close()
	//_ = pdf.OutputFileAndClose("hello.pdf")
	return err
}
