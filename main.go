package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	InitialMigration()
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", rootHandler).Methods("GET")
	myRouter.HandleFunc("/world", HandleGetAllWorlds).Methods("GET")
	myRouter.HandleFunc("/world", HandleWorldOptions).Methods("OPTIONS")
	myRouter.HandleFunc("/world", HandleCreateWorld).Methods("POST")
	myRouter.HandleFunc("/world/{id}", HandleGetWorld).Methods("GET")
	myRouter.HandleFunc("/world/{id}", HandleUpdateWorld).Methods("PUT")
	myRouter.HandleFunc("/world/{id}", HandleDeleteWorld).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello, world")
}
