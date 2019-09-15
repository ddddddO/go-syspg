package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("123456789")
	io.CopyN(os.Stdout, r, 5)
}
