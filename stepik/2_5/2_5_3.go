package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var err error
	stdinReader := bufio.NewReader(os.Stdin)
	X, err := stdinReader.ReadString('\n')
	S, err := stdinReader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading strings: %s", err.Error())
		return
	}
	X = strings.TrimSpace(X)
	S = strings.TrimSpace(S)
	fmt.Println(strings.Index(X, S))
}
