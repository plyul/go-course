package bufferedchan

import (
	"context"
	"fmt"
)

type BufferedChannel struct {
	input  chan string
	output chan string
}

func New(input chan string, size int) *BufferedChannel {
	bc := &BufferedChannel{
		input:  input,
		output: make(chan string, size),
	}
	return bc
}

func (bc *BufferedChannel) Start(ctx context.Context) {
	for {
		select {
		case s := <-bc.input:
			fmt.Printf("Got '%s' from input, buffering\n", s)
			select {
			case bc.output <- s:
				fmt.Printf("Success. Buffer: %v\n", len(bc.output))
			default:
				fmt.Printf("Fail. dropping '%s'\n", s)
			}
		case <-ctx.Done():
			fmt.Println("Stopping executing due cancel signal")
			close(bc.output)
			return
		}
	}
}

func (bc *BufferedChannel) OutChannel() <-chan string {
	return bc.output
}
