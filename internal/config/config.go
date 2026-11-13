package config

import (
	"fmt"
	"os"
)

type Config struct {
	Addr     string
	LogLevel string
}

func Load() (Config, error) {
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	level := os.Getenv("APP_LOG_LEVEL")
	if level == "" {
		level = "INFO"
	}
	if addr[0] != ':' && len(addr) < 3 {
		return Config{}, fmt.Errorf("invalid APP_ADDR %q", addr)
	}
	return Config{Addr: addr, LogLevel: level}, nil
}
