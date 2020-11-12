package main

import "fmt"

// Дано трехзначное число. Найдите сумму его цифр.
func main() {
	var n string
	fmt.Scan(&n)
	var sum int
	for _, v := range n {
		sum += int(v) - 48
	}
	fmt.Println(sum)
}
