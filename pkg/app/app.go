package app

import (
	"context"
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

type App struct {
	config *Config
	http   *http.Server
	chat   *socketio.Server
}

func NewApp(config *Config) *App {
	return &App{config, nil, nil}
}

func (app *App) Start() {
	app.chat = NewChatServer()

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", app.chat)
	app.http = &http.Server{Addr: app.config.Addr, Handler: mux}

	go func() {
		if err := app.http.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	go app.chat.Serve()

	log.Printf("listening on %s", app.config.Addr)
}

func (app *App) Stop() {
	app.chat.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.http.Shutdown(ctx); err != nil {
		log.Fatal("stop failed:", err)
	}
	log.Println("exiting")
}
