package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/florimondmanca/chatroom/chat"
	"github.com/florimondmanca/chatroom/config"
)

var configuration config.Configuration
var serverHostName string

func init() {
	configuration = config.Load()
	serverHostName = fmt.Sprintf(":%s", strconv.Itoa(configuration.Port))
}

func main() {
	chatServer := chat.NewServer()
	go chatServer.Listen()

	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(serverHostName, nil)
}

func handleIndex(rw http.ResponseWriter, request *http.Request) {
	http.ServeFile(rw, request, "index.html")
}
