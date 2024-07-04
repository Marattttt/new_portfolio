package runners

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Marattttt/new_portfoliio/services/jsrunner/config"
)

type RunRequest struct {
	Code string
}

type RunResult struct {
	Output   []byte
	ExitCode int
	Err      error
}

// Formats the result in a human-readable format
func (r RunResult) String() string {
	output := ""
	if r.Output != nil {
		output += fmt.Sprintf("Output: `%s`; ", string(r.Output))
	}

	exitcode := fmt.Sprintf("Exit Code: %d; ", r.ExitCode)

	errout := ""
	if r.Err != nil {
		errout = fmt.Sprintf("Error: %s", r.Err.Error())
	}

	return fmt.Sprintf("%s; %s; %s", output, exitcode, errout)
}

type LocalRunner struct {
	Conf   *config.Runtime
	Logger *slog.Logger
}

func NewLocal(logger *slog.Logger, conf *config.Runtime) LocalRunner {
	return LocalRunner{
		Conf:   conf,
		Logger: logger,
	}
}

func (lr LocalRunner) RunJs(ctx context.Context, req RunRequest) (RunResult, error) {
	r := bytes.NewReader([]byte(req.Code))
	return lr.RunJsReader(ctx, r)
}

// Retunred error is error with starting/finishing the exectution,
func (lr LocalRunner) RunJsReader(ctx context.Context, r io.Reader) (RunResult, error) {
	code, err := io.ReadAll(r)
	if err != nil {
		return RunResult{}, fmt.Errorf("reading code: %w", err)
	}

	lr.Logger.Debug("Received code", slog.String("code", string(code)))

	codefile, path, err := createFile()
	if err != nil {
		return RunResult{}, fmt.Errorf("creting temp file: %w", err)
	}
	// TODO: add config for when to remove files or do a clean-up
	// defer removeFile(codefile)

	lr.Logger.Debug("Created temp file", slog.String("path", path))

	_, err = codefile.Write(code)
	if err != nil {
		return RunResult{}, fmt.Errorf("writing to temp file: %w", err)
	}

	codefile.Close()

	return lr.runFile(ctx, path)
}

func createFile() (*os.File, string, error) {
	tmp, err := os.CreateTemp("", "")
	if err != nil {
		return nil, "", fmt.Errorf("creating temp file: %w", err)
	}

	fullPath, err := filepath.Abs(tmp.Name())
	if err != nil {
		panic(err)
	}

	return tmp, fullPath, nil
}

func removeFile(file *os.File) {
	file.Close()
	fullPath, err := filepath.Abs(file.Name())
	if err != nil {
		return
	}
	os.Remove(fullPath)
}

// Execute js file locaated at a given location
func (lr LocalRunner) runFile(ctx context.Context, absPath string) (RunResult, error) {
	runCtx, cancel := context.WithTimeout(ctx, lr.Conf.RunTimeout)
	defer cancel()

	cmd := exec.CommandContext(runCtx, "node", absPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return RunResult{}, fmt.Errorf("starting node: %w", err)
	}

	return RunResult{
		Output:   output,
		ExitCode: -1,
	}, nil
}
