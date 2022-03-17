package server

import (
	"log"
	"net/http"
)

func Main() {
	http.HandleFunc("/server", websocketServer)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
