package models

type Invoice struct {
	ID          int64      `json:"id"`
	InvoiceAt   string     `json:"invoiceAt"`
	PaymentTerm int32      `json:"paymentTerm"`
	DueAt       string     `json:"dueAt"`
	Salesman    string     `json:"salesman"`
	FinanceID   int32      `json:"financeId"`
	Memo        NullString `json:"memo"`
	Total       float64    `json:"total"`
	AccountId   float64    `json:"accountId"`
}

type InvoiceDetail struct {
	InvoiceID int64   `json:"invoiceId"`
	ID        int64   `json:"id"`
	OrderID   int64   `json:"orderId"`
	Price     float64 `json:"price"`
	Tax       float64 `json:"Tax"`
}
