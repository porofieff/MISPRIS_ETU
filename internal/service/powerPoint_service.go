package service

import (
	"context"

	"MISPRIS/internal/domain"
	"MISPRIS/internal/repository"

	"github.com/google/uuid"
)

// Power Point
type powerPointService struct {
	repo repository.PowerPointRepository
}

func NewPowerPointService(repo repository.PowerPointRepository) PowerPointService {
	return &powerPointService{repo: repo}
}

func (s *powerPointService) Create(ctx context.Context, engineID, invertorID, gearboxID string) (string, error) {
	p := &domain.PowerPoint{
		ID:         uuid.New().String(),
		EngineID:   engineID,
		InverterID: invertorID,
		GearboxID:  gearboxID,
	}
	return s.repo.Create(ctx, p)
}

func (s *powerPointService) Update(ctx context.Context, id string, engineID, invertorID, gearboxID string) error {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	existing.EngineID = engineID
	existing.InverterID = invertorID
	existing.GearboxID = gearboxID
	return s.repo.Update(ctx, existing)
}

func (s *powerPointService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *powerPointService) List(ctx context.Context) ([]domain.PowerPoint, error) {
	rows, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]domain.PowerPoint, len(rows))
	for i, r := range rows {
		result[i] = *r
	}
	return result, nil
}

func (s *powerPointService) GetByID(ctx context.Context, id string) (*domain.PowerPoint, error) {
	return s.repo.GetByID(ctx, id)
}

// Engine
type engineService struct {
	repo repository.EngineRepository
}

func NewEngineService(repo repository.EngineRepository) EngineService {
	return &engineService{repo: repo}
}

func (s *engineService) Create(ctx context.Context, name, engineType, info string) (string, error) {
	e := &domain.Engine{
		ID:         uuid.New().String(),
		EngineName: name,
		EngineType: engineType,
		EngineInfo: info,
	}
	return s.repo.Create(ctx, e)
}

func (s *engineService) Update(ctx context.Context, id, name, engineType, info string) error {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	existing.EngineName = name
	existing.EngineType = engineType
	existing.EngineInfo = info
	return s.repo.Update(ctx, existing)
}

func (s *engineService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *engineService) List(ctx context.Context) ([]domain.Engine, error) {
	engines, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]domain.Engine, len(engines))
	for i, e := range engines {
		result[i] = *e
	}
	return result, nil
}

func (s *engineService) GetByID(ctx context.Context, id string) (*domain.Engine, error) {
	return s.repo.GetByID(ctx, id)
}

// Inverter
type inverterService struct {
	repo repository.InverterRepository
}

func NewInverterService(repo repository.InverterRepository) InverterService {
	return &inverterService{repo: repo}
}

func (s *inverterService) Create(ctx context.Context, name, info string) (string, error) {
	i := &domain.Inverter{
		ID:           uuid.New().String(),
		InverterName: name,
		InverterInfo: info,
	}
	return s.repo.Create(ctx, i)
}

func (s *inverterService) Update(ctx context.Context, id, name, info string) error {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	existing.InverterName = name
	existing.InverterInfo = info
	return s.repo.Update(ctx, existing)
}

func (s *inverterService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *inverterService) List(ctx context.Context) ([]domain.Inverter, error) {
	inverters, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]domain.Inverter, len(inverters))
	for i, inv := range inverters {
		result[i] = *inv
	}
	return result, nil
}

func (s *inverterService) GetByID(ctx context.Context, id string) (*domain.Inverter, error) {
	return s.repo.GetByID(ctx, id)
}

// Gearbox
type gearboxService struct {
	repo repository.GearboxRepository
}

func NewGearboxService(repo repository.GearboxRepository) GearboxService {
	return &gearboxService{repo: repo}
}

func (s *gearboxService) Create(ctx context.Context, name, info string) (string, error) {
	g := &domain.Gearbox{
		ID:          uuid.New().String(),
		GearboxName: name,
		GearboxInfo: info,
	}
	return s.repo.Create(ctx, g)
}

func (s *gearboxService) Update(ctx context.Context, id, name, info string) error {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	existing.GearboxName = name
	existing.GearboxInfo = info
	return s.repo.Update(ctx, existing)
}

func (s *gearboxService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *gearboxService) List(ctx context.Context) ([]domain.Gearbox, error) {
	gearboxes, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]domain.Gearbox, len(gearboxes))
	for i, g := range gearboxes {
		result[i] = *g
	}
	return result, nil
}

func (s *gearboxService) GetByID(ctx context.Context, id string) (*domain.Gearbox, error) {
	return s.repo.GetByID(ctx, id)
}

