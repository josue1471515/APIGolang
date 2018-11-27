package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
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
func deleteDocument(writer http.ResponseWriter, request *http.Request) {
	keys := mux.Vars(request)
	id := keys["ID"]
	var docs []Document
	docs = getDocumentsTypeList()
	sw := false
	docName := ""
	for _, doc := range docs {
		if doc.ID == id {
			sw = true
			docName = doc.Name
			os.Remove("./files/" + doc.Name)
		}
	}
	if !sw {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode("No se encontro el archivo ")
	} else {
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode("Se Elimino correctamente el archivo " + docName)
	}

}
