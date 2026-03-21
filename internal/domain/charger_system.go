package domain

type ChargerSystem struct {
	ID          int64 `json:"charger_system_id" db:"charger_system_id"`
	ChargerID   int64 `json:"charger_id"        db:"charger_id"`
	ConnectorID int64 `json:"connector_id" db:"connector_id"`
}

type Charger struct {
	ID   int64  `json:"charger_id"        db:"charger_id"`
	Name string `json:"charger_name"      db:"charger_name"`
	Info string `json:"charger_info"      db:"charger_info"`
}

type Connector struct {
	ID   int64  `json:"connector_id"        db:"connector_id"`
	Name string `json:"connector_name"      db:"connector_name"`
	Info string `json:"connector_info"      db:"connector_info"`
}
