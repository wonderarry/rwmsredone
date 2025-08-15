package config

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv" // go get github.com/joho/godotenv
)

const EnvModeVar = "RWMS_ENV"

type Config struct {
	HTTP struct {
		Addr         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		IdleTimeout  time.Duration
	}
	DB struct {
		DSN string
	}
	Security struct {
		JWTSecret string
	}
	TemplatesDir string
	LogLevel     string
}

func MustLoad() Config {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}
	return cfg
}

func Load() (Config, error) {
	// If RWMS_ENV is not set, assume local dev and hydrate env from files.
	if os.Getenv(EnvModeVar) == "" {
		_ = godotenv.Load("local.env")
		_ = godotenv.Load("local.secrets")
	}

	var missing []string
	get := func(key string) string {
		if v, ok := os.LookupEnv(key); ok && strings.TrimSpace(v) != "" {
			return v
		}
		missing = append(missing, key)
		return ""
	}
	parseDur := func(key string) time.Duration {
		v := get(key)
		if v == "" {
			return 0
		}
		d, err := time.ParseDuration(v)
		if err != nil {
			missing = append(missing, key+"(invalid duration: "+err.Error()+")")
			return 0
		}
		return d
	}

	var cfg Config
	// HTTP
	cfg.HTTP.Addr = get("RWMS_HTTP__ADDR")
	cfg.HTTP.ReadTimeout = parseDur("RWMS_HTTP__READ_TIMEOUT")
	cfg.HTTP.WriteTimeout = parseDur("RWMS_HTTP__WRITE_TIMEOUT")
	cfg.HTTP.IdleTimeout = parseDur("RWMS_HTTP__IDLE_TIMEOUT")

	// DB & Security
	cfg.DB.DSN = get("RWMS_DB__DSN")
	cfg.Security.JWTSecret = get("RWMS_SECURITY__JWT_SECRET")

	// Misc
	cfg.TemplatesDir = get("RWMS_TEMPLATES_DIR")
	cfg.LogLevel = get("RWMS_LOG_LEVEL")

	if len(missing) > 0 {
		return Config{}, errors.New("missing/invalid required env vars: " + strings.Join(missing, ", "))
	}
	return cfg, nil
}
