package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var PathDirectory = `./files/`

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/documents", getDocuments).Methods("GET")
	router.HandleFunc("/documents/{ID}", getDocumentsById).Methods("GET")
	router.HandleFunc("/documents", createDocument).Methods("POST")
	router.HandleFunc("/documents/{ID}", deleteDocument).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", router))
}

