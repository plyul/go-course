package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cow := "korov"
	if n == 1 || (n > 20 && (n%10 == 1)) {
		cow = "korova"
	}
	if (n < 12 || n > 14) && (n%10 >= 2 && n%10 <= 4) {
		cow = "korovy"
	}
	fmt.Println(n, cow)
}
