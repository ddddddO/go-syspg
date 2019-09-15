package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	zipfile, err := os.Create("new.zip")
	if err != nil {
		panic(err)
	}
	defer zipfile.Close()

	zipWriter := zip.NewWriter(zipfile)
	defer zipWriter.Close()

	xwriter, err := zipWriter.Create("x.txt")
	if err != nil {
		panic(err)
	}
	xreader := strings.NewReader("xxxxx")
	io.Copy(xwriter, xreader)

	ywriter, err := zipWriter.Create("y.txt")
	if err != nil {
		panic(err)
	}
	yreader := strings.NewReader("yyyyy")
	io.Copy(ywriter, yreader)

	zwriter, err := zipWriter.Create("z.txt")
	if err != nil {
		panic(err)
	}
	zreader := strings.NewReader("zzzzz")
	io.Copy(zwriter, zreader)
}
