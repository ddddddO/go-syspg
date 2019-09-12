package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	mr()

	fmt.Println()

	tee()
}

func mr() {
	header := bytes.NewBufferString("-- HEADER --\n")
	content := bytes.NewBufferString("content\n")
	footer := bytes.NewBufferString("-- FOOTER --\n")

	content.Write([]byte("try!!\n"))

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)
}

func tee() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Tee!\n")
	teeReader := io.TeeReader(reader, &buffer)

	// rm
	_, _ = ioutil.ReadAll(teeReader)

	fmt.Println(buffer.String())
}
