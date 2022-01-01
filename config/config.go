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

	//Google Client Id
	ClientId string `mapstructure:"CLIENT_ID"`

	//Google Client Secret
	ClientSecret string `mapstructure:"CLIENT_SECRET"`

	//Redirect Uri
	RedirectUri string `mapstructure:"REDIRECT_URI"`

	SuperUserGoogleId string `mapstructure:"SUPERUSER_GOOGLE_ID"`

	SuperUserName string `mapstructure:"SUPERUSER_NAME"`

	SuperUserEmail string `mapstructure:"SUPERUSER_EMAIL"`
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
	viper.SetDefault("CLIENT_ID", "")
	viper.SetDefault("CLIENT_SECRET", "")
	viper.SetDefault("REDIRECT_URI", "")
	viper.SetDefault("SUPERUSER_GOOGLE_ID", "")
	viper.SetDefault("SUPERUSER_NAME", "")
	viper.SetDefault("SUPERUSER_EMAIL", "")
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
