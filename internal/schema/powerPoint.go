package schema

type CreatePowerPointInput struct {
	EngineID   string `json:"engine_id" binding:"required"`
	InverterID string `json:"inverter_id" binding:"required"`
	GearboxID  string `json:"gearbox_id" binding:"required"`
}
type CreateEngineInput struct {
	Name       string `json:"name" binding:"required"`
	EngineType string `json:"engine_type"`
	Info       string `json:"info"`
}

type CreateInverterInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

type CreateGearboxInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}
