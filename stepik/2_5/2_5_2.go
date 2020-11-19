package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	inputString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading string: %s", err.Error())
		return
	}
	runes := []rune(strings.TrimSpace(inputString))
	lh := 0
	rh := len(runes) - 1
	for runes[lh] == runes[rh] {
		if lh >= rh {
			fmt.Println("Палиндром")
			return
		}
		lh++
		rh--
	}
	fmt.Println("Нет")
}
