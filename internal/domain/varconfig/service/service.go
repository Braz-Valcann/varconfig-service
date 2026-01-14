package service

import (
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

func (s *Service) Create(orgID, benchmarkID string, payload map[string]interface{}) (*varconfig.VarConfig, error) {

	now := time.Now().UTC().Format(time.RFC3339)

	vc := &varconfig.VarConfig{
		ID:          time.Now().UnixNano(),
		OrgID:       orgID,
		BenchmarkID: benchmarkID,
		Payload:     payload,
		CreatedAt:   now,
		UpdateAt:    now,
	}

	err := s.repo.Create(vc)

	return vc, err
}

func (s *Service) Get(orgID, benchmarkID string, id int64) (*varconfig.VarConfig, error) {
	return s.repo.Get(orgID, benchmarkID, id)
}

func (s *Service) List(orgID, benchmarkID string) ([]*varconfig.VarConfig, error) {
	return s.repo.List(orgID, benchmarkID)
}

func (s *Service) Update(orgID, benchmarkID string, id int64, payload map[string]interface{}) (*varconfig.VarConfig, error) {

	vc := &varconfig.VarConfig{
		ID:          id,
		OrgID:       orgID,
		BenchmarkID: benchmarkID,
		Payload:     payload,
		UpdateAt:    time.Now().UTC().Format(time.RFC3339),
	}

	err := s.repo.Update(vc)
	return vc, err
}

func (s *Service) Delete(orgID, benchmarkID string, id int64) error {
	return s.repo.Delete(orgID, benchmarkID, id)
}
