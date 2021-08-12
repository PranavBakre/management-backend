package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	// DB connection string
	DBUri string `mapstructure:"DB_URI"`

	// Secret key for JWT
	JwtSecret string `mapstructure:"JWT_SECRET"`
}

// Unexported variable to implement singleton pattern
var cfg *Config = nil

/*
Init will read all config variables from the .env and environment variables
*/
func Init() {
	// Setup viper
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	// Set default values for config vars
	viper.SetDefault("DB_URI", "")
	viper.SetDefault("JWT_SECRET", "")

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

/*
Get will return the config object set in Init
*/
func Get() *Config {
	return cfg
}
