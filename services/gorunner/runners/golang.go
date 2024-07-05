package runners

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/Marattttt/newportfolio/services/gorunner/config"
)

// Result of running user-provided code
// If code could not be run, a separate error should be returned
type RunResult struct {
	// Code output
	Text string
	// Error with running code
	Err string
}

func InitGoEnv(conf *config.Runners) error {
	if err := os.MkdirAll(*conf.GoRunDir, 0777); err != nil {
		return fmt.Errorf("initialzing runtime directory: %w", err)
	}

	cmd := exec.Command("go", "mod", "init", "portfolio/gorunner")
	cmd.Dir = *conf.GoRunDir

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

// Writes the code to a temporary file and runs it
func (lg *LocalGoRunner) Run(ctx context.Context, code string) (RunResult, error) {
	tempFile, err := os.CreateTemp(*lg.conf.GoRunDir, "run_*.go")
	if err != nil {
		return RunResult{}, fmt.Errorf("creating temp input file: %w", err)
	}

	defer lg.clearAndLogTempFile(tempFile)

	_, err = tempFile.WriteString(code)
	if err != nil {
		return RunResult{}, fmt.Errorf("writing code to a temp file: %w", err)
	}

	lg.logger.Info("Created file", slog.String("path", tempFile.Name()))

	return lg.runFile(ctx, tempFile)
}

func (lg *LocalGoRunner) runFile(ctx context.Context, file *os.File) (RunResult, error) {
	cmd := exec.CommandContext(ctx, "go", "run", file.Name())
	cmd.Dir = *lg.conf.GoRunDir

	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return RunResult{}, fmt.Errorf("starting to run code: %w", err)
	}

	// Begin readign while the command runs
	out, errout, err := readAllOutput(&stdout, &stderr)
	if err != nil {
		return RunResult{}, fmt.Errorf("reading code run output: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return RunResult{}, fmt.Errorf("running compiled file: %w", err)
	}

	return RunResult{Text: string(out), Err: string(errout)}, nil
}

func (lg *LocalGoRunner) clearAndLogTempFile(openFile *os.File) {
	openFile.Close()
	name := openFile.Name()
	fname := slog.String("fname", openFile.Name())

	if err := os.Remove(name); err != nil {
		lg.logger.Warn("Error removing temp file", fname, slog.String("err", err.Error()))
	} else {
		lg.logger.Info("Removed temp file", fname)
	}

	return
}

func readAllOutput(out, errout io.Reader) ([]byte, []byte, error) {
	var (
		outData    = make([]byte, 0)
		erroutData = make([]byte, 0)
		errChan    = make(chan error, 2)
		wg         sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		buf := make([]byte, 1024)
		for {
			n, err := out.Read(buf)
			if err == io.EOF {
				return
			}
			if err != nil {
				errChan <- fmt.Errorf("reading out: %w", err)
			}

			outData = append(outData, buf[:n]...)
		}
	}()

	go func() {
		defer wg.Done()
		buf := make([]byte, 1024)
		for {
			n, err := errout.Read(buf)
			if err == io.EOF {
				return
			}
			if err != nil {
				errChan <- fmt.Errorf("reading errout: %w", err)
			}

			erroutData = append(erroutData, buf[:n]...)
		}
	}()
	wg.Add(1)
	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		errs := make([]error, len(errChan))
		for err := range errChan {
			errs = append(errs, err)
		}
		err := errors.Join(errs...)

		return nil, nil, err
	}

	return outData, erroutData, nil
}
