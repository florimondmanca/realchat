package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/florimondmanca/chatroom/server/chat"
	"github.com/florimondmanca/chatroom/server/config"
)

var configuration config.Configuration
var serverHostName string

func init() {
	configuration = config.Load()
	serverHostName = fmt.Sprintf("%s:%s",
		configuration.Host, strconv.Itoa(configuration.Port))
	log.Println("Host name:", serverHostName)
}

func main() {
	server := chat.NewServer()
	go server.Listen()

	http.HandleFunc("/", handleIndex)
	log.Fatal(http.ListenAndServe(serverHostName, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
