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

func Mtf_GetInvoice(w http.ResponseWriter, r *http.Request) {
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

	var orders []order_unit
	source = (*json.RawMessage)(&invoice.Details)
	err = json.Unmarshal(*source, &orders)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var buf bytes.Buffer

	err = mtf_create_invoice(&buf, &id, &invoice, &finance, orders)

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

func mtf_create_invoice(w io.Writer, invoice_id *int64, inv *invoice_item, finance *models.Finance, orders []order_unit) (err error) {

	var account models.AccCode
	source := (*json.RawMessage)(&inv.Account)
	_ = json.Unmarshal(*source, &account)

	const (
		unit        = "mm"
		size        = "A4"
		orientation = "P"
		font        = "Arial"
		//font       = "Times"
		col1 = "ITEM"
		col2 = "DESCRIPTION"
		col3 = "QTY"
		col4 = "UNIT"
		col5 = "UNIT PRICE\n( Rp )"
		col6 = "TOTAL PRICE\n( Rp )"
	)
	var lh float64 = 5.5
	var mt float64 = 25
	var ml float64 = 20
	var mr float64 = 15

	p := gofpdf.New(orientation, unit, size, "")
	p.SetAutoPageBreak(true, mr)
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
	p.SetFont(font, "", 10)
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
	p.SetFont(font, "B", 16)
	box1 = 50
	p.SetX(ml + (box-box1)/2)
	p.CellFormat(box1, lh+3, "INVOICE", "1", 1, "CM", false, 0, "")

	y := p.GetY() + 7
	p.SetY(y)
	p.SetFont(font, "B", 10)
	p.CellFormat(box, lh, "Kepada Yth", "", 1, "L", false, 0, "")

	p.SetFont(font, "", 10)
	box1 = p.GetStringWidth("Name : ") + 3
	p.CellFormat(box1, lh, "Name : ", "", 0, "L", false, 0, "")
	p.SetFont(font, "B", 10)
	p.CellFormat(box-box1, lh, finance.Name, "", 1, "L", false, 0, "")

	lh = 5.5
	box1 = 25.0
	box2 = 50.0
	box3 := 3.0

	x = pw - box1 - box2 - mr - box3

	p.SetFont(font, "", 10)
	p.SetXY(x, y)
	p.CellFormat(box1, lh, "Invoice No.", "", 0, "R", false, 0, "")
	p.CellFormat(box3, lh, ":", "", 0, "R", false, 0, "")
	p.CellFormat(box2, lh, fyc.CreateInvoiceNumber(inv.ID, inv.InvoiceAt), "", 1, "L", false, 0, "")

	p.SetX(x)
	p.CellFormat(box1, lh, "Date", "0", 0, "R", false, 0, "")
	p.CellFormat(box3, lh, ":", "", 0, "R", false, 0, "")
	p.CellFormat(box2, lh, fyc.CreateIndonesianDate(inv.InvoiceAt, false), "0", 1, "L", false, 0, "")

	box1 = 30
	box2 = box - box1 - box3
	p.SetY(p.GetY() + 5)
	p.CellFormat(box1, lh, "Nama debitur", "", 0, "L", false, 0, "")
	p.CellFormat(box3, lh, ":", "", 0, "R", false, 0, "")
	p.SetTextColor(255, 0, 0)
	p.CellFormat(box2, lh, finance.Name, "", 1, "L", false, 0, "")
	p.SetTextColor(0, 0, 0)
	p.CellFormat(box1, lh, "No. Kontrak", "", 0, "L", false, 0, "")
	p.CellFormat(box3, lh, ":", "", 0, "R", false, 0, "")
	p.SetTextColor(255, 0, 0)
	p.CellFormat(box2, lh, inv.Salesman, "", 1, "L", false, 0, "")
	p.SetTextColor(0, 0, 0)

	p.SetY(p.GetY() + 2)
	lh = 5
	box1 = box * (8.0 / 100.0)
	box2 = box * (40.0 / 100.0)
	box3 = box * (8.0 / 100.0)
	box4 := box * (8.0 / 100.0)
	box5 := box * (16.0 / 100.0)
	box6 := box * (20.0 / 100.0)

	y = p.GetY()
	x = ml + box1 + box2 + box3 + box4

	p.SetFont(font, "B", 10)
	p.SetLineWidth(0.1)
	p.SetXY(x, y)
	p.MultiCell(box5, lh, col5, "1", "CM", false)

	x = x + box5

	p.SetXY(x, y)
	p.MultiCell(box6, lh, col6, "1", "CM", false)

	b := p.GetY() - y

	p.SetY(y)
	p.CellFormat(box1, b, col1, "1", 0, "CM", false, 0, "")
	p.CellFormat(box2, b, col2, "1", 0, "CM", false, 0, "")
	p.CellFormat(box3, b, col3, "1", 0, "CM", false, 0, "")
	p.CellFormat(box4, b, col4, "1", 1, "CM", false, 0, "")

	//y1 := p.GetY()
	p.SetY(p.GetY() + 0.5)

	total := 0.0
	total_pajak := 0.0
	total_finance := 0.0

	p.SetFont(font, "", 10)
	for i := 0; i < len(orders); i++ {
		o := orders[i]
		y = p.GetY()
		top := y + 1

		p.SetXY(ml+box1+1, top+1)
		p.MultiCell(box2-2, 5, fmt.Sprintf("Biaya REPPO FEE %s\nTahun %d\nPPH 23 ( 2%% )",
			o.Unit.Type.Name, o.Unit.Year), "0", "LM", false)

		subtotal := o.BtFinance
		pajak := o.BtFinance * (float64(inv.Ppn) / 100.0)

		p.SetXY(ml+box1+box2+box3+box4+box5, top+1)
		p.MultiCell(box6-1, 5, fmt.Sprintf("%s\n-%s\n%s",
			format_number(subtotal),
			format_number(pajak),
			format_number(subtotal-pajak),
		), "0", "RM", false)

		bottom := p.GetY() + 2

		p.Line(ml, top-1, pw-mr, top-1)
		p.Line(ml, y, ml, bottom)
		p.Line(ml+box1, y, ml+box1, bottom)
		p.Line(ml+box1+box2, y, ml+box1+box2, bottom)
		p.Line(ml+box1+box2+box3, y, ml+box1+box2+box3, bottom)
		p.Line(ml+box1+box2+box3+box4, y, ml+box1+box2+box3+box4, bottom)
		p.Line(ml+box1+box2+box3+box4+box5, y, ml+box1+box2+box3+box4+box5, bottom)
		p.Line(ml+box1+box2+box3+box4+box5+box6, y, ml+box1+box2+box3+box4+box5+box6, bottom)

		lh = bottom - top + 2

		p.SetX(ml)
		p.SetY(top - 1)
		p.CellFormat(box1, lh, fmt.Sprintf("%d.", i+1), "0", 0, "CM", false, 0, "")
		p.SetX(ml + box1 + box2)
		p.CellFormat(box3, lh, fmt.Sprintf("%d", 1), "0", 0, "CM", false, 0, "")
		p.SetX(ml + box1 + box2 + box3)
		p.CellFormat(box4, lh, o.Unit.Type.Wheel.ShortName, "0", 0, "CM", false, 0, "")
		p.SetX(ml + box1 + box2 + box3 + box4)
		p.CellFormat(box5-1, lh, format_number(subtotal), "0", 1, "RM", false, 0, "")

		p.SetY(p.GetY() - 1)

		total = total + (subtotal - pajak)
		total_finance = total_finance + subtotal
		total_pajak = total_pajak + pajak
	}

	p.Line(ml, p.GetY(), pw-mr, p.GetY())

	p.SetY(p.GetY() + 0.5)
	lh = 7.5

	p.CellFormat(box1+box2+box3+box4+box5, lh, "  Subtotal", "1", 0, "LM", false, 0, "")
	//p.CellFormat(box5, lh, " ", "1", 0, "LM", false, 0, "")
	p.SetFont(font, "B", 10)
	p.CellFormat(box6, lh, fmt.Sprintf("%s ", format_number(total_finance)), "1", 1, "RM", false, 0, "")

	p.SetFont(font, "", 10)
	p.CellFormat(box1+box2+box3+box4+box5, lh, "  Pajak", "1", 0, "LM", false, 0, "")
	//p.CellFormat(box5, lh, " ", "1", 0, "LM", false, 0, "")
	p.SetFont(font, "B", 10)
	p.CellFormat(box6, lh, fmt.Sprintf("-%s ", format_number(total_pajak)), "1", 1, "RM", false, 0, "")

	p.SetFont(font, "", 10)
	p.CellFormat(box1+box2+box3+box4+box5, lh, "  Total", "1", 0, "LM", false, 0, "")
	//p.CellFormat(box5, lh, " ", "1", 0, "LM", false, 0, "")
	p.SetFont(font, "B", 10)
	p.CellFormat(box6, lh, fmt.Sprintf("%s ", format_number(total)), "1", 1, "RM", false, 0, "")

	lh = 5.5
	x = pw - mr - box5 - box6
	y = p.GetY() + 10
	box1 = box5 + box6

	p.SetXY(x, y)
	p.CellFormat(box1, lh, "PT SARANA PADMA RIDHO SEPUH", "0", 1, "C", false, 0, "")
	p.SetXY(x, p.GetY()+15)
	p.SetFont(font, "BU", 10)
	p.CellFormat(box1, lh, "DEDDY PRANOTO", "0", 1, "C", false, 0, "")
	p.SetX(x)
	p.SetFont(font, "", 10)
	p.CellFormat(box1, lh, "Direktur", "0", 0, "C", false, 0, "")

	x = ml
	y = p.GetY() - (lh * 2)
	box1 = box1 + box2 + box3

	p.SetXY(x, y)
	p.CellFormat(box1, lh, "Pembayaran dengan transfer dialamatkan ke:", "0", 1, "L", false, 0, "")
	p.CellFormat(box1, lh, account.Name, "0", 1, "L", false, 0, "")
	p.CellFormat(box1, lh, string(account.Descriptions), "0", 1, "L", false, 0, "")

	/*

	   BANK MANDIRI
	   Rekening No. 134-00-0000610-5
	   a/n. PT SARANA PADMA RIDHO SEPUH
	*/
	mft_create_lampiran1(font, p, inv, finance, &total)
	mft_create_lampiran2(font, p, inv, finance, &total)

	err = p.Output(w)
	p.Close()
	//_ = pdf.OutputFileAndClose("hello.pdf")
	return err
}

func mft_create_lampiran1(
	font string,
	p *gofpdf.Fpdf,
	inv *invoice_item,
	finance *models.Finance,
	total *float64,
) {

	ml, _, mr, _ := p.GetMargins()

	lh := 5.5

	pw, _ := p.GetPageSize()
	box := pw - ml - mr
	box1 := 30.0
	box2 := box - box1

	p.AddPage()
	p.SetY(10)
	//log.Println()
	p.SetX(ml)
	p.Image(filepath.Join(os.Getenv("UPLOADFILE_LOCATION"), "logo.jpg"), ml, p.GetY(), box1, 0, false, "", 0, "")

	x := ml + box1

	p.SetX(x)
	p.SetFont(font, "B", 18)
	p.CellFormat(box2, lh, "PT. SARANA PADMA RIDHO SEPUH", "", 1, "C", false, 0, "")
	p.SetFont(font, "", 10)
	p.SetY(p.GetY() + 2)
	p.SetX(x)
	p.CellFormat(box2, lh, "GENERAL SUPPLIER, CONTRACTOR, COLLECTION", "", 1, "C", false, 0, "")
	p.SetX(x)
	p.CellFormat(box2, lh, "Jl. Gator Subroto Villa Gatsu No. 01 - Indramayu", "", 1, "C", false, 0, "")

	x = ml

	p.SetX(x)
	p.SetLineWidth(0.75)
	p.Line(ml, p.GetY()+2, ml+box, p.GetY()+2)
	p.SetLineWidth(0.25)
	p.Line(ml, p.GetY()+3, ml+box, p.GetY()+3)

	p.SetY(p.GetY() + 10)
	p.SetFont(font, "BU", 14)
	p.CellFormat(box, lh, "SURAT PERNYATAAN", "0", 1, "CM", false, 0, "")

	p.SetY(p.GetY() + 10)

	lh = 7

	p.SetFont(font, "", 10)
	p.CellFormat(box, lh, fmt.Sprintf("Indramayu, %s", fyc.CreateIndonesianDate(inv.InvoiceAt, false)),
		"", 1, "L", false, 0, "")
	p.CellFormat(box, lh, "Bersama ini, saya yang bertandatangan di bawah ini:",
		"", 1, "L", false, 0, "")

	box1 = 40
	box2 = 3
	box3 := box - box1 - box2

	lh = 5.5
	p.SetY(p.GetY() + 1)

	p.CellFormat(box1, lh, "Nama", "", 0, "L", false, 0, "")
	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	p.SetFont(font, "B", 10)
	p.CellFormat(box3, lh, "DEDDY PRANOTO", "", 1, "L", false, 0, "")
	p.SetFont(font, "", 10)
	p.CellFormat(box1, lh, "No. KTP", "", 0, "L", false, 0, "")
	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	p.SetFont(font, "B", 10)
	p.CellFormat(box3, lh, "3329130904840004", "", 1, "L", false, 0, "")
	p.SetFont(font, "", 10)
	p.CellFormat(box1, lh, "Pekerjaan", "", 0, "L", false, 0, "")
	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	p.SetFont(font, "B", 10)
	p.CellFormat(box3, lh, "DIREKTUR PT SARANA PADMA RIDHO SEPUH", "", 1, "L", false, 0, "")

	p.SetY(p.GetY() + 3)

	lh = 7

	p.SetFont(font, "", 10)
	p.MultiCell(box, lh, "Telah menyelesaikan tugas penarikan unit PT Mandiri Tunas Finance dengan data-data sbb :", "", "L", false)

	lh = 5.5
	p.CellFormat(box1, lh, "No. Kontrak", "", 0, "L", false, 0, "")
	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	p.SetFont(font, "B", 10)
	p.SetTextColor(255, 0, 0)
	p.CellFormat(box3, lh, "-----------", "", 1, "L", false, 0, "")
	p.SetFont(font, "", 10)

	p.SetTextColor(0, 0, 0)
	p.CellFormat(box1, lh, "Nama Customer", "", 0, "L", false, 0, "")
	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	p.SetFont(font, "B", 10)
	p.SetTextColor(255, 0, 0)
	p.CellFormat(box3, lh, fmt.Sprintf("%s %s", finance.Name, finance.ShortName), "", 1, "L", false, 0, "")

	p.SetFont(font, "", 10)

	p.SetTextColor(0, 0, 0)
	p.CellFormat(box1, lh, "Type/No.Pol", "", 0, "L", false, 0, "")
	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	p.SetFont(font, "B", 10)
	p.SetTextColor(255, 0, 0)
	p.CellFormat(box3, lh, "E - 5565 - KKB", "", 1, "L", false, 0, "")
	p.SetFont(font, "", 10)

	p.SetTextColor(0, 0, 0)
	p.SetY(p.GetY() + 1)
	html := p.HTMLBasicNew()
	_, lineHt := p.GetFontSize()
	lineHt += 2
	html.Write(lineHt, fmt.Sprintf("Dengan total biaya <b>Rp%s,-</b> (<i>** %s Rupiah</i>)",
		format_number(*total), strings.TrimSpace(fyc.Terbilang(*total))))

	html.Write(lineHt,
		fmt.Sprintf(`<br /><br />Bersama ini saya menyatakan dengan sebenarnya bahwa unit tersebut sudah bebas
dari kasus manapun dari masalah serta dapat segera dilakukan lelang, penggunaan
dari seluruh biaya tarik tersebut ( Rp%s,- ) adalah menjadi tanggung jawab saya.
Apabila ada tuntutan di kemudian hari ataupun masalah yang terjadi pada penggunaan
biaya maupun unit tersebut di atas, akan menjadi tanggung jawab saya sepenuhnya dan
tidak melibatkan pihak %s sesuai dengan perjanjian yang telah disepakati bersama,
serta %s terbebas dari segala tuntutan maupun permasalahan yang terjadi di kemudian hari.`,
			format_number(*total), finance.Name, finance.Name))
	html.Write(lineHt, `Demikian surat pernyataan ini saya buat. Tanpa ada paksaan dari pihak manapun.`)

	p.SetY(p.GetY() + 15)

	box1 = box * (45.0 / 100.0)
	box2 = box * (10.0 / 100.0)
	p.CellFormat(box1, lh, "Hormat saya,", "", 0, "C", false, 0, "")
	p.CellFormat(box2, lh, "", "", 0, "L", false, 0, "")
	p.CellFormat(box1, lh, "Mengetahui,", "", 1, "C", false, 0, "")

	p.SetY(p.GetY() + 15)
	p.SetFont(font, "BU", 10)
	p.CellFormat(box1, lh, "Nama : DEDDY PRANOTO,", "", 0, "C", false, 0, "")
	p.CellFormat(box2, lh, "", "", 0, "L", false, 0, "")
	p.SetTextColor(255, 0, 0)
	p.CellFormat(box1, lh, "Nama :  ISKANDAR.G", "", 1, "C", false, 0, "")
	p.SetTextColor(0, 0, 0)
	p.SetFont(font, "", 10)
	p.CellFormat(box1, lh, "Direktur PT SARANA PADMA RIDHO SEPUH", "", 0, "C", false, 0, "")
	p.CellFormat(box2, lh, "", "", 0, "L", false, 0, "")
	p.CellFormat(box1, lh, "Collection Officer Macet", "", 1, "C", false, 0, "")

	p.SetY(p.GetY() + 10)
	p.SetFont(font, "", 8)
	p.CellFormat(box, lh, "NOTARIS MAISARAH PANE, SH", "", 1, "L", false, 0, "")
	p.CellFormat(box, lh, "KEMENKUMHAM.RI NO. AHU-0109242.AH.01.11.tahun 2021-08-24", "", 1, "L", false, 0, "")
	p.CellFormat(box, lh, "42.758.225.9-437.000", "", 1, "L", false, 0, "")

}

func mft_create_lampiran2(
	font string,
	p *gofpdf.Fpdf,
	inv *invoice_item,
	finance *models.Finance,
	total *float64,
) {

	ml, _, mr, _ := p.GetMargins()

	lh := 5.5

	pw, _ := p.GetPageSize()
	box := pw - ml - mr
	box1 := 30.0
	box2 := box - box1

	p.AddPage()
	p.SetY(10)
	//log.Println()
	p.SetX(ml)
	p.Image(filepath.Join(os.Getenv("UPLOADFILE_LOCATION"), "logo.jpg"), ml, p.GetY(), box1, 0, false, "", 0, "")

	x := ml + box1

	p.SetX(x)
	p.SetFont(font, "B", 18)
	p.CellFormat(box2, lh, "PT. SARANA PADMA RIDHO SEPUH", "", 1, "C", false, 0, "")
	p.SetFont(font, "", 10)
	p.SetY(p.GetY() + 2)
	p.SetX(x)
	p.CellFormat(box2, lh, "GENERAL SUPPLIER, CONTRACTOR, COLLECTION", "", 1, "C", false, 0, "")
	p.SetX(x)
	p.CellFormat(box2, lh, "Jl. Gator Subroto Villa Gatsu No. 01 - Indramayu", "", 1, "C", false, 0, "")

	x = ml

	p.SetX(x)
	p.SetLineWidth(0.75)
	p.Line(ml, p.GetY()+2, ml+box, p.GetY()+2)
	p.SetLineWidth(0.25)
	p.Line(ml, p.GetY()+3, ml+box, p.GetY()+3)

	p.SetY(p.GetY() + 10)
	p.SetFont(font, "Bu", 14)
	p.CellFormat(box, lh, "KRONOLOGIS DAN PENGAJUAN BIAYA", "0", 1, "CM", false, 0, "")

	p.SetY(p.GetY() + 10)

	lh = 5.5

	p.SetFont(font, "", 10)
	p.CellFormat(box, lh, "kepada Yth.", "", 1, "L", false, 0, "")
	p.SetFont(font, "B", 10)
	p.CellFormat(box, lh, "Pimpinan", "", 1, "L", false, 0, "")
	p.CellFormat(box, lh, finance.Name, "", 1, "L", false, 0, "")

	p.SetY(p.GetY() + 3)
	p.SetFont(font, "", 10)
	p.CellFormat(box, lh, "Di Tempat,", "", 1, "L", false, 0, "")

	p.SetY(p.GetY() + 3)

	box1 = 40
	box2 = 3
	box3 := box - box1 - box2

	lh = 5.5

	p.CellFormat(box1, lh, "Berikut kronologis penanganan pengamanan asset MTF", "", 1, "L", false, 0, "")
	p.SetY(p.GetY() + 1)

	p.CellFormat(box1, lh, "Nomor Kontrak", "", 0, "L", false, 0, "")
	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	p.SetFont(font, "B", 10)
	p.CellFormat(box3, lh, "---", "", 1, "L", false, 0, "")
	p.SetFont(font, "", 10)
	p.CellFormat(box1, lh, "Merk/ Type", "", 0, "L", false, 0, "")
	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	p.SetFont(font, "B", 10)
	p.CellFormat(box3, lh, "---", "", 1, "L", false, 0, "")
	p.SetFont(font, "", 10)
	p.CellFormat(box1, lh, "Nomor Polisi", "", 0, "L", false, 0, "")
	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	p.SetFont(font, "B", 10)
	p.CellFormat(box3, lh, "---", "", 1, "L", false, 0, "")

	p.SetY(p.GetY() + 3)

	lh = 5.5

	sb := strings.Builder{}
	sb.WriteString("Pencarian keberadaan unit selama ini tidak juga membuahkan hasil,")
	sb.WriteString(" yang pada akhirnya unit terlihat oleh tim excoll PT. SPRS Indramayu. Dicoba dilakukan")
	sb.WriteString(" penarikan kendaraan oleh tim PT.SPRS yang ternyata kondisi tidak kondusif karena unit")
	sb.WriteString(" ternyata dipegang oleh oknum preman yang cukup disegani di daerah Segeran Indramayu")
	sb.WriteString(" yang di mana PT. SPRS didatangi oleh massa yang cukup banyak karena kampung Segeran")
	sb.WriteString(" meminta bantuan ke kampung Tugu yang membuat keadaan jadi tidak kondusif dan sempat")
	sb.WriteString(" terjadi keributan antara massa pemegang unit dengan tim PT. SPRS. Untuk menjaga")
	sb.WriteString(" kondusifitas kantor PT. SPRS serta Kantor MTF Cirebon tim PT meminta bantuan Dalmas")
	sb.WriteString(" untuk memediasi antara massa pemegang unit dengan tim excoll. Melalui mediasi dari")
	sb.WriteString(" pihak Polres Indramayu dan menjaga kondusifitas maka dicapai kesepakatan untuk pemegang")
	sb.WriteString(" unit melunasi kredit kendaraan tersebut. Tetapi pemegang unit keberatan dengan nominal")
	sb.WriteString(" pelunasan yang timbul karena merasa sudah keluar uang sebesar Rp.0 untuk terima gadai")
	sb.WriteString(" unit tersebut dan juga perbaikan kendaraan selama ini serta pengurusan pajak kendaran.")
	sb.WriteString(" Sudah diajukan penawaran lelang khusus tapi pemegang unit tidak juga ada penyelesaian.")
	sb.WriteString(" Oleh tim excoll akhirnya unit dilakukan pengamanan asset.")

	p.SetFont(font, "", 10)
	p.MultiCell(box, lh, sb.String(), "", "J", false)
	p.SetY(p.GetY() + 3)

	p.MultiCell(box, lh, "Oleh karena kronologis tersebut di atas maka kami mengajukan penambahan biaya dengan total biaya Rp.20,000,000,- (include fee standard) dengan rincian sebagai berikut:", "", "J", false)

	box1 = 5
	box2 = 65
	box3 = 3
	box4 := box - box1 - box2 - box3
	p.SetY(p.GetY() + 3)

	p.CellFormat(box1, lh, "1)", "", 0, "L", false, 0, "")
	p.CellFormat(box2, lh, "Biaya Koordinasi & Mediasi", "", 0, "L", false, 0, "")
	p.CellFormat(box3, lh, ":", "", 0, "L", false, 0, "")
	p.CellFormat(box4, lh, "Rp 0,-", "", 1, "L", false, 0, "")

	p.CellFormat(box1, lh, "2)", "", 0, "L", false, 0, "")
	p.CellFormat(box2, lh, "Biaya Operasional & sukses fee", "", 0, "L", false, 0, "")
	p.CellFormat(box3, lh, ":", "", 0, "L", false, 0, "")
	p.CellFormat(box4, lh, "Rp 0,-", "", 1, "L", false, 0, "")

	p.SetFont(font, "B", 10)
	p.CellFormat(box1+box2, lh, "Total", "", 0, "C", false, 0, "")
	p.CellFormat(box3, lh, ":", "", 0, "L", false, 0, "")
	p.CellFormat(box4, lh, "Rp 0,-", "", 1, "L", false, 0, "")
	p.SetY(p.GetY() + 3)
	p.SetFont(font, "", 10)

	p.MultiCell(box, lh, "PT. SPRS dan team lapangan bersedia membuat pernyataan unit bebas kasus dan siap untuk dilelang. Demikian pengajuan ini kami sampaikan. Terima kasih", "", "J", false)

	/*
	   	lh = 5.5
	   	p.CellFormat(box1, lh, "No. Kontrak", "", 0, "L", false, 0, "")
	   	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	   	p.SetFont(font, "B", 10)
	   	p.CellFormat(box3, lh, "", "", 1, "L", false, 0, "")
	   	p.SetFont(font, "", 10)

	   	p.CellFormat(box1, lh, "Nama Customer", "", 0, "L", false, 0, "")
	   	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	   	p.SetFont(font, "B", 10)
	   	p.CellFormat(box3, lh, "3329130904840004", "", 1, "L", false, 0, "")
	   	p.SetFont(font, "", 10)

	   	p.CellFormat(box1, lh, "Type/No.Pol", "", 0, "L", false, 0, "")
	   	p.CellFormat(box2, lh, ":", "", 0, "L", false, 0, "")
	   	p.SetFont(font, "B", 10)
	   	p.CellFormat(box3, lh, "DIREKTUR PT SARANA PADMA RIDHO SEPUH", "", 1, "L", false, 0, "")
	   	p.SetFont(font, "", 10)

	   	p.SetY(p.GetY() + 1)
	   	html := p.HTMLBasicNew()
	   	_, lineHt := p.GetFontSize()
	   	lineHt += 2
	   	html.Write(lineHt, fmt.Sprintf("Dengan total biaya <b>Rp%s,-</b> (<i>** %s Rupiah</i>)",
	   		format_number(*total), strings.TrimSpace(convertion.Terbilang(*total))))

	   	html.Write(lineHt,
	   		fmt.Sprintf(`<br /><br />Bersama ini saya menyatakan dengan sebenarnya bahwa unit tersebut sudah bebas
	   dari kasus manapun dari masalah serta dapat segera dilakukan lelang, penggunaan
	   dari seluruh biaya tarik tersebut ( Rp%s,- ) adalah menjadi tanggung jawab saya.
	   Apabila ada tuntutan di kemudian hari ataupun masalah yang terjadi pada penggunaan
	   biaya maupun unit tersebut di atas, akan menjadi tanggung jawab saya sepenuhnya dan
	   tidak melibatkan pihak %s sesuai dengan perjanjian yang telah disepakati bersama,
	   serta %s terbebas dari segala tuntutan maupun permasalahan yang terjadi di kemudian hari.`,
	   			format_number(*total), finance.Name, finance.Name))
	   	html.Write(lineHt, `Demikian surat pernyataan ini saya buat. Tanpa ada paksaan dari pihak manapun.`)

	   	p.SetY(p.GetY() + 10)

	   	box1 = box * (45.0 / 100.0)
	   	box2 = box * (10.0 / 100.0)
	   	p.CellFormat(box1, lh, "Hormat saya,", "", 0, "C", false, 0, "")
	   	p.CellFormat(box2, lh, "", "", 0, "L", false, 0, "")
	   	p.CellFormat(box1, lh, "Mengetahui,", "", 1, "C", false, 0, "")

	   	p.SetY(p.GetY() + 15)
	   	p.SetFont(font, "BU", 10)
	   	p.CellFormat(box1, lh, "Nama : DEDDY PRANOTO,", "", 0, "C", false, 0, "")
	   	p.CellFormat(box2, lh, "", "", 0, "L", false, 0, "")
	   	p.CellFormat(box1, lh, "Nama :  ISKANDAR.G", "", 1, "C", false, 0, "")

	   	p.SetFont(font, "", 10)
	   	p.CellFormat(box1, lh, "Direktur PT SARANA PADMA RIDHO SEPUH", "", 0, "C", false, 0, "")
	   	p.CellFormat(box2, lh, "", "", 0, "L", false, 0, "")
	   	p.CellFormat(box1, lh, "Collection Officer Macet", "", 1, "C", false, 0, "")

	   	p.SetY(p.GetY() + 10)
	   	p.SetFont(font, "", 8)
	   	p.CellFormat(box, lh, "NOTARIS MAISARAH PANE, SH", "", 1, "L", false, 0, "")
	   	p.CellFormat(box, lh, "KEMENKUMHAM.RI NO. AHU-0109242.AH.01.11.tahun 2021-08-24", "", 1, "L", false, 0, "")
	   	p.CellFormat(box, lh, "42.758.225.9-437.000", "", 1, "L", false, 0, "")
	*/
}
