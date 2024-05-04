package restapi

import (
	"log/slog"
	"net/http"

	"github.com/Marattttt/newportfolio/config"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func useMiddleware(conf *config.AppConfig, mux *chi.Mux) {
	mux.Use(
		middleware.Heartbeat("/ping"),
		middleware.RequestID,
		middleware.Recoverer,
		useChiLogging(conf),
	)
}

// Middleware to configure and use chi httplog package
func useChiLogging(conf *config.AppConfig) func(next http.Handler) http.Handler {
	loglevel := slog.LevelInfo
	if conf.IsDebug() {
		loglevel = slog.LevelDebug
	}

	logger := httplog.NewLogger("portfolio", httplog.Options{
		LogLevel:       loglevel,
		Concise:        conf.IsDebug(),
		RequestHeaders: !conf.IsDebug(),
		Tags: map[string]string{
			"env": conf.Mode,
		},
	})

	return httplog.RequestLogger(logger)
}
