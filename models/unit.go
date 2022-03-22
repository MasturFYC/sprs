package models

type Unit struct {
	OrderID int64 `json:"orderId"`

	// nomor polisi / kendaraan
	Nopol string `json:"nopol"`

	// tahun keluaran
	Year int64 `json:"year"`

	// nomor rangka
	FrameNumber NullString `json:"frameNumber"`

	//nomor mesin
	MachineNumber NullString `json:"machineNumber"`

	//bpkb atas nama
	// BpkbName string `json:"bpkbName"`

	//warna kendaraan
	Color NullString `json:"color"`
	// Dealer   string `json:"dealer"`
	// Surveyor string `json:"surveyor"`

	// tipe unit
	TypeID int64 `json:"typeId"`
	Type   Type  `json:"type,omitempty"`

	// gudang lokasi unit
	WarehouseID int       `json:"warehouseId"`
	Warehouse   Warehouse `json:"warehouse,omitempty"`
}
