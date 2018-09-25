package main

import (
	//"flag"
	"log"
	"net/http"

	"github.com/andyvauliln/prediction-markets-api/app"
)

//var addr = flag.String("addr", "localhost:8080", "http service address")
var done = make(chan bool)
var config = app.NewDefaultConfiguration()

func main() {
	//flag.Parse()

	server := app.NewServer()

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		server.OpenWebsocketConnection(w, r)
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
