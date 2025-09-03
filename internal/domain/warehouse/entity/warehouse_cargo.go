package entity

import "time"

type WarehouseCargo struct {
	ID          uint64 `json:"id"`
	WarehouseID uint64 `json:"warehouse_id"`
	CargoTypeID uint64 `json:"cargo_type_id"`
	Reserved    uint64 `json:"reserved"`
	OnHand      uint64 `json:"on_hand"`

	Warehouse *Warehouse `json:"warehouse"`
	CargoType *CargoType `json:"cargo_type"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
