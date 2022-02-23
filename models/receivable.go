package models

import "time"

// table tunggakan
type Receivable struct {
	// nomor order
	OrderID int64 `json:"orderId"`

	// tgl perjanjian
	CovenantAt time.Time `json:"covenantAt"`

	// tgl jatuh tempo
	DueAt time.Time `json:"dueAt"`

	// angsuran per bulan
	MortgageByMonth float64 `json:"mortgageByMonth"`

	// angsuran tunggakan
	MortgageReceivable float64 `json:"mortgageReceivable"`

	// denda berjalan
	RunningFine float64 `json:"runningFine"`

	// sisa denda
	RestFine float64 `json:"restFine"`

	// jasa penagihan
	BillService float64 `json:"billService"`

	// bayar titipan
	PayDeposit float64 `json:"payDeposit"`

	// sisa piutang
	RestReceivable float64 `json:"restReceivable"`

	// sisa pokok
	RestBase float64 `json:"restBase"`

	// jangka waktu
	DayPeriod int `json:"dayPeriod"`

	// angsuran yg ke
	MortgageTo int `json:"mortgageTo"`

	// jumlah hari angsuran
	DayCount int64 `json:"dayCount"`
}
