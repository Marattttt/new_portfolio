package config

type App struct {
	Runners
}

type Runners struct {
	// Must be full path, not relative
	GoRunDir string `env:"RUN_GO_DIR, default=/tmp/gorunner/"`
}
