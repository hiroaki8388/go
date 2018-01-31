package main

import "time"

func outPutNum() chan int {
	result := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// 仮想的な重い処理
			time.Sleep(time.Second)
			reslt <- i
		}
		close(result)
	}()

	return result
}
