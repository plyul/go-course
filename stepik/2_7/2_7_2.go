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
	output := make([]rune, 0, len(runes)*2)
	for _, v := range runes {
		output = append(output, v, '*')
	}
	output = output[:len(output)-1]
	fmt.Println(string(output))
}
