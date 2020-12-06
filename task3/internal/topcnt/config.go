package topcnt

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"runtime"
)

const (
	minWordLengthDefault  = 4
	numTopWordsDefault    = 10
	chunkSizeBytesDefault = 10000
)

type Configuration struct {
	InputFileName  string
	MinWordLen     int
	NumTopWords    int
	ChunkSizeBytes int64
	NumWorkers     int
}

func configure() *Configuration {
	var c Configuration
	var needHelp bool
	pflag.BoolVar(&needHelp, "help", false, "Show available configuration options")
	pflag.StringVar(&c.InputFileName, "in", "", "Input text file name")
	pflag.IntVar(&c.MinWordLen, "min-len", minWordLengthDefault, "Minimal word length to count")
	pflag.IntVar(&c.NumTopWords, "num-top", numTopWordsDefault, "Number of top words to output")
	pflag.Int64Var(&c.ChunkSizeBytes, "chunk-size", chunkSizeBytesDefault, "Size of chunk for parallel computation, bytes")
	pflag.IntVar(&c.NumWorkers, "num-workers", runtime.NumCPU(), "Number of workers")
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
