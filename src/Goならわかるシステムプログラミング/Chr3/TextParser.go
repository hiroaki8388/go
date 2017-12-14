package main

import (
	"os"
)

func main() {

	path := "/Users/hiroaki/memo/test.txt"
	parser(path)
}

func parser(path string) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	os.Stdout(file)

}
