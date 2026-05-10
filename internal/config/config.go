package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EnvFileLoaded bool
	DBPath        string
	ListenAddr    string
}

type AuthSecrets struct {
	AdminPassword string
	SessionKey    string
}

func Load() (Config, error) {
	cfg := Config{DBPath: "portfolio.db?_foreign_keys=on", ListenAddr: ":8080"}

	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(".env"); err != nil {
			return Config{}, err
		}
		cfg.EnvFileLoaded = true
	}

	if v := os.Getenv("PORTFOLIO_DB"); v != "" {
		cfg.DBPath = v
	}
	if p := os.Getenv("PORT"); p != "" {
		cfg.ListenAddr = ":" + p
	}
	return cfg, nil
}

func LoadAuthSecrets() AuthSecrets {
	return AuthSecrets{
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),
		SessionKey:    os.Getenv("ADMIN_SESSION_KEY"),
	}
}
