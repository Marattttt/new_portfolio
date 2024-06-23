package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/Marattttt/newportfolio/services/multirunner/config"
	"github.com/Marattttt/newportfolio/services/multirunner/runners"
)

func main() {
	conf, err := config.NewApp(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	if err := runners.InitGoEnv(&conf.Runners); err != nil {
		log.Fatal(err)
	}

	runner := runners.NewGoRunner(slog.Default(), &conf.Runners)

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	if res, err := runner.Run(context.Background(), string(input)); err != nil {
		log.Fatal(err)
	} else {
		log.Println(res)
	}
}
