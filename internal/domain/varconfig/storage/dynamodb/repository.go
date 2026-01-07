package dynamodbConfig

import (
	"context"
	"fmt"

	"github.com/Braz-Valcann/varconfig-service/internal/domain/varconfig"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Repository struct {
	client    *dynamodb.Client
	tableName string
}

// confirma que tu completa
var _ varconfig.Repository = (*Repository)(nil)

func New(client *dynamodb.Client, table string) *Repository {
	return &Repository{
		client:    client,
		tableName: table,
	}
}

func pk(orgID, benchmarkID string) string {
	return fmt.Sprintf("ORG#%s#BENCH#%s", orgID, benchmarkID)
}

func sk(id int64) string {
	return fmt.Sprintf("VAR#%d", id)
}

func (r *Repository) Create(ctx context.Context, varConfig *varconfig.VarConfig) error {
	item := map[string]interface{}{
		"PK":          pk(varConfig.OrgID, varConfig.BenchmarkID),
		"SK":          sk(varConfig.ID),
		"id":          varConfig.ID,
		"orgId":       varConfig.OrgID,
		"benchmarkId": varConfig.BenchmarkID,
		"payload":     varConfig.Payload,
		"createdAt":   varConfig.CreatedAt,
		"updatedAt":   varConfig.UpdateAt,
	}

	av, _ := attributevalue.MarshalMap(item)

	_, err := r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	})

	return err
}

func (r *Repository) Get(ctx context.Context, orgID, benchmarkID string, id int64) (*varconfig.VarConfig, error) {

	key, err := attributevalue.MarshalMap(map[string]string{
		"PK": pk(orgID, benchmarkID),
		"SK": sk(id),
	})

	if err != nil {
		return nil, err
	}

	out, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &r.tableName,
		Key:       key,
	})

	if err != nil || out.Item == nil {
		return nil, err
	}

	var varConfig varconfig.VarConfig
	_ = attributevalue.UnmarshalMap(out.Item, &varConfig)

	return &varConfig, nil
}

func (r *Repository) List(ctx context.Context, orgID, benchmarkID string) ([]*varconfig.VarConfig, error) {
	out, err := r.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              &r.tableName,
		KeyConditionExpression: aws.String("PK = :pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: pk(orgID, benchmarkID)},
		},
	})

	if err != nil {
		return nil, err
	}

	var varConfigs []*varconfig.VarConfig
	_ = attributevalue.UnmarshalListOfMaps(out.Items, &varConfigs)

	return varConfigs, nil
}

func (r *Repository) Update(ctx context.Context, varConfig *varconfig.VarConfig) error {
	item := map[string]interface{}{
		"PK":          pk(varConfig.OrgID, varConfig.BenchmarkID),
		"SK":          sk(varConfig.ID),
		"id":          varConfig.ID,
		"orgId":       varConfig.OrgID,
		"benchmarkId": varConfig.BenchmarkID,
		"payload":     varConfig.Payload,
		"updatedAt":   varConfig.UpdateAt,
	}

	av, _ := attributevalue.MarshalMap(item)

	_, err := r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	})

	return err
}

func (r *Repository) Delete(ctx context.Context, orgID, benchmarkID string, id int64) error {
	key, _ := attributevalue.MarshalMap(map[string]string{
		"PK": pk(orgID, benchmarkID),
		"SK": sk(id),
	})

	_, err := r.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &r.tableName,
		Key:       key,
	})

	return err
}
