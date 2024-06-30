package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/Marattttt/new_portfoliio/services/jsrunner/coderun"
	"github.com/Marattttt/new_portfoliio/services/jsrunner/internal/config"
)

func main() {
	conf := config.DefaultConf()
	runner := coderun.LocalRunner{
		Logger: slog.Default(),
		Conf:   &conf.Runtime,
	}

	slog.SetLogLoggerLevel(slog.LevelDebug)
	result, err := runner.Exec(context.Background(), os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)

}

func getStdIn() ([]byte, error) {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("reading stdin: %w", err)
	}

	return stdin, nil
}
