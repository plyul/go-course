package main

import "fmt"

// Дано трехзначное число. Переверните его, а затем выведите.
func main() {
	var n int
	fmt.Scan(&n)
	d1 := (n - n%100) / 100
	d2 := (n % 100) / 10
	d3 := n % 10
	res := d3*100 + d2*10 + d1
	fmt.Println(res)
}
