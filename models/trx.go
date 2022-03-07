package models

type AccType struct {
	ID           int32      `json:"id"`
	Name         string     `json:"name"`
	Descriptions NullString `json:"descriptions"`
}

type AccCode struct {
	ID           int32      `json:"id"`
	Name         string     `json:"name"`
	AccTypeID    int32      `json:"accTypeId"`
	Descriptions NullString `json:"descriptions"`
}

type TrxType struct {
	ID           int32      `json:"id"`
	Name         string     `json:"name"`
	Descriptions NullString `json:"descriptions"`
}

type Trx struct {
	ID           int64       `json:"id"`
	TrxTypeID    int32       `json:"trxTypeId"`
	RefID        int64       `json:"refId"`
	Division     string      `json:"divison"`
	Descriptions string      `json:"descriptions"`
	TrxDate      string      `json:"trxDate"`
	Memo         NullString  `json:"memo"`
	Saldo        float64     `json:"saldo"`
	Details      []TrxDetail `json:"details,omitempty"`
}

type TrxDetail struct {
	ID        int64   `json:"id"`
	AccCodeID int32   `json:"accCodeId"`
	TrxID     int64   `json:"trxId"`
	Debt      float64 `json:"debt"`
	Cred      float64 `json:"cred"`
}

type AccCodeType struct {
	ID           int32      `json:"id"`
	Name         string     `json:"name"`
	TypeID       int32      `json:"typeId"`
	TypeName     string     `json:"typeName"`
	Descriptions NullString `json:"descriptions"`
}
