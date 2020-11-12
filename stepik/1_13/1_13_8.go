package main

import "fmt"

// Найдите количество минимальных элементов в последовательности.
func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	var c int
	for _, v := range s {
		if v == min {
			c++
		}
	}
	fmt.Println(c)
}
