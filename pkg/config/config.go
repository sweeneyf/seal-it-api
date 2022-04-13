package config

type Configuration struct {
	Server       ServerConfiguration
	Database     DatabaseConfiguration
	Hedera       HederaConfiguration
	Log          LogConfiguration
	Cloud        CloudConfiguration
	EXAMPLE_PATH string
	EXAMPLE_VAR  string
}

// ServerConfigurations exported
type ServerConfiguration struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfiguration struct {
	DBName     string
	DBUser     string
	DBPassword string
}

// Log configuration exported
type LogConfiguration struct {
	Filename string
	Level    string
}

type CloudConfiguration struct {
	Region string
}

// HederaConfigurations exported
type HederaConfiguration struct {
	AccountId  string
	PrivateKey string
	TopicId    string
}
