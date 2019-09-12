package main

import (
	"crypto/rand"
	"os"
)

func main() {
	b := make([]byte, 1024)

	rReader := rand.Reader
	rReader.Read(b)

	f, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(b)
}
