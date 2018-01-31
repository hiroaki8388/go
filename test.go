package main

func main() {
	println(fib(10))

	println(fib2(10))
}

func fib(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	return fib(i-1) + fib(i-2)

}

// 少しリファクタ
func fib2(i int) int {
	if i == 0 || i == 1 {
		return i
	}
	return fib2(i-1) + fib2(i-2)
}
