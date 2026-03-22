package domain

type Body struct {
	ID        string `json:"body_id"    db:"body_id"`
	CarcassID string `json:"carcass_id" db:"carcass_id"`
	DoorsID   string `json:"doors_id"   db:"doors_id"`
	WingsID   string `json:"wings_id"   db:"wings_id"`
}

type Carcass struct {
	ID   string `json:"carcass_id"   db:"carcass_id"`
	Name string `json:"carcass_name" db:"carcass_name"`
	Info string `json:"carcass_info" db:"carcass_info"`
}

type Doors struct {
	ID   string `json:"doors_id"   db:"doors_id"`
	Name string `json:"doors_name" db:"doors_name"`
	Info string `json:"doors_info" db:"doors_info"`
}

type Wings struct {
	ID   string `json:"wings_id"   db:"wings_id"`
	Name string `json:"wings_name" db:"wings_name"`
	Info string `json:"wings_info" db:"wings_info"`
}
