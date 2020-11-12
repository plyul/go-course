package main

import "fmt"

// Цифровой корень
func main() {
	var n int
	fmt.Scan(&n)
	digits := make([]int, 8) // 10^7 = 10000000, 8 digits
	sum := n                 // Closes case when n < 10
	for n > 9 {
		digits = digits[0:0]
		for n > 9 {
			digits = append(digits, n%10)
			n /= 10
		}
		digits = append(digits, n)
		sum = 0
		for _, v := range digits {
			sum = sum + v
		}
		n = sum
	}

	fmt.Println(sum)
}
