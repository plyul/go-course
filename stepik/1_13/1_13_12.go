package main

import "fmt"

// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.
func main() {
	var n int
	fmt.Scan(&n)
	res := 1
	for {
		fmt.Printf("%d ", res)
		res = 2 * res
		if res > n {
			break
		}
	}
}
