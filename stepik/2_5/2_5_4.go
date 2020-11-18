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
	var output []rune
	for i, rune := range runes {
		if i%2 != 0 {
			output = append(output, rune)
		}
	}
	fmt.Println(string(output))
}
