package main

import (
	"fmt"
)

func f(str string) {
	fmt.Printf(str)
}

// ListenChannel:myChからの通知をうけとる
func ListenChannel(myCh chan int) {
	// myChから送信を受け取る
	for out := range myCh {
		fmt.Println(out)
	}
}
