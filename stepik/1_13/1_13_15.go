package main

import "fmt"

// Из натурального числа удалить заданную цифру.
func main() {
	var n, f, r int
	fmt.Scan(&n)
	fmt.Scan(&f)

	k := 1
	for {
		d := n % 10
		if n < 10 && n != f {
			r += d * k
			break
		}
		n /= 10
		if d != f {
			r += d * k
			k *= 10
		}
	}
	fmt.Println(r)
}
