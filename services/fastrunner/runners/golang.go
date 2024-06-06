package runners

import (
	"context"
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

// Runresult contains errors with the code, error returned along with it contains a system error
// FIXME: no use of context parameter
func (lg *LocalGoRunner) Run(ctx context.Context, code string, id string) (RunResult, error) {
	// FIXME: no delete after use
	inFile, err := os.CreateTemp(lg.conf.GoRunDir, id+"*.go")
	if err != nil {
		return RunResult{}, fmt.Errorf("creating temp input file: %w", err)
	}
	defer lg.clearFile(inFile)

	// TODO: Add maxfile check?
	_, err = inFile.WriteString(code)
	if err != nil {
		return RunResult{}, fmt.Errorf("writing code to a temp file: %w", err)
	}

	args := []string{"run", inFile.Name()}
	lg.logger.Info("Created file", slog.String("path", inFile.Name()))
	cmd := exec.Command("go", args...)
	cmd.Dir = lg.conf.GoRunDir

	output, err := cmd.CombinedOutput()

	fmt.Println(string(output))

	// Remove comment about command line arguments used for executing
	stripped := strings.Split(string(output), "\n")[1]
	return RunResult{stripped, err}, nil
}

func (lg *LocalGoRunner) clearFile(openFile *os.File) error {
	openFile.Close()
	name := openFile.Name()
	if err := os.Remove(name); err != nil {
		return err
	}
	return nil
}
