package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

var contents = []string{
	" これは、私わたしが小さいときに、村の茂平もへいというおじいさんからきいたお話です。",
	" むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって、",
	" 中山さまというおとのさまが、おられたそうです。",
	" その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。",
	" ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。",
	" そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。",
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	log.Println("Server is running localhost:8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go processSession(conn)
	}
}

func processSession(conn net.Conn) {
	defer conn.Close()
	log.Printf("Accept %v\n", conn.RemoteAddr())

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				log.Println("TIMEOUT")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}

		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text-plain",
			"Transfer-Encoding: chunked",
			"",
			"",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}
