package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/Marattttt/new_portfoliio/services/jsrunner/config"
	"github.com/Marattttt/new_portfoliio/services/jsrunner/runners"
)

func main() {
	conf, err := config.NewApp(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	runner := runners.NewLocal(slog.Default(), &conf.Runtime, runners.LocalNode{})

	slog.SetLogLoggerLevel(slog.LevelDebug)
	result, err := runner.RunJsReader(context.Background(), os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
