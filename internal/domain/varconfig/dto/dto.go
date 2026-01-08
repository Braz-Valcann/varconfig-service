package dto

type CreateRequest struct {
	Payload map[string]interface{} `json:"payload" binding:"required"`
}

type UpdateRequest struct {
	Payload map[string]interface{} `json:"payload" binding:"required"`
}

// A ideia é padrozinar resposta
type Response struct {
	ID          int64                  `json: "id"`
	OrgID       string                 `json: "orgId"`
	BenchmarkID string                 `json: "benchmark_id"`
	Payload     map[string]interface{} `json: "payload"`
	CreatedAt   string                 `json: "createdAt"`
	UpdateAt    string                 `json: "updateAt"`
}
