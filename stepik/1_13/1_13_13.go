package main

import "fmt"

// Номер числа Фибоначчи
func main() {
	var a int
	fmt.Scan(&a)
	n1 := 1
	n2 := 1
	c := 3
	for {
		nextFibo := n1 + n2
		if nextFibo == a {
			fmt.Println(c)
			break
		}
		if nextFibo > a {
			fmt.Println(-1)
			break
		}
		c++
		n1 = n2
		n2 = nextFibo
	}
}
