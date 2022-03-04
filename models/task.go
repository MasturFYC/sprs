package models

//import "time"

// table tugas tugas
type Task struct {
	OrderID int64 `json:"orderId"`

	// periode awal tagihan
	PeriodFrom string `json:"periodFrom"`
	// periode akhir tagihan
	PeriodTo string `json:"periodTo"`

	// nama penerima tugas
	RecipientName string `json:"recipientName"`
	// jabatan penerima tugas
	RecipientPosition string `json:"recipientPosition"`

	// nama pemberi tugas
	GiverName string `json:"giverName"`
	// jabatan pemberi tugas
	GiverPosition string `json:"giverPosition"`

	// keterangan
	Descriptions NullString `json:"descriptions"`
}
