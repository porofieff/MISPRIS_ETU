package schema

type CreateElectronicsInput struct {
	ControllerID string `json:"controller_id" binding:"required"`
	SensorID     string `json:"sensor_id" binding:"required"`
	WiringID     string `json:"wiring_id" binding:"required"`
}

type CreateControllerInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

type CreateSensorInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

type CreateWiringInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}
