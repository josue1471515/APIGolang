package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getDocuments(w http.ResponseWriter, r *http.Request) {
	//puntero lleva a la la variable
	//referencia copia
	var docs []Document
	docs = getDocumentsTypeList(PathDirectory)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)

}

func getDocumentsById(w http.ResponseWriter, r *http.Request) {
	//puntero lleva a la la variable
	//referencia copia
	//var docs []Document

	keys, ok := r.URL.Query()["ID"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		json.NewEncoder(w).Encode("Url Param 'key' is missing")
		return
	}
	//docs = getDocumentsTypeList(PathDirectory)
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(docs)
	key := keys[0]
	json.NewEncoder(w).Encode(key)
	log.Println("Url Param 'key' is: " + string(key))
}

func getDocumentsTypeList(pathDocument string) []Document{
	var docs []Document

	var fileList = getListFiles(pathDocument)
	for _, value := range fileList {

		absolutePath := pathDocument + value
		verifyValidPath(absolutePath)
		hashMd5, err := hashFileMd5(absolutePath)
		fileSize :=  getSizeFile(absolutePath)

		if err == nil {
			docs = append(docs,
				Document{ID: hashMd5, Name: value, Size: fileSize})
		}
	}
	return docs
}

func getIdMd5FIle(id string){

}
