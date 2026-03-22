package service

import (
	"MISPRIS/internal/domain"
	"MISPRIS/internal/repository"
	"context"

	"github.com/jmoiron/sqlx"
)

type EmobileService interface {
	Create(ctx context.Context, name string, powerPointID string,
		batteryID string, charg_sysID string, chassisID string, bodyID string, electonicsID string) (string, error)
	Update(ctx context.Context, id string, name string, powerPointID string,
		batteryID string, charg_sysID string, chassisID string, bodyID string, electonicsID string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Emobile, error)
	GetByID(ctx context.Context, id string) (*domain.Emobile, error)
}

type BatteryService interface {
	Create(ctx context.Context, name string, batteryType string, batteryCapacity string, batteryInfo string) (string, error)
	Update(ctx context.Context, id string, name string, batteryType string, batteryCapacity string, batteryInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Battery, error)
	GetByID(ctx context.Context, id string) (*domain.Battery, error)
}

type BodyService interface {
	Create(ctx context.Context, carcassID string, doorsID string, wingsID string) (string, error)
	Update(ctx context.Context, id string, carcassID string, doorsID string, wingsID string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Body, error)
	GetByID(ctx context.Context, id string) (*domain.Body, error)
}

type CarcassService interface {
	Create(ctx context.Context, carcassName string, carcassInfo string) (string, error)
	Update(ctx context.Context, id string, carcassName string, carcassInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Carcass, error)
	GetByID(ctx context.Context, id string) (*domain.Carcass, error)
}

type DoorsService interface {
	Create(ctx context.Context, doorsName string, doorInfo string) (string, error)
	Update(ctx context.Context, id string, doorsName string, doorInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Doors, error)
	GetByID(ctx context.Context, id string) (*domain.Doors, error)
}

type WingsService interface {
	Create(ctx context.Context, wingsName string, wingsInfo string) (string, error)
	Update(ctx context.Context, id string, wingsName string, wingsInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Wings, error)
	GetByID(ctx context.Context, id string) (*domain.Wings, error)
}

type ChargerSystemService interface {
	Create(ctx context.Context, chargerID string, connectorID string) (string, error)
	Update(ctx context.Context, id string, chargerID string, connectorID string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.ChargerSystem, error)
	GetByID(ctx context.Context, id string) (*domain.ChargerSystem, error)
}

type ChargerService interface {
	Create(ctx context.Context, chargerName, chargerInfo string) (string, error)
	Update(ctx context.Context, id string, chargerName, chargerInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Charger, error)
	GetByID(ctx context.Context, id string) (*domain.Charger, error)
}

type ConnectorService interface {
	Create(ctx context.Context, connectorName, connectorInfo string) (string, error)
	Update(ctx context.Context, id string, connectorName, connectorInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Connector, error)
	GetByID(ctx context.Context, id string) (*domain.Connector, error)
}

type ChassisService interface {
	Create(ctx context.Context, frameID, suspensionID, breakSystemID string) (string, error)
	Update(ctx context.Context, id string, frameID, suspensionID, breakSystemID string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Chassis, error)
	GetByID(ctx context.Context, id string) (*domain.Chassis, error)
}

type FrameService interface {
	Create(ctx context.Context, frameName, frameInfo string) (string, error)
	Update(ctx context.Context, id string, frameName, frameInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Frame, error)
	GetByID(ctx context.Context, id string) (*domain.Frame, error)
}

type SuspensionService interface {
	Create(ctx context.Context, suspensionName, suspensionInfo string) (string, error)
	Update(ctx context.Context, id string, suspensionName, suspensionInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Suspension, error)
	GetByID(ctx context.Context, id string) (*domain.Suspension, error)
}

type BreakSystemService interface {
	Create(ctx context.Context, breakSystemName, breakSystemInfo string) (string, error)
	Update(ctx context.Context, id string, breakSystemName, breakSystemInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.BreakSystem, error)
	GetByID(ctx context.Context, id string) (*domain.BreakSystem, error)
}

type ElectronicsService interface {
	Create(ctx context.Context, controllerID, sensorID, wiringID string) (string, error)
	Update(ctx context.Context, id string, controllerID, sensorID, wiringID string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Electronics, error)
	GetByID(ctx context.Context, id string) (*domain.Electronics, error)
}

type ControllerService interface {
	Create(ctx context.Context, controllerName, controllerInfo string) (string, error)
	Update(ctx context.Context, id string, controllerName, controllerInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Controller, error)
	GetByID(ctx context.Context, id string) (*domain.Controller, error)
}

