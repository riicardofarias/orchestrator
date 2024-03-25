package mongodb

import "github.com/spf13/viper"

type Config struct {
	Database string
	Host     string
	Port     string
}

func GetDatabaseConfig() *Config {
	return &Config{
		Database: viper.GetString("database.name"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
	}
}
