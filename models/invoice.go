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
	Tax         float64    `json:"tax"`
	AccountId   float64    `json:"accountId"`
}

type InvoiceDetail struct {
	InvoiceID int64 `json:"invoiceId"`
	OrderID   int64 `json:"orderId"`
}