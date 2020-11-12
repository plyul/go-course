package main

import "fmt"

func main() {
	var workArray [10]uint8
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}
	for i := 0; i < 3; i++ {
		var i1, i2 int
		fmt.Scan(&i1)
		fmt.Scan(&i2)
		b := workArray[i1]
		workArray[i1] = workArray[i2]
		workArray[i2] = b
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", workArray[i])
	}
}
