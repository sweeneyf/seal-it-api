package main

import (
	"github.com/spf13/viper"
	"github.com/sweeneyf/seal-it-api/app"
	"github.com/sweeneyf/seal-it-api/infra/logfactory"
	"github.com/sweeneyf/seal-it-api/pkg/config"
	"github.com/sweeneyf/seal-it-api/pkg/logger"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var config config.Configuration
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	logger.SetLogger(logfactory.CreateLogger(logger.GetLogLevel(config.Log.Level), config.Log.Filename))
	logger.Log.Info("Application started")

	app.StartApp(config)

}
