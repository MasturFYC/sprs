package models

type AccGroup struct {
	ID           int32      `json:"id"`
	Name         string     `json:"name"`
	Descriptions NullString `json:"descriptions"`
}
type AccType struct {
	GroupID      int32      `json:"groupId"`
	ID           int32      `json:"id"`
	Name         string     `json:"name"`
	Descriptions NullString `json:"descriptions"`
}

type AccCode struct {
	TypeID           int32      `json:"typeId"`
	ID               int32      `json:"id"`
	Name             string     `json:"name"`
	Descriptions     NullString `json:"descriptions"`
	IsActive         bool       `json:"isActive"`
	IsAutoDebet      bool       `json:"isAutoDebet"`
	ReceivableOption int        `json:"option"`
}

type AccInfo struct {
	AccCode
	TypeName  string     `json:"typeName"`
	TypeDesc  NullString `json:"typeDesc"`
	GroupName string     `json:"groupName"`
	GroupDesc NullString `json:"groupDesc"`
}

type Trx struct {
	ID           int64       `json:"id"`
	RefID        int64       `json:"refId"`
	Division     string      `json:"division"`
	Descriptions string      `json:"descriptions"`
	TrxDate      string      `json:"trxDate"`
	Memo         NullString  `json:"memo"`
	Saldo        float64     `json:"saldo"`
	Details      []TrxDetail `json:"details,omitempty"`
}

type TrxDetailsToken struct {
	Trx     Trx         `json:"trx"`
	Details []TrxDetail `json:"details"`
	Token   string      `json:"token"`
}

type TrxDetail struct {
	ID     int64   `json:"id"`
	CodeID int32   `json:"codeId"`
	TrxID  int64   `json:"trxId"`
	Debt   float64 `json:"debt"`
	Cred   float64 `json:"cred"`
}

type AccCodeType struct {
	ID           int32      `json:"id"`
	Name         string     `json:"name"`
	TypeID       int32      `json:"typeId"`
	TypeName     string     `json:"typeName"`
	Descriptions NullString `json:"descriptions"`
	IsActive     bool       `json:"isActive"`
}
