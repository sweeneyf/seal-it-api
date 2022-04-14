package hedera

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/sweeneyf/seal-it-api/entity"
	"github.com/sweeneyf/seal-it-api/pkg/config"
	"github.com/sweeneyf/seal-it-api/pkg/logger"
)

type HederaClient struct{}

func NewHederaClient() *HederaClient {
	return &HederaClient{}
}

func (h *HederaClient) SealDeed(config config.Configuration, deed entity.Deed) (entity.Deed, error) {

	logger.Log.Info("In Record deed")
	accountId, err := hedera.AccountIDFromString(config.Hedera.AccountId)
	if err != nil {
		return deed, err
	}
	logger.Log.Info(accountId.String())
	privateKey, err := hedera.PrivateKeyFromString(config.Hedera.PrivateKey)
	if err != nil {
		logger.Log.Error(err.Error(), err)
		return deed, err
	}
	logger.Log.Info(privateKey.String())

	topicId, err := hedera.TopicIDFromString(config.Hedera.TopicId)
	if err != nil {
		return deed, err
	}
	logger.Log.Info(topicId.String())
	client := hedera.ClientForTestnet()
	client.SetOperator(accountId, privateKey)

	//Send "Hello, HCS!" to the topic
	submitMessage, err := hedera.NewTopicMessageSubmitTransaction().SetMessage([]byte("Hello, Test")).
		SetTopicID(topicId).
		Execute(client)

	if err != nil {
		println(err.Error(), ": error submitting to topic")
		return deed, err
	}

	//Get the receipt of the transaction
	receipt, err := submitMessage.GetReceipt(client)

	//Get the transaction status
	transactionStatus := receipt.Status
	logger.Log.Info("The message transaction status " + transactionStatus.String())

	return deed, nil
}
