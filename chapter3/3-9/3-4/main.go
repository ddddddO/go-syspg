package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/try", tryHandler{})
	http.Handle("/zip", zipHandler{})

	http.ListenAndServe(":8888", nil)
}

type tryHandler struct{}

func (th tryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TRY!!\n"))
	w.Write([]byte(r.RemoteAddr))
}

type zipHandler struct{}

func (zh zipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := genZipFile()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("sorry.."))
		return
	}

	w.Header().Set("content-type", "application/zip")
	w.Header().Set("content-disposition", "attachment; filename=server.zip")

	http.ServeFile(w, r, "./server.zip")
}

func genZipFile() error {
	file, err := os.Create("server.zip")
	if err != nil {
		return err
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	xwriter, err := zipWriter.Create("x.txt")
	if err != nil {
		return err
	}
	io.Copy(xwriter, strings.NewReader("server xxx"))

	return nil
}
