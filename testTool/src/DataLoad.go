package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	csvData := readCSV("../resource/testData.csv")
	print(csvData)
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error")
	}
}

func readCSV(readFile string) {
	fp, err := os.Open(readFile)
	failOnError(err)
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.Comma = ','

	// 行列の宣言がうまく行ってないっぽい
	var data [][]string{}

	for {
		recode, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		data = append(data, recode)
	}
	return data
}

// func readEXCEL(readFile string) {
//
// }
