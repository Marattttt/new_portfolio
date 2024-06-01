package runners

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"

	"github.com/Marattttt/newportfolio/services/fastrunner/config"
)

// ONLY SUPPORTS LOCAL
func InitGoEnv(conf *config.Runners) error {
	if err := os.MkdirAll(conf.GoRunDir, 0777); err != nil {
		return fmt.Errorf("initialzing runtime directory: %w", err)
	}

	args := []string{"mod", "init", "portfolio/gorunner"}
	cmd := exec.Command("go", args...)
	cmd.Dir = conf.GoRunDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		if !strings.Contains(string(output), "already exists") {
			return fmt.Errorf("output: %s, executing command: %w", string(output), err)
		}
	}

	return nil
}

func NewGoRunner(logger *slog.Logger, conf *config.Runners) LocalGoRunner {
	return LocalGoRunner{
		logger: logger,
		conf:   conf,
	}
}

type LocalGoRunner struct {
	logger *slog.Logger
	conf   *config.Runners
}

func (lg *LocalGoRunner) Run(code string, name string) (string, error) {
	// FIXME: no delete after use
	inFile, err := os.CreateTemp(lg.conf.GoRunDir, name+"*.go")
	if err != nil {
		return "", fmt.Errorf("creating temp input file: %w", err)
	}
	defer inFile.Close()

	// FIXME:
	_, err = inFile.WriteString(code)
	if err != nil {
		return "", fmt.Errorf("writing code to a temp file: %w", err)
	}

	args := []string{"run", inFile.Name()}
	fmt.Println(inFile.Name())
	cmd := exec.Command("go", args...)
	cmd.Dir = lg.conf.GoRunDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("executing go code with %s, output: %s: %w", cmd.String(), string(output), err)
	}

	return string(output), nil
}
