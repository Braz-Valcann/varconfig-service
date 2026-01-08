package service

import (
	"context"
	"time"

	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig"
)

type Service struct {
	repo varconfig.IRepository
}

var _ varconfig.IService = (*Service)(nil)

func New(repo varconfig.IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, orgID, benchmarkID string, payload map[string]interface{}) (*varconfig.VarConfig, error) {

	now := time.Now().UTC().Format(time.RFC3339)

	vc := &varconfig.VarConfig{
		ID:          time.Now().UnixNano(),
		OrgID:       orgID,
		BenchmarkID: benchmarkID,
		Payload:     payload,
		CreatedAt:   now,
		UpdateAt:    now,
	}

	err := s.repo.Create(ctx, vc)

	return vc, err
}

func (s *Service) Get(ctx context.Context, orgID, benchmarkID string, id int64) (*varconfig.VarConfig, error) {
	return s.repo.Get(ctx, orgID, benchmarkID, id)
}

func (s *Service) List(ctx context.Context, orgID, benchmarkID string) ([]*varconfig.VarConfig, error) {
	return s.repo.List(ctx, orgID, benchmarkID)
}

func (s *Service) Update(ctx context.Context, orgID, benchmarkID string, id int64, payload map[string]interface{}) (*varconfig.VarConfig, error) {

	vc := &varconfig.VarConfig{
		ID:          id,
		OrgID:       orgID,
		BenchmarkID: benchmarkID,
		Payload:     payload,
		UpdateAt:    time.Now().UTC().Format(time.RFC3339),
	}

	err := s.repo.Update(ctx, vc)
	return vc, err
}

func (s *Service) Delete(ctx context.Context, orgID, benchmarkID string, id int64) error {
	return s.repo.Delete(ctx, orgID, benchmarkID, id)
}
