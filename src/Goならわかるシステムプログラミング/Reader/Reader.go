package main

import (
	"io"
	"os"
)

func FileReader(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	io.Copy(os.Stdout, file)
}

/*
func main() {
	// ファイルの中身を標準出力する
	FileReader("src/resouce/output.txt")
}
*/
