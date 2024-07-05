package main

import (
	"context"
	"log"
	"log/slog"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Marattttt/new_portfoliio/services/jsrunner/config"
	"github.com/Marattttt/new_portfoliio/services/jsrunner/internal/grpc"
	"github.com/Marattttt/new_portfoliio/services/jsrunner/runners"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	conf, err := config.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	runner := runners.NewLocal(slog.Default(), &conf.Runtime, runners.LocalNode{})

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
