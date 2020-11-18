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
	for _, r := range runes {
		if strings.Count(string(runes), string(r)) < 2 {
			fmt.Print(string(r))
		}
	}
}
