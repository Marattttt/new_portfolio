package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/Marattttt/newportfolio/services/gorunner/config"
	"github.com/Marattttt/newportfolio/services/gorunner/grpc/grpcgen/gogen"
	"github.com/Marattttt/newportfolio/services/gorunner/runners"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var ErrInternal = fmt.Errorf("Something went wrong")

func FormatRunResult(res runners.RunResult) *gogen.RunResponse {
	err := ""
	if res.Err != nil {
		err = res.Err.Error()
	}
	return &gogen.RunResponse{
		Error:  err,
		Output: res.Text,
	}
}

type CodeRunner interface {
	Run(ctx context.Context, code string) (runners.RunResult, error)
}

type Server struct {
	Logger   *slog.Logger
	GoRunner CodeRunner

	gogen.UnimplementedGoRunnerServer
	grpcServer *grpc.Server
}

func NewServer(logger *slog.Logger, goRunner CodeRunner) Server {
	return Server{
		Logger:   logger,
		GoRunner: goRunner,
	}
}

func (s *Server) ListenAndServe(ctx context.Context, conf *config.Server) error {
	if s.grpcServer != nil {
		return fmt.Errorf("Server already started")
	}

	listenOn := fmt.Sprintf(":%d", conf.GrpcPort)
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("starting a listener: %w", err)
	}

	s.grpcServer = grpc.NewServer()
	gogen.RegisterGoRunnerServer(s.grpcServer, s)
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

func (s Server) RunGoLang(ctx context.Context, req *gogen.RunGoRequest) (*gogen.RunResponse, error) {
	res, err := s.GoRunner.Run(ctx, req.Code)

	if err != nil {
		s.Logger.Error("Running code", slog.String("lang", "go"), slog.String("err", err.Error()))

		return nil, ErrInternal
	}

	return FormatRunResult(res), nil
}
