package config

import "time"

type (
	Config struct {
	}

	Server struct {
		Port           int           `mapstructure:"port" validate:"required"`
		AllowedOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit      string        `mapstructure:"bodyLimit" validate:"required"`
		TimeOut        time.Duration `mapstructure:"timeout" validate:"required"`
	}

	OAuth struct {
	}

	State struct {
	}

	Database struct {
	}
)
