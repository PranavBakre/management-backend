package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	// DB connection string
	DBUri string `mapstructure:"DB_URI"`
}

var cfg *Config = nil

/*
GetConfig will set all config variables in cfg if not already set, and cfg
*/
func GetConfig() *Config {
	if cfg == nil {
		// Setup viper
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")

		// Set default values for config vars
		viper.SetDefault("DB_URI", "")

		// Automatically override values in config file with those in environment
		viper.AutomaticEnv()

		// Read config file
		err := viper.ReadInConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
			} else {
				// Config file was found but another error was produced
				log.Fatal(err)
			}
		}

		// Set config object
		err = viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatal(err)
		}
	}

	return cfg
}
