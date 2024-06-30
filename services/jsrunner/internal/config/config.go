package config

import "time"

func DefaultConf() *App {
	return &App{
		Runtime: Runtime{RunTimeout: time.Second * 10},
	}
}

type App struct {
	Runtime
}

type Runtime struct {
	RunTimeout time.Duration
}
