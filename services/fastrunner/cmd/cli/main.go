package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/Marattttt/newportfolio/services/fastrunner/config"
	"github.com/Marattttt/newportfolio/services/fastrunner/runners"
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

	if res, err := runner.Run(string(input), "cli-app"); err != nil {
		log.Fatal(err)
	} else {
		log.Println(res)
	}
}
