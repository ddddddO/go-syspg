package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	//stdin() // go run main.go < main.go
	//file()
	//receiveNet()
	//strReader()
	//cutReader()
	endian()
}

// 標準入力
func stdin() {
	for {
		buf := make([]byte, 5)
		size, err := os.Stdin.Read(buf)
		if err == io.EOF {
			os.Stdout.Write([]byte("EOF!\n"))
			break
		}
		fmt.Printf("size=%d input=%s\n", size, string(buf))
	}
}

// ファイル入力
func file() {
	f, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(os.Stdout, f)
}

func receiveNet() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)

	fmt.Println(res.Status)
	fmt.Println()
	fmt.Println(res.Header)
	fmt.Println()
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

func strReader() {
	sr := strings.NewReader("ReaderReaderrrrrrrrr")
	io.Copy(os.Stdout, sr)
}

// 必要なとこを切り出す
func cutReader() {
	s := "abcdefghi Section jklmn..."
	sr := strings.NewReader(s)

	// 先頭から10バイトのみ読み込む
	lReader := io.LimitReader(sr, 10)
	io.Copy(os.Stdout, lReader)

	fmt.Println()

	// 指定位置から7バイト読み込む
	sReader := io.NewSectionReader(sr, 10, 7)
	io.Copy(os.Stdout, sReader)
}

// エンディアン変換 ref(エンディアンとは): https://wa3.i-3-i.info/diff112endiannes.html
// よくわからない。p48の図でなんとなくわかりそう
func endian() {
	// 32bitのビッグエンディアンデータ(10000)
	data := []byte{0x0, 0x0, 0x27, 0x10}
	fmt.Printf("data(string): %s\n", string(data)) // for debug 意味なさそう

	var i int32
	// エンディアン変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)

	// --for debug-- 以下の結果と↑の結果を比較すると、binary.Readで変換しないとダメらしい
	d := []byte{0x10, 0x27, 0x0, 0x0}
	var ii int32
	binary.Read(bytes.NewReader(d), binary.LittleEndian, &ii)
	fmt.Printf("d: %d\n", ii)
}
