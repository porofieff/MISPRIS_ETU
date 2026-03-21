package domain

type PowerPoint struct {
	ID         int64 `json:"power_point_id" db:"power_point_id"`
	EngineID   int64 `json:"engine_id"        db:"engine_id"`
	InverterID int64 `json:"inverter_id" db:"inverter_id"`
	GearboxID  int64 `json:"gearbox_id" db:"gearbox_id"`
}

type Engine struct {
	ID         int64  `json:"engine_id" db:"engine_id"`
	EngineName string `json:"engine_name" db:"engine_name"`
	EngineType string `json:"engine_type" db:"engine_type"`
	EngineInfo string `json:"engine_info" db:"engine_info"`
}

type Inverter struct {
	ID           int64  `json:"inverter_id" db:"inverter_id"`
	InverterName string `json:"inverter_name" db:"inverter_name"`
	InverterInfo string `json:"inverter_info" db:"inverter_info"`
}

type Gearbox struct {
	ID          int64  `json:"gearbox_id" db:"gearbox_id"`
	GearboxName string `json:"gearbox_name" db:"gearbox_name"`
	GearboxInfo string `json:"gearbox_info" db:"gearbox_info"`
}
