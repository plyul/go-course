package main

import (
	"fmt"
	"go-course/task3/internal/config"
	"go-course/task3/internal/textparser"
	"go-course/task3/internal/wordscounter"
	"os"
)

func main() {
	c := config.Get()
	f, err := os.Open(c.InputFileName)
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("Error closing file %s: %v\n", f.Name(), err)
		}
	}()
	if err != nil {
		fmt.Printf("Error opening file %s: %v", f.Name(), err)
		return
	}
	p := textparser.New(f)
	wc := wordscounter.New(p)
	output := wc.Top(c.NumTopWords, c.MinWordLen)
	for _, wd := range output {
		fmt.Printf("%s [%d] @%d\n", wd.Word, wd.Count, wd.InsertIndex)
	}
}
