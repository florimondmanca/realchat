package app

import (
	"context"
	"log"
	"net/http"
	"time"
)

type App struct {
	config *Config
	http   *http.Server
	server *Server
}

func NewApp(config *Config) *App {
	server := NewChatServer()

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", server.Handler())
	srv := &http.Server{Addr: config.Addr(), Handler: mux}

	return &App{config, srv, server}
}

func (app *App) Start() {
	go func() {
		if err := app.http.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	app.server.Start()

	log.Printf("listening on %s", app.config.Addr())
}

func (app *App) Stop() {
	app.server.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.http.Shutdown(ctx); err != nil {
		log.Fatal("stop failed:", err)
	}
	log.Println("exiting")
}
