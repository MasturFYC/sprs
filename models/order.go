package models

//import "time"

// table order SPK
type Order struct {
	ID int64 `json:"id"`

	// nomor SPK
	Name string `json:"name"`
	// tanggal penerimaan SPK
	OrderAt     string     `json:"orderAt"`
	PrintedAt   string     `json:"printedAt"`
	BtFinance   float64    `json:"btFinance"`
	BtPercent   float32    `json:"btPercent"`
	BtMatel     float64    `json:"btMatel"`
	Ppn         float32    `json:"ppn"`
	Nominal     float64    `json:"nominal"`
	Subtotal    float64    `json:"subtotal"`
	UserName    string     `json:"userName"`
	VerifiedBy  NullString `json:"verifiedBy"`
	ValidatedBy NullString `json:"validatedBy"`
	FinanceID   int        `json:"financeId"`
	BranchID    int        `json:"branchId"`
	IsStnk      bool       `json:"isStnk"`
	StnkPrice   float64    `json:"stnkPrice"`

	// * badan keuangan
	Finance Finance `json:"finance,omitempty"`
	// * cabang yg menangani
	Branch Branch `json:"branch,omitempty"`

	// data pelanggan
	Customer Customer `json:"customer,omitempty"`
	// data tunggakan
	Receivable Receivable `json:"receivable,omitempty"`
	// data unit kendaraan
	Unit Unit `json:"unit,omitempty"`
	// tindakan2 yg pernah dilakukan
	Actions []Action `json:"actions,omitempty"`

	// data pemberian tugas
	Task Task `json:"task,omitempty"`

	HomeAddress   HomeAddress   `json:"homeAddress,omitempty"`
	OfficeAddress OfficeAddress `json:"officeAddress,omitempty"`
	// alamat penagihan
	PostAddress PostAddress `json:"postAddress,omitempty"`
	// alamat kendaraan berdasarkan ktp
	KtpAddress KtpAddress `json:"ktpAddress,omitempty"`
}
