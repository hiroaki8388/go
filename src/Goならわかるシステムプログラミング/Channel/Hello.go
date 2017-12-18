package main

import (
	"fmt"
)

func f(str string) {
	fmt.Printf(str)
}

func ListenChannel(myCh chan int) {
	// myChから送信を受け取る
	for out := range myCh {
		fmt.Println(out)
	}
}
