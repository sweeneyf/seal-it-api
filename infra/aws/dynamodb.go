package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/sweeneyf/seal-it-api/entity"
	"github.com/sweeneyf/seal-it-api/pkg/config"
	"github.com/sweeneyf/seal-it-api/pkg/logger"
)

type DynamoDbClient struct {
	region string
}

func NewDynamoDbClient(region string) *DynamoDbClient {
	return &DynamoDbClient{
		region: region,
	}
}

func (h *DynamoDbClient) SaveDeed(config config.Configuration, deed entity.Deed) error {

	logger.Log.Info("Saving deed to dynamodb")
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(), func(o *awsconfig.LoadOptions) error {
		o.Region = config.Cloud.Region
		return nil
	})
	if err != nil {
		logger.Log.Error(err.Error(), err)
		return err
	}

	// Create an Amazon S3 service client
	svc := dynamodb.NewFromConfig(cfg)
	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("deed"),
		Item: map[string]types.AttributeValue{
			"id":   &types.AttributeValueMemberS{Value: uuid.New().String()},
			"deed": &types.AttributeValueMemberS{Value: "test"},
			"type": &types.AttributeValueMemberS{Value: "test"},
		},
	})
	if err != nil {
		logger.Log.Error(err.Error(), err)
		return err
	}
	logger.Log.Info("Deed saved")

	return nil
}
