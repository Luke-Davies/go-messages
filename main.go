package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// StoredMessages contains all messages sent to the server
var StoredMessages = Messages{}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Path("/messages").Methods("GET").HandlerFunc(GetMessages)
	router.Path("/messages").Methods("POST").HandlerFunc(CreateMessage)
	router.Path("/messages").Methods("DELETE").HandlerFunc(DeleteMessages)
	router.Path("/messages/{msgId}").Methods("GET").HandlerFunc(GetMessage)

	handler := handlers.LoggingHandler(os.Stdout, router)

	port := ":3000"
	if v := os.Getenv("PORT"); v != "" {
		port = ":" + v
	}
	log.Fatal(http.ListenAndServe(port, handler))
}
