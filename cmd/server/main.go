package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/florimondmanca/realchat/internal/chat"
)

var (
	hostName = "localhost"
	addr     string
)

func init() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr = fmt.Sprintf("%s:%s", hostName, port)
	log.Println("Address:", addr)
}

func main() {
	app := NewApp()
	app.Start()
	defer app.Shutdown()
	app.Wait()
}

// App represents the main application
type App struct {
	server      *http.Server
	chat        *chat.Server
	wg          *sync.WaitGroup
	interrupted chan os.Signal
}

// NewApp builds an app and returns it
func NewApp() *App {
	chat := chat.NewServer()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/chat", chat.HandleChat)
	http.HandleFunc("/messages", chat.HandleMessages)

	server := &http.Server{Addr: addr, Handler: nil}

	wg := &sync.WaitGroup{}

	interrupted := make(chan os.Signal, 1)
	signal.Notify(interrupted, os.Interrupt)

	return &App{
		server,
		chat,
		wg,
		interrupted,
	}
}

func (app *App) Start() {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()
		app.chat.Listen()
	}()

	go func() {
		if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server: error: %s\n", err)
		}
	}()
}

func (app *App) Wait() {
	sig := <-app.interrupted
	log.Printf("Got signal: %s. Stopping...", sig)
}

func (app *App) Shutdown() {
	app.chat.Stop()
	app.wg.Wait()

	// Close connections within a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown failed: %s", err)
	}

	log.Println("Stopped")
}
