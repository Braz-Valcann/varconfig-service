package benchmark

import (
	"context"
	"errors"
)

type PathParameter struct {
	ID          int64                  `json:"id"`
	OrgID       string                 `json:"orgId"`
	BenchmarkID string                 `json:"benchmark_id"`
	Payload     map[string]interface{} `json:"payload"`
	CreatedAt   string                 `json:"createdAt"`
	UpdateAt    string                 `json:"updateAt"`
	Context     context.Context        `json:"-"`
}

func (pp *PathParameter) Validate() error {
	if pp.BenchmarkID == "" {
		return errors.New("benchmarkId cannot be empty")
	}
	return nil
}
