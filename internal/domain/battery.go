package domain

type Battery struct {
	ID              string `json:"battery_id"    db:"battery_id"`
	BatteryName     string `json:"battery_name" db:"battery_name"`
	BatteryType     string `json:"battery_type"   db:"battery_type"`
	BatteryCapacity string `json:"battery_capacity" db:"battery_capacity"`
	BatteryInfo     string `json:"battery_info"   db:"battery_info"`
}
