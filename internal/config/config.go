package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	portDefault = "8088"
	envDefault  = EnvLocal
)

// Reading config
// https://github.com/ilyakaznacheev/cleanenv
type Config struct {
	Env         Env    `yaml:"env" env:"ENV" env-default:"local" env-required:"true"`
	PortDefault string `yaml:"port_default" env-required:"true"`
}

type StartValues struct {
	ConfigPath string //might be extracted later// do not need to keep it after setup is done
	Env        Env
	AppPort    string
}

type Env string

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

// Env Variables reading
func readEnvValues() (StartValues, error) {
	//todo secrets db etc read here
	if err := godotenv.Load(); err != nil {
		return StartValues{}, fmt.Errorf("Can not load godotenv %w", err)
	}
	return StartValues{
		ConfigPath: os.Getenv("CONFIG_PATH"),
	}, nil
}

// Config values reading..
func MustLoad() StartValues {
	startValues, err := readEnvValues()
	if err != nil {
		log.Fatalf("can not load initial env vars: %v", err)
	}
	if startValues.ConfigPath == "" {
		log.Fatal("CONFIG_PATH env var is not set")
	}
	if _, err := os.Stat(startValues.ConfigPath); err != nil {
		log.Fatalf("Config file does not exist: : %s", startValues.ConfigPath)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(startValues.ConfigPath, &cfg); err != nil {
		log.Fatalf("Can not read config file: : %s: %v", startValues.ConfigPath, err)
	}

	if cfg.PortDefault != "" {
		startValues.AppPort = cfg.PortDefault
	} else {
		startValues.AppPort = portDefault
	}
	if cfg.Env != "" {
		startValues.Env = cfg.Env
	} else {
		startValues.Env = envDefault
	}
	return startValues
}
