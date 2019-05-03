package main

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func problems() {
	// Q2.1 ファイルに対するフォーマット出力
	a1 := func() {
		file, err := os.Create("a1.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		var (
			d = 123
			s = "sss"
			f = 99.999999
		)
		fmt.Fprintf(file, "%d日は%sの日です。温度は%fです。\n", d, s, f)
	}

	// Q2.2 CSV出力
	a2 := func() {
		data := [][]string{
			{"first_name", "last_name", "username"},
			{"Rob", "Pike", "rob"},
			{"Ken", "Thompson", "ken"},
			{"Robert", "Griesemer", "gri"},
		}

		// 標準出力に出力
		a2_1 := func() {
			stdoutWriter := csv.NewWriter(os.Stdout)
			for _, record := range data {
				stdoutWriter.Write(record)
			}
			stdoutWriter.Flush()

			if err := stdoutWriter.Error(); err != nil {
				panic(err)
			}
		}

		// ファイルに出力
		a2_2 := func() {
			f, err := os.Create("a2_2.csv")
			if err != nil {
				panic(err)
			}
			defer f.Close()

			fileWriter := csv.NewWriter(f)
			fileWriter.WriteAll(data)

			if err = fileWriter.Error(); err != nil {
				panic(err)
			}
		}

		a2_1()
		a2_2()
	}

	// Q2.3 gzipされたJSON出力をしながら、標準出力にログを出力
	a3 := func() {
		handler := func(w http.ResponseWriter, r *http.Request) {
			// ref: https://blog.mmmcorp.co.jp/blog/2017/08/23/golang-s3-zip/
			w.Header().Set("Content-Type", "application/gzip")                           // edit/ ref: http://www.kyoto-su.ac.jp/ccinfo/use_web/mine_contenttype/
			w.Header().Set("Content-Disposition", `attachment; filename="resp.json.gz"`) // add

			// json化する元データ
			source := map[string]string{
				"Hello": "World",
			}
			// ここにコードを書く

			// バッファに整形済jsonを詰める
			var buf bytes.Buffer
			bufEncoder := json.NewEncoder(&buf)
			bufEncoder.SetIndent("", "    ")
			bufEncoder.Encode(source)

			/* 不要
			f, err := os.Create("a3.gz")
			if err != nil {
				panic(err)
			}
			defer f.Close()

			// ファイルの書き出し用意
			gzfw := gzip.NewWriter(f)
			*/

			gzfw := gzip.NewWriter(w)
			defer gzfw.Close()
			gzfw.Header.Name = "resp.json"

			// マルチな(ファイル・標準出力)writer生成
			writer := io.MultiWriter(gzfw, os.Stdout)
			// 詰めたjsonをgzファイル・標準出力へわたす。
			io.Copy(writer, &buf)
			// gzファイルに書き込み完了(無くても書き込まれている。。)
			gzfw.Flush()
		}

		// サーバー起動
		http.HandleFunc("/", handler)
		http.ListenAndServe(":8099", nil)
	}

	go a1()
	go a2()
	a3()
}

// io.Writerインタフェースを満たすものら
func main() {
	//outFile()
	//outDisplay()
	//outBufferByBytes()
	//sendNet()
	//ioDecorators()
	//formats()

	problems()
}

// ファイル出力
func outFile() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte("aaaaatest"))
}

// 画面出力
func outDisplay() {
	os.Stdout.Write([]byte("out display?\n"))
	os.Stdout.WriteString("out display??\n")
}

// 書かれた内容を記憶しておくバッファ
func outBufferByBytes() {
	// bytes.Buffer
	f1 := func(s string) {
		os.Stdout.WriteString("PATTERN: " + s + "\n")

		var buf bytes.Buffer
		buf.Write([]byte("bybyby\n"))
		buf.WriteString("BYBY\n")
		os.Stdout.WriteString(buf.String())

		os.Stdout.WriteString("-1-\n")
		os.Stdout.WriteString(buf.String())

		l1, err := buf.ReadString('\n')
		if err != nil {
			panic(err)
		}
		os.Stdout.WriteString("-2-\n")
		os.Stdout.WriteString(l1)

		l2, err := buf.ReadString('\n')
		if err != nil {
			panic(err)
		}
		os.Stdout.WriteString("-3-\n")
		os.Stdout.WriteString(l2)

		os.Stdout.WriteString("-4-\n")
		os.Stdout.WriteString(buf.String())
	}

	// strings.Builder
	f2 := func(s string) {
		os.Stdout.WriteString("PATTERN: " + s + "\n")

		var builder strings.Builder
		builder.Write([]byte("stststs"))
		builder.WriteString("STSTSTR\n")

		os.Stdout.WriteString(builder.String())
	}

	f1("bytes.Buffer")
	f2("strings.Builder")
}

// インターネットアクセスの送信
func sendNet() {
	//conn, err := net.Dial("tcp", "ascii.jp:80")
	conn, err := net.Dial("tcp", "localhost:4567")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: localhost\r\n\r\n")
	io.Copy(os.Stdout, conn)
}

// io.Writerのデコレータ
func ioDecorators() {
	f1 := func() {
		f, err := os.Create("multiwriter.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		writer := io.MultiWriter(f, os.Stdout, os.Stderr)
		io.WriteString(writer, "Hello World!!!\n")
	}

	f2 := func() {
		f, err := os.Create("compless.txt.gz")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		writer := gzip.NewWriter(f)
		defer writer.Close()
		writer.Header.Name = "compless.txt"
		io.WriteString(writer, "COMPLESS\n")
	}

	f1()
	f2()
}

// フォーマットしてデータをio.Writerに書き出す
func formats() {
	f1 := func() {
		fmt.Fprintf(os.Stdout, "TIME: %v, PLACE: %v\n\n", time.Now(), "aaa")
	}

	f2 := func() {
		data := map[string]string{
			"example": "encoding/json",
			"Hello":   "World",
		}

		f2_1 := func() {
			// jsonを画面に出力
			stdoutEncoder := json.NewEncoder(os.Stdout)
			stdoutEncoder.SetIndent("", "	")
			stdoutEncoder.Encode(data)
		}

		f2_2 := func() {
			// jsonをファイルに出力
			f, err := os.Create("test.json")
			if err != nil {
				panic(err)
			}
			defer f.Close()
			fileEncoder := json.NewEncoder(f)
			fileEncoder.SetIndent("", "    ")
			fileEncoder.Encode(data)
		}

		f2_1()
		f2_2()
	}

	f1()
	f2()
}
