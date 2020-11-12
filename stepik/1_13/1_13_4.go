package main

import "fmt"

// Заданы три числа - a,b,c(a<b<c) - длины сторон треугольника.
// Нужно проверить, является ли треугольник прямоугольным.
func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	if (c * c) == (a*a)+(b*b) {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
