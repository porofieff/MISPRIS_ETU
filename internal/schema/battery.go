package schema

type CreateBatteryInput struct {
	Name     string `json:"name" binding:"required"`
	Type     string `json:"battery_type"`
	Capacity string `json:"battery_capacity"`
	Info     string `json:"battery_info"`
}
