package main

import "fmt"

// Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток.
func main() {
	var k int
	fmt.Scan(&k)
	h := k / 3600
	m := (k - (h * 3600)) / 60
	fmt.Println("It is", h, "hours", m, "minutes.")
}
