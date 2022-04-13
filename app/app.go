package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sweeneyf/seal-it-api/app/handler"
	"github.com/sweeneyf/seal-it-api/app/middleware"
	"github.com/sweeneyf/seal-it-api/infra/aws"
	"github.com/sweeneyf/seal-it-api/infra/hedera"
	"github.com/sweeneyf/seal-it-api/pkg/config"
	"github.com/sweeneyf/seal-it-api/usecase/deed"
	"github.com/sweeneyf/seal-it-api/usecase/ledger"
)

var (
	router      *gin.Engine
	ds1         deed.Service
	deedHandler *handler.DeedHandler
)

func StartApp(config config.Configuration) {

	router = gin.Default()
	router.Use(middleware.CORS())
	router.Use(middleware.AddConfig(config))

	hederaClient := hedera.NewHederaClient()
	ledgerService := ledger.NewService(hederaClient)
	dynamoDbClient := aws.NewDynamoDbClient(config.Cloud.Region)
	deedService := deed.NewService(dynamoDbClient)
	deedHandler = handler.NewDeedHandler(ledgerService, deedService, config)

	mapUrls(router)
	if err := router.Run(fmt.Sprintf(":%d", config.Server.Port)); err != nil {
		panic(err)
	}
}

func mapUrls(router *gin.Engine) {
	apiV1 := router.Group("/v1")
	apiV1.GET("/ping", handler.Ping)
	apiV1.POST("/deed", deedHandler.SealAndSaveDeed)
}
