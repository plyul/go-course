package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", s[i])
		}
	}
}
