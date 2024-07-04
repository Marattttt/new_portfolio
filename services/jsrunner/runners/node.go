package runners

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
)

type LocalNode struct{}

func (ln LocalNode) NodeRun(ctx context.Context, code []byte) (io.ReadCloser, io.ReadCloser, error) {
	cmd := exec.CommandContext(ctx, "node")
	in := bytes.NewReader(code)
	cmd.Stdin = in

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("getting stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("getting stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("starting node process: %w", err)
	}

	return stdout, stderr, nil
}
