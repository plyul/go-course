package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	for _, v := range []byte(input) {
		v -= 48
		fmt.Print(v * v)
	}
}
