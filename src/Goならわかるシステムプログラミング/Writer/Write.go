package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	const exWrite = "テスト出力"

	fmt.Println("start")

	// ウェブサーバからブラウザに書き出し
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)

	// バッファに書き出し（一時保存）
	// 	BufferWrite(exWrite)

	// ファイルに書き出し
	// 	FileWrite("/Users/hasegawahiroaki/go/src/resouce/output.txt")

	// コンソールに書き出し
	// 	ConsoleWrite(exWrite)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "sample")
}

// バッファに書き出し（一時保存）
func BufferWrite(str string) {
	var buffer bytes.Buffer
	buffer.Write([]byte(str))

	fmt.Println(buffer.String())
}

func FileWrite(path string) {

	file, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	const exp = "ファイル書き出し"
	file.Write([]byte(exp))

	return
}

func ConsoleWrite(str string) {
	os.Stdin.Write([]byte(str))
	return
}