type SensorService interface {
	Create(ctx context.Context, sensorName string, sensorInfo string) (string, error)
	Update(ctx context.Context, id string, sensorName string, sensorInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Sensor, error)
	GetByID(ctx context.Context, id string) (*domain.Sensor, error)
}

type WiringService interface {
	Create(ctx context.Context, wiringName string, wiringInfo string) (string, error)
	Update(ctx context.Context, id string, wiringName string, wiringInfo string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Wiring, error)
	GetByID(ctx context.Context, id string) (*domain.Wiring, error)
}

type PowerPointService interface {
	Create(ctx context.Context, engineID, invertorID, gearboxID string) (string, error)
	Update(ctx context.Context, id string, engineID, invertorID, gearboxID string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.PowerPoint, error)
	GetByID(ctx context.Context, id string) (*domain.PowerPoint, error)
}

type EngineService interface {
	Create(ctx context.Context, name, engineType, info string) (string, error)
	Update(ctx context.Context, id, name, engineType, info string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Engine, error)
	GetByID(ctx context.Context, id string) (*domain.Engine, error)
}

type InverterService interface {
	Create(ctx context.Context, name, info string) (string, error)
	Update(ctx context.Context, id, name, info string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Inverter, error)
	GetByID(ctx context.Context, id string) (*domain.Inverter, error)
}

type GearboxService interface {
	Create(ctx context.Context, name, info string) (string, error)
	Update(ctx context.Context, id, name, info string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Gearbox, error)
	GetByID(ctx context.Context, id string) (*domain.Gearbox, error)
}

type Service struct {
	PowerPoint PowerPointService
	Engine     EngineService
	Inverter   InverterService
	Gearbox    GearboxService
	//////////////////////////////////////
	Electronics ElectronicsService
	Controller  ControllerService
	Sensor      SensorService
	Wiring      WiringService
	//////////////////////////////////////////////
	ChargerSystem ChargerSystemService
	Charger       ChargerService
	Connector     ConnectorService
	////////////////////////////////////////////
	Chassis     ChassisService
	Frame       FrameService
	Suspension  SuspensionService
	BreakSystem BreakSystemService
	/////////////////////////////////////////////
	Battery BatteryService
	/////////////////////////////////////////////
	Emobile EmobileService

	/////////////////////////////////
	Carcass CarcassService
	Doors   DoorsService
	Wings   WingsService
	Body    BodyService
	///////////////////////////////////
}

func NewService(db *sqlx.DB, repo *repository.Repository) *Service {
	frame := NewFrameService(repo.Frame)
	suspension := NewSuspensionService(repo.Suspension)
	breakSystem := NewBreakSystemService(repo.BreakSystem)

	charger := NewChargerService(repo.Charger)
	connector := NewConnectorService(repo.Connector)

	carcass := NewCarcassService(repo.Carcass)
	doors := NewDoorsService(repo.Doors)
	wings := NewWingsService(repo.Wings)

	sensor := NewSensorService(repo.Sensor)
	wiring := NewWiringService(repo.Wiring)
	controller := NewControllerService(repo.Controller)

	engine := NewEngineService(repo.Engine)
	inverter := NewInverterService(repo.Inverter)
	gearbox := NewGearboxService(repo.Gearbox)

	return &Service{
		Charger:       charger,
		Connector:     connector,
		ChargerSystem: NewChargerSystemService(db, repo.ChargerSystem, charger, connector),

		Frame:       frame,
		Suspension:  suspension,
		BreakSystem: breakSystem,
		Chassis:     NewChassisService(db, repo.Chassis, frame, suspension, breakSystem),

		Sensor:      sensor,
		Wiring:      wiring,
		Controller:  controller,
		Electronics: NewElectronicsService(db, repo.Electronics, wiring, controller, sensor),

		Carcass: carcass,
		Doors:   doors,
		Wings:   wings,
		Body:    NewBodyService(db, repo.Body, carcass, doors, wings),

		Battery: NewBatteryService(repo.Battery, db),

		Engine:     engine,
		Inverter:   inverter,
		Gearbox:    gearbox,
		PowerPoint: NewPowerPointService(repo.PowerPoint, db, engine, inverter, gearbox),

		Emobile: NewEmobileService(db, repo.Emobile,
			NewChargerSystemService(db, repo.ChargerSystem, charger, connector),
			NewBodyService(db, repo.Body, carcass, doors, wings),
			NewElectronicsService(db, repo.Electronics, wiring, controller, sensor),
			NewChassisService(db, repo.Chassis, frame, suspension, breakSystem),
			NewBatteryService(repo.Battery, db),
			NewPowerPointService(repo.PowerPoint, db, engine, inverter, gearbox)),
	}
}
