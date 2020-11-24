package main

import (
	"fmt"
	"go-course/task3/parser"
	"os"
)

func main() {
	wordFreqs := make(map[string]int)
	wordAppearOrder := make([]string, 0, 1500)

	args := os.Args
	if len(args) < 2 {
		fmt.Print("Usage: \n\n")
		fmt.Printf("%s <text_file>\n", args[0])
		return
	}
	f, err := os.Open(args[1])
	defer f.Close()
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	textParser := parser.New(f)
	var w string
	for {
		w = textParser.Scan()
		if w != "" {
			if _, ok := wordFreqs[w]; !ok {
				wordAppearOrder = append(wordAppearOrder, w)
			}
			wordFreqs[w]++
		} else {
			break
		}
	}

	// TODO: Сделать более эффективно
	var topTen [10]string
	topTenMap := map[string]int{}
	for curTopPos := 0; curTopPos < 10; curTopPos++ {
		for i := range wordAppearOrder {
			word := wordAppearOrder[i]
			if wordFreqs[topTen[curTopPos]] < wordFreqs[word] {
				topTen[curTopPos] = word
			}
		}
		topTenMap[topTen[curTopPos]] = wordFreqs[topTen[curTopPos]]
		delete(wordFreqs, topTen[curTopPos])
	}
	for _, word := range wordAppearOrder{
		for _, topWord := range topTen {
			if word == topWord {
				fmt.Println(word, topTenMap[word])
			}
		}
	}
}
