package main

import "fmt"

func main() {
	cache := make(map[int]int)
	for i := 0; i < 10; i++ {
		var input int
		fmt.Scan(&input)
		output, ok := cache[input]
		if !ok {
			output = work(input)
			cache[input] = output
		}
		fmt.Printf("%d ", output)
	}
}
