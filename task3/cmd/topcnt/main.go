package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"go-course/task3/internal/cfg"
	"go-course/task3/internal/textparser"
	"go-course/task3/internal/wordscounter"
	"os"
)

func main() {
	c := configure()
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

func configure() *cfg.ConfigStruct {
	const (
		minWordLengthDefault = 4
		numTopWordsDefault   = 10
	)
	var c cfg.ConfigStruct
	var needHelp bool
	pflag.BoolVar(&needHelp, "help", false, "Show available configuration options")
	pflag.StringVar(&c.InputFileName, "in", "", "Input text file name")
	pflag.IntVar(&c.MinWordLen, "min-len", minWordLengthDefault, "Minimal word length to count")
	pflag.IntVar(&c.NumTopWords, "num-top", numTopWordsDefault, "Number of top words to output")
	pflag.Parse()
	if needHelp {
		pflag.Usage()
		os.Exit(0)
	}
	if c.InputFileName == "" {
		fmt.Println("Mandatory argument --in is missing")
		os.Exit(1)
	}
	return &c
}
