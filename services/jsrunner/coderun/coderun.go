package coderun

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Marattttt/new_portfoliio/services/jsrunner/internal/config"
)

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

	output += fmt.Sprintf("Exit Code: %d; ", r.ExitCode)

	if r.Err != nil {
		output += fmt.Sprintf("Error: %s", r.Err.Error())
	}

	return strings.TrimSpace(output)
}

type LocalRunner struct {
	Conf   *config.Runtime
	Logger *slog.Logger
}

// Retunred error is error with starting/finishing the exectution,
func (lr LocalRunner) Exec(ctx context.Context, file io.Reader) (RunResult, error) {
	code, err := io.ReadAll(file)
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

	return lr.Run(ctx, path)
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
func (lr LocalRunner) Run(ctx context.Context, absPath string) (RunResult, error) {
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
