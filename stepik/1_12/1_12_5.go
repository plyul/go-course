package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	var c int
	for i := 0; i < n; i++ {
		if s[i] > 0 {
			c++
		}
	}
	fmt.Println(c)
}
