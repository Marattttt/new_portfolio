package config

import (
	"context"
	"fmt"
	"time"

	"github.com/sethvargo/go-envconfig"
)

func NewApp(ctx context.Context) (*App, error) {
	var app App
	if err := envconfig.Process(ctx, &app); err != nil {
		return nil, fmt.Errorf("processing env: %w", err)
	}

	return &app, nil
}

type App struct {
	Runtime
	Server
}

type Runtime struct {
	RunTimeout time.Duration `env:"RUN_TIMEOUT, default=1m"`
}

type Server struct {
	Port int `env:"PORT, default=3002"`
}
