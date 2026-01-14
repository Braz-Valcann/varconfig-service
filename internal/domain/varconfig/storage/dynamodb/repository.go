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
	context   context.Context
}

// Assertion
var _ varconfig.IRepository = (*Repository)(nil)

func New(client *dynamodb.Client, table string) *Repository {
	return &Repository{
		client:    client,
		tableName: table,
		context:   context.TODO(),
	}
}

func pk(orgID, benchmarkID string) string {
	return fmt.Sprintf("ORG#%s#BENCH#%s", orgID, benchmarkID)
}

func sk(id int64) string {
	return fmt.Sprintf("VAR#%d", id)
}

func (r *Repository) Create(varConfig *varconfig.VarConfig) error {
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

	_, err := r.client.PutItem(r.context, &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	})

	return err
}

func (r *Repository) Get(orgID, benchmarkID string, id int64) (*varconfig.VarConfig, error) {

	key, err := attributevalue.MarshalMap(map[string]string{
		"PK": pk(orgID, benchmarkID),
		"SK": sk(id),
	})

	if err != nil {
		return nil, err
	}

	out, err := r.client.GetItem(r.context, &dynamodb.GetItemInput{
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

func (r *Repository) List(orgID, benchmarkID string) ([]*varconfig.VarConfig, error) {
	out, err := r.client.Query(r.context, &dynamodb.QueryInput{
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

func (r *Repository) Update(varConfig *varconfig.VarConfig) error {
	item := map[string]interface{}{
		"PK":          pk(varConfig.OrgID, varConfig.BenchmarkID),
		"SK":          sk(varConfig.ID),
		"id":          varConfig.ID,
		"orgId":       varConfig.OrgID,
		"benchmarkId": varConfig.BenchmarkID,
		"payload":     varConfig.Payload,
		"createdAt":   varConfig.CreatedAt,
		"updateAt":    varConfig.UpdateAt,
	}

	av, _ := attributevalue.MarshalMap(item)

	_, err := r.client.PutItem(r.context, &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	})

	return err
}

func (r *Repository) Delete(orgID, benchmarkID string, id int64) error {
	key, _ := attributevalue.MarshalMap(map[string]string{
		"PK": pk(orgID, benchmarkID),
		"SK": sk(id),
	})

	_, err := r.client.DeleteItem(r.context, &dynamodb.DeleteItemInput{
		TableName: &r.tableName,
		Key:       key,
	})

	return err
}
