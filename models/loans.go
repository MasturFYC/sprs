package models

type Lent struct {
	OrderID   int64      `json:"orderId"`          // Nomor order unit
	Name      string     `json:"name"`             // Nama peminjam kendaraan
	Descripts string     `json:"descripts"`        // Nama peminjam kendaraan
	Street    NullString `json:"street,omitempty"` // nama jalan
	City      NullString `json:"city,omitempty"`   // nama kota
	Phone     NullString `json:"phone,omitempty"`  // no. telephone
	Cell      NullString `json:"cell,omitempty"`   // no. hp / cellular
	Zip       NullString `json:"zip,omitempty"`    // no. kode pos
}

type Loan struct {
	ID     int64      `json:"id"`
	Name   string     `json:"name"`
	Persen float32    `json:"persen"`
	Street NullString `json:"street,omitempty"`
	City   NullString `json:"city,omitempty"`
	Phone  NullString `json:"phone,omitempty"`
	Cell   NullString `json:"cell,omitempty"`
	Zip    NullString `json:"zip,omitempty"`
}
