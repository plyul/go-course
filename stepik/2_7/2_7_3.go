package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	var max byte
	for _, v := range []byte(input) {
		if v > max {
			max = v
		}
	}
	fmt.Println(max - 48)
}
