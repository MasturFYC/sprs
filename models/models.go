package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
)

type NullFloat64 sql.NullFloat64

// Scan implements the Scanner interface.
func (ni *NullFloat64) Scan(value interface{}) error {
	var i sql.NullFloat64
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullFloat64{i.Float64, false}
	} else {
		*ni = NullFloat64{i.Float64, true}
	}
	return nil
}

// MarshalJSON for NullInt64
func (ni *NullFloat64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
		//return []byte("0"), nil
	}
	return json.Marshal(ni.Float64)
}

type NullString string

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	strVal, ok := value.(string)
	if !ok {
		return errors.New("column is not a string")
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
	ActionAt     string     `json:"actionAt"`
	OrderId      int64      `json:"orderId"`
	Pic          string     `json:"pic"`
	Descriptions NullString `json:"descriptions"`
	// kode tindakan
	//Code     string     `json:"code"`
	FileName NullString `json:"fileName"`
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
	OrderID         int64      `json:"orderId"`
	Name            string     `json:"name"`
	AgreementNumber NullString `json:"agreementNumber"`
	PaymentType     string     `json:"paymentType"`
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
	GroupID   int        `json:"groupId"`
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
	MerkID  int    `json:"merkId"`
	Merk    Merk   `json:"merk,omitempty"`
	Wheel   Wheel  `json:"wheel,omitempty"`
}

// ------------------------------------------

type Property struct {
	ID   int64  `json:"id"`
	Name string `json:"string"`
}

type SearchGroup struct {
	Txt string `json:"txt"`
}

type FinanceGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// type PropertyMap map[string]interface{}

// func (p PropertyMap) Value() (driver.Value, error) {
// 	j, err := json.Marshal(p)
// 	return j, err
// }

// func (p *PropertyMap) Scan(src interface{}) error {
// 	source, ok := src.([]byte)
// 	if !ok {
// 		return errors.New("Type assertion .([]byte) failed.")
// 	}

// 	var i interface{}
// 	err := json.Unmarshal(source, &i)
// 	if err != nil {
// 		return err
// 	}

// 	*p, ok = i.(map[string]interface{})
// 	if !ok {
// 		return errors.New("Type assertion .(map[string]interface{}) failed.")
// 	}

// 	return nil
// }
