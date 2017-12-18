package main

import (
	"fmt"
	"time"
)

func main() {
	//	message := make(chan string)

	// 	go func() { message <- "ping" }()
	// 	msg := <-message
	// 	fmt.Printf(msg)

	go f("hello")

	go func() {
		fmt.Printf("hello_go")
	}()

	// 非同期で処理しているので、mainが終わらないように意図的に止めている
	time.Sleep(time.Second)
}

func f(str string) {
	fmt.Printf(str)
}

// func ListenChannel(myCh chan int) {
// 	// myChから送信を受け取る
// 	for out := range myCh {
// 		fmt.Printf(out)
// 	}
// }
