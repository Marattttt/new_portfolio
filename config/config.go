package config

import (
	"context"
	"strings"

	"github.com/sethvargo/go-envconfig"
)

var Conf *AppConfig

func ReadConfig(ctx context.Context) error {
	var temp AppConfig
	if err := envconfig.Process(ctx, &temp); err != nil {
		return err
	}
	Conf = &temp
	return nil
}

type AppConfig struct {
	Mode   string       `env:"MODE, default=debug"`
	Server ServerConfig `env:", prefix=SERV_"`
}

func (ac AppConfig) IsDebug() bool {
	return strings.ToLower(ac.Mode) == "debug"
}

type ServerConfig struct {
	Host string `env:"HOST, default=0.0.0.0"`
	Port int    `env:"PORT, default=2024"`
}
