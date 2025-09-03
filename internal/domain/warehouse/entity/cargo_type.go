package entity

type CargoType struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	Weight    uint64 `json:"weight"`
	Stackable bool   `json:"stackable"`
	Hazard    bool   `json:"hazard"`
	Fragility bool   `json:"fragility"`

	// Название для маппинга в dcs
	InternalName string `json:"internal_name"`
	// Размер груза для dcs, кол-во ракет/бомб/груза
	InternalSize uint64 `json:"internal_size"`
}
