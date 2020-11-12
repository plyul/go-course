package main

import "fmt"

// Найдите самое большее число на отрезке от a до b, кратное 7.
func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	var res int
	for i := a; i <= b; i++ {
		if i%7 == 0 && i > res {
			res = i
		}
	}
	if res != 0 {
		fmt.Println(res)
	} else {
		fmt.Println("NO")
	}
}
