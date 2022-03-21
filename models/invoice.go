package models

type Invoice struct {
	ID          int64      `json:"id"`
	InvoiceAt   string     `json:"invoiceAt"`
	PaymentTerm int32      `json:"paymentTerm"`
	DueAt       string     `json:"dueAt"`
	Salesman    string     `json:"salesman"`
	FinanceID   int32      `json:"financeId"`
	Subtotal    float64    `json:"subtotal"`
	Ppn         float32    `json:"ppn"`
	Tax         float64    `json:"tax"`
	Total       float64    `json:"total"`
	AccountId   int32      `json:"accountId"`
	Memo        NullString `json:"memo"`
}

type InvoiceDetail struct {
	InvoiceID int64 `json:"invoiceId"`
	OrderID   int64 `json:"orderId"`
}
