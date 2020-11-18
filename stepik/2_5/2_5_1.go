package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	inputString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading string: %s", err.Error())
		return
	}
	runes := []rune(strings.TrimSpace(inputString))
	isFirstCapital := unicode.IsUpper(runes[0])
	isLastDot := runes[len(runes)-1] == '.'
	if isFirstCapital && isLastDot {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
