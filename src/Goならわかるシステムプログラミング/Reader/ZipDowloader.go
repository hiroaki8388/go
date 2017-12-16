package main

import (
	"archive/zip"
	"io"
	"net/http"
	// 	"net/http"
	"os"
)

// pathに指定してあるファイルをzip化し、対象のポートからダウンロード出来る
func zipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; sample.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	path := "/Users/hasegawahiroaki/test.txt"

	a, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	defer a.Close()

	b, err := zipWriter.Create("data.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(b, a)

}

// func main() {
//
// 	http.HandleFunc("/", zipHandler)
// 	http.ListenAndServe(":8080", nil)
// }
