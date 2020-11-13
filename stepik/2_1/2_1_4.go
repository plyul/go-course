package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, ":", fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n < 3 {
		return 1
	}
	fiPPrev := 1
	fiPrev := 1
	fiCur := 2
	for i := 3; i <= n; i++ {
		fiCur = fiPPrev + fiPrev
		fiPPrev = fiPrev
		fiPrev = fiCur
	}
	return fiCur
}
