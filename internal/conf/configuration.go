package conf

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type GlobalConfig struct {
	APPLICATION_NAME      string `mapstructure:"APPLICATION_NAME"`
	APPLICATION_ENV       string `mapstructure:"APPLICATION_ENV"`
	VERSION               string `mapstructure:"VERSION"`
	DB_HOST               string `mapstructure:"DB_HOST"`
	DB_USER               string `mapstructure:"DB_USER"`
	DB_PASSWORD           string `mapstructure:"DB_PASSWORD"`
	DB_NAME               string `mapstructure:"DB_NAME"`
	DB_PORT               string `mapstructure:"DB_PORT"`
	JWT_SECRET            string `mapstructure:"JWT_SECRET"`
	RequestIDHeader       string `mapstructure:"REQUEST_ID_HEADER"`
	STRIPE_SECRET_KEY     string `mapstructure:"STRIPE_SECRET_KEY"`
	SENDGRID_API_KEY      string `mapstructure:"SENDGRID_API_KEY"`
	SENDGRID_EMAIL_FROM   string `mapstructure:"SENDGRID_EMAIL_FROM"`
	BUCKET_NAME           string `mapstructure:"BUCKET_NAME"`
	KEY_ELATED            string `mapstructure:"KEY_ELATED"`
	IV_ELATED             string `mapstructure:"IV_ELATED"`
	STRIPE_WEBHOOK_SECRET string `mapstructure:"STRIPE_WEBHOOK_SECRET"`
}

var (
	logger = log.Default()
)

func InitEnv() GlobalConfig {
	// Load .env file; godotenv sets them in the process's environment
	err := godotenv.Load() // Loads from .env by default, specify a path as an argument if needed
	if err != nil {
		logger.Printf("Error loading .env file: %s", err)
	}

	// Initialize your config struct
	c := GlobalConfig{}

	// Tell Viper to automatically read from environment variables
	viper.AutomaticEnv()

	// Optionally, reflect over struct tags to bind specific environment variables if needed
	t := reflect.TypeOf(c)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("mapstructure")
		envVar := strings.ToUpper(strings.ReplaceAll(tag, "_", ""))
		viper.BindEnv(tag, envVar)
	}

	// Set defaults or read other configurations as needed
	viper.SetDefault("APPLICATION_NAME", "")
	viper.SetDefault("APPLICATION_ENV", "development")
	viper.SetDefault("VERSION", "1.0.0")
	viper.SetDefault("PORT", "8080")

	// Unmarshal environment configurations into your struct
	if err := viper.Unmarshal(&c); err != nil {
		log.Printf("Unable to decode into struct: %s", err)
	}

	// Debugging: print DB_HOST to verify
	fmt.Println("DB_HOST:", c.DB_HOST)

	return c

}
