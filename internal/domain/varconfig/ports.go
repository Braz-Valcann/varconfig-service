package varconfig

import "context"

type VarConfig struct {
	ID          int64                  `json: "id"`
	OrgID       string                 `json: "orgId"`
	BenchmarkID string                 `json: "benchmark_id"`
	Payload     map[string]interface{} `json: "payload"`
	CreatedAt   string                 `json: "createdAt"`
	UpdateAt    string                 `json: "updateAt"`
}

type Repository interface {
	Create(ctx context.Context, varConfig *VarConfig) error
	Get(ctx context.Context, orgID, benchmarkID string, id int64) (*VarConfig, error)
	List(ctx context.Context, orgID, benchmarkID string) ([]*VarConfig, error)
	Update(ctx context.Context, varConfig *VarConfig) error
	Delete(ctx context.Context, orgID, benchmarkID string, id int64) error
}
