package models

import "time"

// table tugas tugas
type Task struct {
	OrderID int64 `json:"orderId"`

	// periode awal tagihan
	PeriodFrom time.Time `json:"periodFrom"`
	// periode akhir tagihan
	PeriodTo time.Time `json:"periodTo"`

	// nama penerima tugas
	RecipientName string `json:"recipientName"`
	// jabatan penerima tugas
	RecipientPosition string `json:"recipientPosition"`

	// nama pemberi tugas
	GiverName string `json:"giverName"`
	// jabatan pemberi tugas
	GiverPosition string `json:"giverPosition"`

	// keterangan
	Descriptions string `json:"descriptions"`
}
