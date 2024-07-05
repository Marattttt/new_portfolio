package runners

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/Marattttt/new_portfoliio/services/jsrunner/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNode struct {
	mock.Mock
}

func (m *MockNode) NodeRun(ctx context.Context, code string) (io.Reader, io.Reader, error) {
	args := m.Called(ctx, code)
	arg1 := args[0].(io.Reader)
	arg2 := args[1].(io.Reader)

	// If a nil errro is returned, it fails assertion to error type
	var arg3 error
	if args[2] == nil {
		arg3 = nil
	} else {
		arg3 = args[2].(error)
	}

	return arg1, arg2, arg3
}

func TestRunJsHelloWorld(t *testing.T) {
	mockNode := new(MockNode)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	mockNode.On("NodeRun", ctx, `console.log("Hello world")`).
		Return(bytes.NewBuffer([]byte("Hello world")), &bytes.Buffer{}, nil)

	logOut := bytes.Buffer{}
	logger := slog.New(slog.NewTextHandler(&logOut, nil))

	runner := NewLocal(logger, &config.DefaultConfig().Runtime, mockNode)

	req := RunRequest{Code: `console.log("Hello world")`}
	res, err := runner.RunJs(ctx, req)
	if assert.NoError(t, err) {
		assert.Equal(t, res, RunResult{Output: "Hello world", ExitCode: 0, Err: ""})
	}
}

func TestRunJsEmpty(t *testing.T) {
	mockNode := new(MockNode)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	mockNode.On("NodeRun", ctx, "").
		Return(bytes.NewBuffer([]byte{}), &bytes.Buffer{}, nil)

	logOut := bytes.Buffer{}
	logger := slog.New(slog.NewTextHandler(&logOut, nil))

	runner := NewLocal(logger, &config.DefaultConfig().Runtime, mockNode)

	req := RunRequest{}
	res, err := runner.RunJs(ctx, req)
	if assert.NoError(t, err) {
		assert.Equal(t, res, RunResult{Output: "", ExitCode: 0, Err: ""})
	}
}

func TestReadAll(t *testing.T) {
	var (
		out    = []byte{1, 2}
		errout = []byte{3, 4}
	)

	resout, reserrout, err := readAllOutput(bytes.NewBuffer(out), bytes.NewBuffer(errout))
	if assert.NoError(t, err) {
		assert.Equal(t, resout, out)
		assert.Equal(t, reserrout, errout)
	}
}

func TestReadAllTimeout(t *testing.T) {
	const timeout = time.Second / 2
	msg := []byte("done")

	outr, outw := io.Pipe()
	erroutr, erroutw := io.Pipe()

	go writeCloseAfter(outw, msg, timeout)
	// Ensure Does not return on first timeout
	go writeCloseAfter(erroutw, msg, timeout*2)

	resOut, resErrout, err := readAllOutput(outr, erroutr)
	if assert.NoError(t, err) {
		assert.Equal(t, msg, resOut, "Should wait for timeout")
		assert.Equal(t, msg, resErrout, "Should wait for both timeouts")
	}
}

func writeCloseAfter(to io.WriteCloser, msg []byte, after time.Duration) {
	time.Sleep(after)
	to.Write(msg)
	to.Close()
}
