package main

import "fmt"

// По данным числам, определите количество чисел, которые равны нулю.
func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	var c int
	for _, v := range s {
		if v == 0 {
			c++
		}
	}
	fmt.Println(c)
}
