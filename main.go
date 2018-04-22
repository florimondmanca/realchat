package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/florimondmanca/go-live-chat/chat"
)

var serverHostName string

func init() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverHostName = fmt.Sprintf(":%s", port)
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
