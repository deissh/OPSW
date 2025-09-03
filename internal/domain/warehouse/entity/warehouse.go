package entity

import "time"

type Warehouse struct {
	ID   uint64 `json:"warehouse_id"`
	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
