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
	password, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading password: %s", err.Error())
		return
	}
	runes := []rune(strings.TrimSpace(password))
	if len(runes) < 5 {
		fmt.Println("Wrong password")
		return
	}
	for _, v := range runes {
		if !(unicode.Is(unicode.Latin, v) || unicode.IsDigit(v)) {
			fmt.Println("Wrong password")
			return
		}
	}
	fmt.Println("Ok")
}
