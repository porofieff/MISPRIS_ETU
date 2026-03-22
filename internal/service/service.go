package service

import (
	"context"

	"MISPRIS/internal/domain"
	"MISPRIS/internal/repository"
)

type EmobileService interface {
	Create(ctx context.Context, id string, name string, powerPointID string,
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
	Create(ctx context.Context, id string, carcassID string, doorsID string, wingsID string) (string, error)
	Update(ctx context.Context, id string, carcassID string, doorsID string, wingsID string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Battery, error)
	GetByID(ctx context.Context, id string) (*domain.Battery, error)
}

type ChargerService interface {
	Create(ctx context.Context, id string, chargerID string, connectorID string) (string, error)
	Update(ctx context.Context, id string, chargerID string, connectorID string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Charger, error)
	GetByID(ctx context.Context, id string) (*domain.Charger, error)
}

type ChassisService interface {
	Create(ctx context.Context, id string, frameID, suspensionID, breakSystemID string) (string, error)
	Update(ctx context.Context, id string, frameID, suspensionID, breakSystemID string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Chassis, error)
	GetByID(ctx context.Context, id string) (*domain.Chassis, error)
}

type ElectronicsService interface {
	Create(ctx context.Context, id string, controllerID, sensorID, wiringID string) (string, error)
	Update(ctx context.Context, id string, controllerID, sensorID, wiringID string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.Electronics, error)
	GetByID(ctx context.Context, id string) (*domain.Electronics, error)
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
	PowerPoint  PowerPointService
	Electronics ElectronicsService
	ChargerSys  ChargerService
	Chassis     ChassisService
	Battery     BatteryService
	Emobile     EmobileService
	Body        BodyService
	Engine      EngineService
	Inverter    InverterService
	Gearbox     GearboxService
}

func NewService(repo *repository.Repository) Service {
	return &Service{
		PowerPoint:  NewPowerPointService(repo.PowerPoint),
		Electronics: NewElectronicsService(repo.Electronics),
		ChargerSys:  NewChargerSysService(repo.ChargerSystem),
		Chassis:     NewChassisService(repo.Chassis),
		Battery:     NewBatteryService(repo.Battery),
		Emobile:     NewEmobileService(repo.Emobile),
		Body:        NewBodyService(repo.Body),
		Engine:      NewEngineService(repo.Engine),
		Inverter:    NewInverterService(repo.Inverter),
		Gearbox:     NewGearboxService(repo.Gearbox),
	}
}
