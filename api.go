package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func getDocuments(w http.ResponseWriter, r *http.Request) {
	//puntero lleva a la la variable
	//referencia copia
	var docs []Document
	docs = getDocumentsTypeList()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(docs)

}

func getDocumentsById(w http.ResponseWriter, r *http.Request) {
	keys := mux.Vars(r)
	id := keys["ID"]
	log.Println(keys)
	doc, err := getDocumentByIdMd5(id)
	if err == nil {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(doc)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
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

func createDocument(writer http.ResponseWriter, request *http.Request) {
	//request.ParseMultipartForm(32 << 20)
	file, handler, err := request.FormFile("file")
	defer file.Close()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errors.New("No se puede extraer el archivo")
	}
	f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errors.New("No se puedo guardar el archivo")
	}

	defer f.Close()
	io.Copy(f, file)
	writer.WriteHeader(http.StatusForbidden)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode("Se guardo correctamente el archivo " + handler.Filename)
}
