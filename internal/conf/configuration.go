package conf

import (
	"log"

	"github.com/spf13/viper"
)

type GlobalConfig struct {
	APPLICATION_NAME    string `mapstructure:"APPLICATION_NAME"`
	APPLICATION_ENV     string `mapstructure:"APPLICATION_ENV"`
	VERSION             string `mapstructure:"VERSION"`
	DB_HOST             string `mapstructure:"DB_HOST"`
	DB_USER             string `mapstructure:"DB_USER"`
	DB_PASSWORD         string `mapstructure:"DB_PASSWORD"`
	DB_NAME             string `mapstructure:"DB_NAME"`
	DB_PORT             string `mapstructure:"DB_PORT"`
	JWT_SECRET          string `mapstructure:"JWT_SECRET"`
	RequestIDHeader     string `mapstructure:"REQUEST_ID_HEADER"`
	STRIPE_SECRET_KEY   string `mapstructure:"STRIPE_SECRET_KEY"`
	SENDGRID_API_KEY    string `mapstructure:"SENDGRID_API_KEY"`
	SENDGRID_EMAIL_FROM string `mapstructure:"SENDGRID_EMAIL_FROM"`
	BUCKET_NAME         string `mapstructure:"BUCKET_NAME"`
	KEY_ELATED          string `mapstructure:"KEY_ELATED"`
	IV_ELATED           string `mapstructure:"IV_ELATED"`
}

var (
	logger = log.Default()
)

func InitEnv() GlobalConfig {
	c := GlobalConfig{}
	vi := viper.New()
	vi.SetConfigType("env")
	vi.SetConfigFile(".env")
	vi.AddConfigPath(".")
	vi.AutomaticEnv()

	vi.SetDefault("APPLICATION_NAME", "Zentive-Resource-Server")
	vi.SetDefault("APPLICATION_ENV", "development")
	vi.SetDefault("VERSION", "1.0.0")

	err := vi.ReadInConfig()

	if err != nil {
		logger.Fatalf("Error while reading config file %s", err)
	}

	err = vi.Unmarshal(&c)

	if err != nil {
		logger.Fatalf("Unable to decode into struct %s", err)
	}

	return c
}
