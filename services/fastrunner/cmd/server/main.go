package main

import (
	"context"
	"log"
	"log/slog"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Marattttt/newportfolio/services/fastrunner/config"
	"github.com/Marattttt/newportfolio/services/fastrunner/grpc"
	"github.com/Marattttt/newportfolio/services/fastrunner/runners"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	conf, err := config.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := runners.InitGoEnv(&conf.Runners); err != nil {
		log.Fatal(err)
	}

	runner := runners.NewGoRunner(slog.Default(), &conf.Runners)

	server := grpc.NewServer(slog.Default(), &runner)

	go func() {
		if err := server.ListenAndServe(ctx, &conf.Server); err != nil {
			slog.Error("Unexpected grpc shutdown", slog.String("err", err.Error()))
		}
	}()

	<-ctx.Done()
	const shutdownTimeOut = time.Second * 10
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeOut)
	defer shutdownCancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		stopGrpc(shutdownCtx, slog.Default(), &server)
	}()

	wg.Wait()
	slog.Info("Finished shutdown")
}

func stopGrpc(ctx context.Context, logger *slog.Logger, server *grpc.Server) {
	if err := server.Stop(ctx); err != nil {
		logger.Error("Shutting down grpc server", slog.String("err", err.Error()))
	}
}
