package main

import (
	"api-crud-gorilla/handlers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/messages", handlers.GetMessages).Methods("GET")
	r.HandleFunc("/messages/create", handlers.CreateMessage).Methods("POST")
	r.HandleFunc("/messages/update/{id}", handlers.UpdateMessage).Methods("PUT")
	r.HandleFunc("/messages/delete/{id}", handlers.DeleteMessage).Methods("DELETE")

	server := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening at port 8080...")
	log.Fatal(server.ListenAndServe())

}
