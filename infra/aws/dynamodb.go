package aws

import (
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
	return nil
}
