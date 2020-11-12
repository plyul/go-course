package main

import "fmt"

func main() {
	var a [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}
	var max int
	for i := 0; i < 5; i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	fmt.Println(max)
}
