package main

import (
	"io/ioutil"
	"testing"
)

func TestGetListFiles(t *testing.T) {
	files := getListFiles("./files/")
	numberFiles := len(files)
	casca, _ := ioutil.ReadDir("./files/")
	expected := len(casca)
	if numberFiles != expected {
		t.Errorf("numero de archivos incorecto: %d, Son : %d.", numberFiles, expected)
	}
}

func TestHashFileMd5(t *testing.T) {
	expected := "765e6aec6d47d10b08dd325aa6b5128d"
	idMd5, _ := hashFileMd5(`C:\Users\PC\go\src\dss-persistence\files\fileB`)

	if idMd5 != expected {
		t.Errorf(" los valores no coinciden : " + idMd5 + " Esperado :" + expected)
	}

}
