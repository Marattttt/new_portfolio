package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Marattttt/newportfolio/config"
	"github.com/Marattttt/newportfolio/restapi"
)

func main() {
	cancelsignals := []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM}
	appCtx, appcancel := signal.NotifyContext(context.Background(), cancelsignals...)
	defer appcancel()

	if err := config.ReadConfig(appCtx); err != nil {
		panic(err)
	}

	printConfig(config.Conf)

	go func() {
		server := restapi.Server(appCtx, config.Conf)
		slog.Info("Serving http", slog.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil {
			slog.Error("Unexpected server shutdown", slog.String("err", err.Error()))
		}
		appcancel()
	}()

	<-appCtx.Done()
	slog.Info("Server shut down")
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
