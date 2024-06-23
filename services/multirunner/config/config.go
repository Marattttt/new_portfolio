package config

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/sethvargo/go-envconfig"
)

func NewApp(ctx context.Context) (*App, error) {
	var app App
	if err := envconfig.Process(ctx, &app, toLowerMutator{}); err != nil {
		return nil, fmt.Errorf("processsing env: %w", err)
	}

	if err := Validate(app); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	return &app, nil
}

func Validate(conf App) error {
	supportedServers := []string{"http", "grpc"}

	for _, serve := range conf.Supported {
		if slices.Index(supportedServers, serve) == -1 {
			return fmt.Errorf("serve method not supported: %s", serve)
		}
	}

	if slices.Index(conf.Supported, "http") == -1 {
		conf.HttpPort = -1
	} else if conf.HttpPort == 0 {
		return fmt.Errorf("http method specified without port")
	}

	if slices.Index(conf.Supported, "grpc") == -1 {
		conf.GrpcPort = -1
	} else if conf.GrpcPort == 0 {
		return fmt.Errorf("http method specified without port")
	}

	return nil
}

type App struct {
	Server
	Runners
}

type Server struct {
	Supported []string `env:"SERVE, default=http,grpc"`
	HttpPort  int      `env:"HTTP_PORT, default=3000"`
	GrpcPort  int      `env:"GRPC_PORT, default=3001"`
}

type Runners struct {
	// Must be full path, not relative
	GoRunDir string `env:"RUN_GO_DIR, default=/tmp/gorunner/"`
}

// Wrapper for formatting all config values from using the envconfig package
type toLowerMutator struct{}

// Turns all string values to lowercase
func (_ toLowerMutator) EnvMutate(ctx context.Context, originalKey, resolvedKey, originalValue, currentValue string) (newValue string, stop bool, err error) {
	return strings.ToLower(currentValue), false, nil
}

type configValueValidator struct{}