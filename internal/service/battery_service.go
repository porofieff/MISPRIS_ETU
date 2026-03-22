package service

import (
	"context"

	"MISPRIS/internal/domain"
	"MISPRIS/internal/repository"

	"github.com/google/uuid"
)

type batteryService struct {
	repo repository.BatteryRepository
}

func NewBatteryService(repo repository.BatteryRepository) BatteryService {
	return &batteryService{repo: repo}
}

func (s *batteryService) Create(ctx context.Context, name string, batteryType string, batteryCapacity string, batteryInfo string) (string, error) {
	b := &domain.Battery{
		ID:              uuid.New().String(),
		BatteryName:     name,
		BatteryType:     batteryType,
		BatteryCapacity: batteryCapacity,
		BatteryInfo:     batteryInfo,
	}
	return s.repo.Create(ctx, b)
}

func (s *batteryService) Update(ctx context.Context, id string, name string, batteryType string, batteryCapacity string, batteryInfo string) error {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	existing.BatteryName = name
	existing.BatteryType = batteryType
	existing.BatteryCapacity = batteryCapacity
	existing.BatteryInfo = batteryInfo
	return s.repo.Update(ctx, existing)
}

func (s *batteryService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *batteryService) List(ctx context.Context) ([]domain.Battery, error) {
	batteries, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]domain.Battery, len(batteries))
	for i, b := range batteries {
		result[i] = *b
	}
	return result, nil
}

func (s *batteryService) GetByID(ctx context.Context, id string) (*domain.Battery, error) {
	return s.repo.GetByID(ctx, id)
}

