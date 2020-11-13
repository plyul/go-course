package main

import (
	"fmt"
)

func main() {
	minimumFromFour()
}

func minimumFromFour() int {
	var b int
	min := 1<<63 - 1
	for i := 0; i < 4; i++ {
		fmt.Scan(&b)
		if b < min {
			min = b
		}
	}
	return min
}
