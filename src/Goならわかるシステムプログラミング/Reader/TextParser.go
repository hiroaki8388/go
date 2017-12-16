package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("ファイルの読み込み開始")
	var path string
	if len(os.Args) < 2 {
		path = "/Users/hasegawahiroaki/golang/src/Goならわかるシステムプログラミング/Chr3/csvFile.csv"
	} else {
		path = os.Args[1]
	}
	csvParser(path)
}

// csvをパースして出力する
func csvParser(path string) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanaer := csv.NewReader(file)
	// scanaer.Comma = ","
	for {
		record, err := scanaer.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}
}
