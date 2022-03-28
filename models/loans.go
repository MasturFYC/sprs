package models

type Lent struct {
	OrderID   int64      `json:"orderId"`             // Nomor order unit
	Name      string     `json:"name"`                // Nama peminjam kendaraan
	Descripts NullString `json:"descripts,omitempty"` // keterangan
	Street    NullString `json:"street,omitempty"`    // nama jalan
	City      NullString `json:"city,omitempty"`      // nama kota
	Phone     NullString `json:"phone,omitempty"`     // no. telephone
	Cell      NullString `json:"cell,omitempty"`      // no. hp / cellular
	Zip       NullString `json:"zip,omitempty"`       // no. kode pos
}

type LentDetail struct {
	OrderID   int64      `json:"orderId"`
	PaymentAt string     `json:"paymentAt"`
	ID        int64      `json:"id"`
	Descripts NullString `json:"descripts,omitempty"`
	Debt      float64    `json:"debt"`
	Cred      float64    `json:"cred"`
}

type Loan struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Descripts NullString `json:"descripts,omitempty"`
	Street    NullString `json:"street,omitempty"`
	City      NullString `json:"city,omitempty"`
	Phone     NullString `json:"phone,omitempty"`
	Cell      NullString `json:"cell,omitempty"`
	Zip       NullString `json:"zip,omitempty"`
}

type LoanDetail struct {
	LoanID    int64      `json:"orderId"`
	ID        int64      `json:"id"`
	PaymentAt string     `json:"paymentAt"`
	Descripts NullString `json:"descripts,omitempty"`
	Debt      float64    `json:"debt"`
	Cred      float64    `json:"cred"`
}
