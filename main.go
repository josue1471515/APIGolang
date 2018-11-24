package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
var PathDirectory =`C:\Users\PC\go\src\dss-persistence\files\`

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/documents", getDocuments).Methods("GET")
	router.HandleFunc("/document", getDocumentsById).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", router))
}
