package domain

type Emobile struct {
	ID              int64  `json:"emobile_id"        db:"emobile_id"`
	Name            string `json:"emobile_name"       db:"emobile_name"`
	PowerPointID    int64  `json:"power_point_id"     db:"power_point_id"`
	BatteryID       int64  `json:"battery_id"         db:"battery_id"`
	ChargerSystemID int64  `json:"charger_system_id"  db:"charger_system_id"`
	ChassisID       int64  `json:"chassis_id"         db:"chassis_id"`
	BodyID          int64  `json:"body_id"            db:"body_id"`
	ElectronicsID   int64  `json:"electronics_id"     db:"electronics_id"`
}
