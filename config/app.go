package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Application struct {
	Port      string
	FilePath  string
	Region    string
	Queue     string
	NRLicense string
}

func Load() (*Application, error) {
	if os.Getenv("APP_ENV") == "local" {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	viper.AutomaticEnv()

	app := &Application{
		Port:      viper.GetString("APP_PORT"),
		FilePath:  viper.GetString("OUTPUT_FILE_PATH"),
		Region:    viper.GetString("AWS_REGION"),
		Queue:     viper.GetString("AWS_SQS_QUEUE_URL"),
		NRLicense: viper.GetString("NEW_RELIC_LICENSE_KEY"),
	}

	return app, nil
}
