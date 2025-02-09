package utils

import "github.com/spf13/viper"

// Config stores all configs of the application
// The values are read by viper from a config file or environment variables
type Config struct {
	DbSource          string `mapstructure:"DB_SOURCE"`
	HttpServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GrpcServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
}

// LoadConfig reads configs from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
