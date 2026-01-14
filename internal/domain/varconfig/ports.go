package varconfig

type VarConfig struct {
	ID          int64                  `json: "id"`
	OrgID       string                 `json: "orgId"`
	BenchmarkID string                 `json: "benchmark_id"`
	Payload     map[string]interface{} `json: "payload"`
	CreatedAt   string                 `json: "createdAt"`
	UpdateAt    string                 `json: "updateAt"`
}

type IRepository interface {
	Create(varConfig *VarConfig) error
	Get(orgID, benchmarkID string, id int64) (*VarConfig, error)
	List(orgID, benchmarkID string) ([]*VarConfig, error)
	Update(varConfig *VarConfig) error
	Delete(orgID, benchmarkID string, id int64) error
}

type IService interface {
	Create(orgID, benchmarkID string, payload map[string]interface{}) (*VarConfig, error)
	Get(orgID, benchmarkID string, id int64) (*VarConfig, error)
	List(orgID, benchmarkID string) ([]*VarConfig, error)
	Update(orgID, benchmarkID string, id int64, payload map[string]interface{}) (*VarConfig, error)
	Delete(orgID, benchmarkID string, id int64) error
}
