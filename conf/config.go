package conf

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	OpenId   OpenIDCredentials
	Sms      SmsConfigurations
	Postgres PostgresConfigurations
}

type OpenIDCredentials struct {
	ClientId     string
	ClientSecret string
}

type SmsConfigurations struct {
	AccountSid   string
	ClientSecret string
}

type PostgresConfigurations struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbName   string
	PgDriver           string
	PostgresqlSslMode  string
}

func RequireConfigurations(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file does not exist")
		}
		return nil, err
	}
	return v, nil
}

func ParseConfigurations(v *viper.Viper) (*Config, error) {
	var config Config
	err := v.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &config, nil
}
