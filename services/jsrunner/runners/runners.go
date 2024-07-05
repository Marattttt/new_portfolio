package runners

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"sync"

	"github.com/Marattttt/new_portfoliio/services/jsrunner/config"
)

type RunRequest struct {
	Code string
}

type RunResult struct {
	Output   string
	ExitCode int
	Err      string
}

// Formats the result in a human-readable format
func (r RunResult) String() string {
	output := ""
	if r.Output != "" {
		output += fmt.Sprintf("Output: `%s`; ", string(r.Output))
	}

	exitcode := fmt.Sprintf("Exit Code: %d; ", r.ExitCode)

	errout := ""
	if r.Err != "" {
		errout = fmt.Sprintf("Error: %s", string(r.Err))
	}

	return fmt.Sprintf("%s; %s; %s", output, exitcode, errout)
}

type NodeRunner interface {
	NodeRun(context.Context, string) (io.Reader, io.Reader, error)
}

type LocalRunner struct {
	Conf   *config.Runtime
	Logger *slog.Logger
	node   NodeRunner
}

func NewLocal(logger *slog.Logger, conf *config.Runtime, node NodeRunner) LocalRunner {
	return LocalRunner{
		Conf:   conf,
		Logger: logger,
		node:   node,
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

	out, errout, err := lr.node.NodeRun(ctx, string(code))
	if err != nil {
		return RunResult{}, fmt.Errorf("node: %w", err)
	}

	outData, errData, err := readAllOutput(out, errout)
	if err != nil {
		return RunResult{}, fmt.Errorf("running node: %w", err)
	}

	return RunResult{
		Output: string(outData),
		Err:    string(errData),
	}, nil
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
