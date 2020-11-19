package main

import "fmt"

const CharZeroASCIICode = 48

func main() {
	var input string
	fmt.Scan(&input)
	var max byte
	for _, v := range []byte(input) {
		if v > max {
			max = v
		}
	}
	fmt.Println(max - CharZeroASCIICode) // Преобразование значения max из цифры в кодировке ASCII в просто цифру
}
