package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/Marattttt/new_portfoliio/services/jsrunner/config"
	"github.com/Marattttt/new_portfoliio/services/jsrunner/internal/grpc/jsgen"
	"github.com/Marattttt/new_portfoliio/services/jsrunner/runners"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var ErrInternal = fmt.Errorf("Something went wrong")

type CodeRunner interface {
	RunJs(context.Context, runners.RunRequest) (runners.RunResult, error)
}

type Server struct {
	Logger   *slog.Logger
	JsRunner CodeRunner

	jsgen.UnimplementedJsRunnerServer
	grpcServer *grpc.Server
}

func NewServer(logger *slog.Logger, goRunner CodeRunner) Server {
	return Server{
		Logger:   logger,
		JsRunner: goRunner,
	}
}

func (s Server) RunJs(ctx context.Context, req *jsgen.JsRunRequest) (*jsgen.JsRunResponse, error) {
	res, err := s.JsRunner.RunJs(ctx, formatGenRequest(req))

	if err != nil {
		s.Logger.Warn("Error running code", slog.String("err", err.Error()))

		return nil, ErrInternal
	}

	return formatRunResult(res), nil
}

func (s *Server) ListenAndServe(ctx context.Context, conf *config.Server) error {
	if s.grpcServer != nil {
		return fmt.Errorf("Server already started")
	}

	listenOn := fmt.Sprintf(":%d", conf.Port)
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("starting a listener: %w", err)
	}

	s.grpcServer = grpc.NewServer()
	jsgen.RegisterJsRunnerServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)

	s.Logger.Info("Starting new grpc server", slog.String("on", listenOn))
	if err := s.grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.Logger.Info("Stopping serving grpc")
	stopChan := make(chan struct{})
	go func() {
		s.grpcServer.GracefulStop()
		stopChan <- struct{}{}
	}()

	select {
	case <-stopChan:
		s.Logger.Info("Finished serving grpc")
		return nil
	case <-ctx.Done():
		s.Logger.Warn("Cannot gracefully close grpc server, beginning force stop")
		s.grpcServer.Stop()
		s.Logger.Warn("Failed to gracefully close grpc server")
		return fmt.Errorf("Provided context done before graceful shutdown of grpc server")
	}
}

func formatGenRequest(req *jsgen.JsRunRequest) runners.RunRequest {
	return runners.RunRequest{
		Code: req.Code,
	}

}

func formatRunResult(res runners.RunResult) *jsgen.JsRunResponse {
	err := ""
	if res.Err != nil {
		err = res.Err.Error()
	}
	return &jsgen.JsRunResponse{
		Error:  err,
		Output: string(res.Output),
	}
}
