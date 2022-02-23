package models

import (
	"database/sql/driver"
	"errors"
	"time"
)

type NullString string

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	strVal, ok := value.(string)
	if !ok {
		return errors.New("Column is not a string")
	}
	*s = NullString(strVal)
	return nil
}

func (s NullString) Value() (driver.Value, error) {
	if len(s) == 0 { // if nil or empty string
		return nil, nil
	}
	return string(s), nil
}

// table tindakan

type Action struct {
	ID           int64      `json:"id"`
	ActionAt     time.Time  `json:"actionAt"`
	OrderId      int64      `json:"orderId"`
	Pic          string     `json:"pic"`
	Descriptions NullString `json:"Descriptions"`
	// kode tindakan
	Code string `json:"code"`
}

// table cabang
type Branch struct {
	ID     int        `json:"id"`
	Street NullString `json:"street"`
	City   NullString `json:"city"`
	Phone  NullString `json:"phone"`
	Cell   NullString `json:"cell"`
	Zip    NullString `json:"zip"`
	Email  NullString `json:"email"`

	// nama cabang
	Name string `json:"name"`
	// nama kepala cabang
	HeadBranch string `json:"headBranch"`
}

type Customer struct {
	OrderID         int64  `json:"orderId"`
	Name            string `json:"name"`
	AgreementNumber string `json:"agreementNumber"`
	PaymentType     string `json:"paymentType"`
}

type Finance struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	ShortName string     `json:"shortName"`
	Street    NullString `json:"street"`
	City      NullString `json:"city"`
	Phone     NullString `json:"phone"`
	Cell      NullString `json:"cell"`
	Zip       NullString `json:"zip"`
	Email     NullString `json:"email"`
	Orders    []Order    `json:"orders,omitempty"`
}

type HomeAddress struct {
	OrderID int64      `json:"orderId"`
	Street  NullString `json:"street"`
	Region  NullString `json:"region"`
	City    NullString `json:"city"`
	Phone   NullString `json:"phone"`
	Cell    NullString `json:"cell"`
	Zip     NullString `json:"zip"`
}

type OfficeAddress struct {
	OrderID int64      `json:"orderId"`
	Street  NullString `json:"street"`
	Region  NullString `json:"region"`
	City    NullString `json:"city"`
	Phone   NullString `json:"phone"`
	Cell    NullString `json:"cell"`
	Zip     NullString `json:"zip"`
}

type PostAddress struct {
	OrderID int64      `json:"orderId"`
	Street  NullString `json:"street"`
	Region  NullString `json:"region"`
	City    NullString `json:"city"`
	Phone   NullString `json:"phone"`
	Cell    NullString `json:"cell"`
	Zip     NullString `json:"zip"`
}
type KtpAddress struct {
	OrderID int64      `json:"orderId"`
	Street  NullString `json:"street"`
	Region  NullString `json:"region"`
	City    NullString `json:"city"`
	Phone   NullString `json:"phone"`
	Cell    NullString `json:"cell"`
	Zip     NullString `json:"zip"`
}

// table merk kendaraan seperti honda, yamaha, dll
type Merk struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Types []Type `json:"types,omitempty"`
}

// table jenis kendaraan berdasarkan jumlah roda
type Wheel struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Types     []Type `json:"types,omitempty"`
}

// table gudang unit kendaraan
type Warehouse struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Descriptions NullString `json:"descriptions"`
	Units        []Unit     `json:"units,omitempty"`
}

// table tipe kendaraan berdasarkan nama kendaraan
type Type struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	WheelID int    `json:"wheelId"`
	MerkID  int    `json:"wheelMerk"`
	Merk    Merk   `json:"merk,omitempty"`
	Wheel   Wheel  `json:"wheel,omitempty"`
}

// ------------------------------------------

type Property struct {
	ID   int64  `json:"id"`
	Name string `json:"string"`
}
