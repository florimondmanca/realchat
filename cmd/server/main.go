package main

import (
	"os"
	"os/signal"

	"github.com/florimondmanca/realchat/pkg/app"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	config := app.NewConfig()
	app := app.NewApp(config)

	app.Start()
	defer app.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
