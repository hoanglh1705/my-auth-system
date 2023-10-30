package config

import (
	"fmt"
	"os"
)

type Configuration struct {
	Stage        string `env:"UP_STAGE"`
	Region       string `env:"REGION"`
	Host         string `env:"HOST"`
	Port         int    `env:"PORT"`
	ReadTimeout  int
	WriteTimeout int
	Debug        bool
	AllowOrigins []string

	JwtAdminSecret    string `env:"JWT_ADMIN_SECRET"`
	JwtAdminDuration  int    `env:"JWT_ADMIN_DURATION"`
	JwtAdminAlgorithm string `env:"JWT_ADMIN_ALGORITHM"`

	JwtCustomerSecret    string `env:"JWT_CUSTOMER_SECRET"`
	JwtCustomerDuration  int    `env:"JWT_CUSTOMER_DURATION"`
	JwtCustomerAlgorithm string `env:"JWT_CUSTOMER_ALGORITHM"`
}

// Load returns Configuration struct
func Load() (*Configuration, error) {
	appName := os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
	if configName := os.Getenv("CONFIG_NAME"); configName != "" {
		appName = configName
	}

	stage := os.Getenv("UP_STAGE")

	cfg := new(Configuration)
	if err := LoadWithAPS(cfg, appName, stage); err != nil {
		return nil, fmt.Errorf("error parsing environment config: %s", err)
	}

	return cfg, nil
}
