package main

import "fmt"

func main() {
	numCha := outPutNum()
	// メッセージをループしながら受信。channelが閉じられるまでループ
	// 値がくるまで回る。個数が未定の動的配列
	for n := range numCha {
		fmt.Println(n)
	}
}
