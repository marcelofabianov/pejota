package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel  string `validate:"required,string,enum=debug|info|warn|error|fatal|panic"`
	LogFormat string `validate:"required,string,enum=text|json|logfmt"`
	LogOutput string `validate:"required,string,enum=stdout|stderr"`
	DbHost    string `validate:"required,string"`
	DbPort    int    `validate:"required,int"`
	DbUser    string `validate:"required,string"`
	DbPass    string `validate:"required,string"`
	DbName    string `validate:"required,string"`
}

func Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("./bootstrap")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.AutomaticEnv()

	// Log
	viper.BindEnv("log.level", "LOG_LEVEL")
	viper.BindEnv("log.format", "LOG_FORMAT")
	viper.BindEnv("log.output", "LOG_OUTPUT")

	// Database
	viper.BindEnv("db.host", "DB_HOST")
	viper.BindEnv("db.port", "DB_PORT")
	viper.BindEnv("db.user", "DB_USER")
	viper.BindEnv("db.password", "DB_PASSWORD")
	viper.BindEnv("db.name", "DB_NAME")

	return nil
}

func NewConfig() *Config {
	return &Config{
		LogLevel:  viper.GetString("log.level"),
		LogFormat: viper.GetString("log.format"),
		LogOutput: viper.GetString("log.output"),
		DbHost:    viper.GetString("db.host"),
		DbPort:    viper.GetInt("db.port"),
		DbUser:    viper.GetString("db.user"),
		DbPass:    viper.GetString("db.password"),
		DbName:    viper.GetString("db.name"),
	}
}
