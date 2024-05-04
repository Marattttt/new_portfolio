package restapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Marattttt/newportfolio/chats"
	"github.com/Marattttt/newportfolio/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func Server(ctx context.Context, conf *config.AppConfig) *http.Server {
	router := chi.NewRouter()
	useMiddleware(conf, router)

	router.Handle("/", http.HandlerFunc(Test))

	listenOn := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)

	return &http.Server{
		Addr:    listenOn,
		Handler: router,
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	logger := httplog.LogEntry(r.Context())
	chats.HandleEcho(w, r, logger)
}
