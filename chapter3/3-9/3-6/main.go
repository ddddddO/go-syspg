package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	a := io.NewSectionReader(programming, 5, 1)
	s := io.NewSectionReader(system, 0, 1)
	c := io.NewSectionReader(computer, 0, 1)
	i1 := io.NewSectionReader(programming, 8, 1)

	i2, pw := io.Pipe()

	go func() {
		io.CopyN(pw, io.NewSectionReader(programming, 8, 1), 1)
		pw.Close()
	}()

	stream = io.MultiReader(a, s, c, i1, i2)

	io.Copy(os.Stdout, stream)
}
