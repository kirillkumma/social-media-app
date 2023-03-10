package app

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"strings"
	"time"
)

type Config struct {
	HTTP     HTTP     `mapstructure:"http" validate:"required"`
	Postgres Postgres `mapstructure:"postgres" validate:"required"`
}

type HTTP struct {
	Host string `mapstructure:"host" validate:"required"`
	Port string `mapstructure:"port" validate:"required"`
}

type Postgres struct {
	Conn                string        `mapstructure:"conn" validate:"required"`
	MaxOpenConn         int           `mapstructure:"max_open_conn" validate:"required"`
	MaxIdleConn         int           `mapstructure:"max_idle_conn" validate:"required"`
	ConnMaxLifetime     time.Duration `mapstructure:"conn_max_lifetime" validate:"required"`
	IdleConnMaxLifetime time.Duration `mapstructure:"idle_conn_max_lifetime" validate:"required"`
}

func LoadConfig() (*Config, error) {

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	cfg := &Config{
		HTTP: HTTP{
			Host: "localhost",
			Port: "8080",
		},
		Postgres: Postgres{
			Conn:                "postgresql://root:pass@localhost:5432/social-media?sslmode=disable",
			ConnMaxLifetime:     30,
			IdleConnMaxLifetime: 15,
			MaxOpenConn:         30,
			MaxIdleConn:         10,
		},
	}

	err := viper.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return nil, err
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	err = validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
