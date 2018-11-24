package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getDocuments(w http.ResponseWriter, r *http.Request) {
	//puntero lleva a la la variable
	//referencia copia
	var docs []Document
	docs = getDocumentsTypeList()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)

}

func getDocumentsById(w http.ResponseWriter, r *http.Request) {
	keys := mux.Vars(r)
	id := keys["ID"]
	log.Println(keys)
	doc, err := getDocumentByIdMd5(id)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(doc)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}
}

func getDocumentsTypeList() []Document {
	var docs []Document

	var fileList = getListFiles(PathDirectory)
	for _, value := range fileList {

		absolutePath := PathDirectory + value
		verifyValidPath(absolutePath)
		hashMd5, err := hashFileMd5(absolutePath)
		fileSize := getSizeFile(absolutePath)

		if err == nil {
			docs = append(docs,
				Document{ID: hashMd5, Name: value, Size: fileSize})
		}
	}
	return docs
}

func getDocumentByIdMd5(id string) (Document, error) {

	var docs []Document
	docs = getDocumentsTypeList()
	for _, doc := range docs {
		if doc.ID == id {
			return doc, nil
		}
	}
	return Document{}, errors.New("No se encontro el ID ")
}
