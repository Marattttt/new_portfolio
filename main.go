package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Marattttt/newportfolio/chats"
	"github.com/Marattttt/newportfolio/config"
	"github.com/Marattttt/newportfolio/restapi"
	"golang.org/x/sync/errgroup"
)

func main() {
	cancelsignals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	appCtx, appcancel := signal.NotifyContext(context.Background(), cancelsignals...)
	defer appcancel()

	logger := slog.Default()

	if err := config.ReadConfig(appCtx); err != nil {
		panic(err)
	}

	printConfig(config.Conf)

	chats.GlobalHub = chats.NewHub()

	server := restapi.Server(appCtx, config.Conf)

	go func() {
		logger.Info("Serving http", slog.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Unexpected server shutdown", slog.String("err", err.Error()))
		}
		appcancel()
	}()

	<-appCtx.Done()

	shutdown(config.Conf, logger, server)
}

func shutdown(conf *config.AppConfig, logger *slog.Logger, server *http.Server) {
	timeout := conf.ShutDownTimeout
	logger.Info("Beginning shutdown", slog.String("timeout", timeout.String()))

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var eg errgroup.Group

	eg.Go(shutdownGlobalChat)

	eg.Go(func() error {
		return shutdownHttp(ctx, server)
	})

	waitAll := func(eg *errgroup.Group) chan error {
		done := make(chan error)
		go func() {
			done <- eg.Wait()

		}()
		return done
	}

	select {
	case <-ctx.Done():
		logger.Error("Shutdown timeout", slog.String("timeout", timeout.String()))
	case err := <-waitAll(&eg):
		if err != nil {
			logger.Error("Shutdown not successful", slog.String("errs", err.Error()))
		}
	}

	logger.Info("Server shut down")
}

func shutdownGlobalChat() error {
	err := chats.GlobalHub.Close()
	if err != nil {
		return fmt.Errorf("closing global chat hub: %w", err)
	}

	return nil
}

func shutdownHttp(ctx context.Context, server *http.Server) error {
	err := server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("shutting down rest server: %w", err)
	}
	return nil
}

func printConfig(conf *config.AppConfig) {
	if !conf.IsDebug() {
		slog.Info("Using config", slog.Any("config", conf))
		return
	}

	// Print a pretrier version of config
	marshalledConf, err := json.MarshalIndent(conf, "", strings.Repeat(" ", 4))
	if err != nil {
		slog.Error("Marshalling created config", err)
		os.Exit(1)
	}

	slog.Info("Using config: \n" + string(marshalledConf))
}
