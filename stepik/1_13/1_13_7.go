package main

import "fmt"

// По данным числам, определите количество чисел, которые равны нулю.
func main() {
	var n, b, c int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		if b == 0 {
			c++
		}
	}
	fmt.Println(c)
}
