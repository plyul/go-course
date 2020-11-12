package main

import "fmt"

// Даны два числа. Найти их среднее арифметическое.
func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(float64(a+b) / 2)
}
