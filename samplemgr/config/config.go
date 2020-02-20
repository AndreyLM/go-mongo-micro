package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	appName string
	appPort int
)

// Load - loads configuration
func Load() {
	viper.SetDefault("APP_NAME", "app")
	viper.SetDefault("APP_PORT", "8002")

	viper.SetConfigName("application.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.ReadInConfig()
	// viper.AutomaticEnv()
}

// AppName - application name
func AppName() string {
	if appName == "" {
		appName = ReadEnvString("APP_NAME")
	}

	return appName
}

// ServicePort - service port
func ServicePort() string {
	return ReadEnvString("APP_PORT")
}

// ReadEnvString - reads variable from environment
func ReadEnvString(key string) string {
	checkIfSet(key)
	return viper.GetString(key)
}

// ReadEnvInt - reads int variable from environment
func ReadEnvInt(key string) int {
	checkIfSet(key)
	return viper.GetInt(key)
}

func checkIfSet(key string) {
	if !viper.IsSet(key) {
		err := fmt.Errorf("Key %s is not set", key)
		panic(err)
	}
}
