package config

import (
	"context"
	"fmt"
	"path"
	"slices"
	"strings"

	"github.com/sethvargo/go-envconfig"
)

type App struct {
	Server
	Runners
}

type Server struct {
	Methods  []string `env:"SERVE_METHODS, default=http,grpc"`
	HttpPort int      `env:"HTTP_PORT, default=3000"`
	GrpcPort int      `env:"GRPC_PORT, default=3001"`
}

type Runners struct {
	SupportedLangs       []Lang
	SupportedLangStrings []string `env:"RUN_SUPPORTED"`
	// Must be full path, not relative
	GoRunDir *string `env:"RUN_GO_DIR, default=/tmp/go_runner/"`
	JsRunDir *string `env:"RUN_JS_DIR, default=/tmp/js_runner/"`
}

// Parses string versions of langs and fills the suppored langs list
func (r *Runners) parseLangs() error {
	if r.SupportedLangs == nil {
		r.SupportedLangs = make([]Lang, 0)
	}

	for _, str := range r.SupportedLangStrings {
		parsed := ParseLang(str)
		if parsed == nil {
			return fmt.Errorf("could not parse langauge: %s", str)
		}
		r.SupportedLangs = append(r.SupportedLangs, *parsed)
	}
	return nil
}

type Lang int

const (
	GoLang Lang = iota
	JavaScript
)

const UnsupportedLanguageMsg = "<unsupported>"

// The name is turned to lowercase and trimmed to check against common names of used languages
func ParseLang(name string) *Lang {
	formatted := strings.ToLower(name)
	formatted = strings.TrimSpace(formatted)

	// Reserve space for possible variants
	result := new(Lang)

	switch formatted {
	case "go", "golang":
		*result = GoLang
	case "js", "javascript":
		*result = JavaScript
	default:
		result = nil
	}

	return result
}

// Lowercase name fo the language
func (l Lang) String() string {
	switch l {
	case GoLang:
		return "golang"
	case JavaScript:
		return "javascript"
	}

	return UnsupportedLanguageMsg
}

func NewApp(ctx context.Context) (*App, error) {
	var app App
	if err := envconfig.Process(ctx, &app, toLowerMutator{}); err != nil {
		return nil, fmt.Errorf("processsing env: %w", err)
	}

	if err := app.Runners.parseLangs(); err != nil {
		return nil, err
	}

	if err := Validate(app); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	return &app, nil
}

// Check all fields and return the first error encountered
func Validate(conf App) error {
	if err := validateServingMethods(&conf.Server); err != nil {
		return err
	}
	if err := validateRunnerConfig(&conf.Runners); err != nil {
		return err
	}
	return nil
}

func validateServingMethods(conf *Server) error {
	supportedServers := []string{"http", "grpc"}

	for _, serve := range conf.Methods {
		if slices.Index(supportedServers, serve) == -1 {
			return fmt.Errorf("serve method not supported: %s", serve)
		}
	}

	if slices.Index(conf.Methods, "http") == -1 {
		conf.HttpPort = -1
	} else if conf.HttpPort == 0 {
		return fmt.Errorf("http method specified without port")
	}

	if slices.Index(conf.Methods, "grpc") == -1 {
		conf.GrpcPort = -1
	} else if conf.GrpcPort == 0 {
		return fmt.Errorf("http method specified without port")
	}
	return nil
}

// Ensure present absolute paths for needed languages
//
// Requires the languages to be parsed to check which language configs to check
func validateRunnerConfig(conf *Runners) error {
	// Ensure all needed paths are provided
	for _, l := range conf.SupportedLangs {
		switch l {
		case GoLang:
			if conf.GoRunDir != nil && !path.IsAbs(*conf.GoRunDir) {
				return fmt.Errorf("go path specified and not absolute")
			}
		case JavaScript:
			if conf.JsRunDir != nil && !path.IsAbs(*conf.JsRunDir) {
				return fmt.Errorf("javascript path specified and not absolute")
			}
		default:
			return fmt.Errorf("lanuage %s is not supported", l.String())
		}
	}
	return nil
}

// Wrapper for formatting all config values from using the envconfig package
type toLowerMutator struct{}

// Turns all string values to lowercase
func (_ toLowerMutator) EnvMutate(ctx context.Context, originalKey, resolvedKey, originalValue, currentValue string) (newValue string, stop bool, err error) {
	return strings.ToLower(currentValue), false, nil
}
