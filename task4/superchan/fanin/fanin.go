package fanin

import (
	"context"
	"fmt"
	"sync"
)

type FanIn struct {
	sync.Mutex
	ctx context.Context
	out chan string
}

func New(ctx context.Context) *FanIn {
	f := &FanIn{
		ctx: ctx,
		out: make(chan string),
	}
	go func(fanin *FanIn) {
		<-fanin.ctx.Done()
		fmt.Println("Closing output channel due cancel signal")
		fanin.Lock()
		close(fanin.out)
		fanin.out = nil
		fanin.Unlock()
	}(f)
	return f
}

func (fanin *FanIn) Add(c <-chan string) {
	go func(c <-chan string) {
		fmt.Printf("Starting reader for channel %v\n", c)
		for v := range c {
			fanin.Lock()
			if fanin.out == nil { // выходной канал был закрыт, поэтому заканчиваем работу
				fmt.Printf("Output channel is closed. Stopping reader for channel %v\n", c)
				fanin.Unlock()
				return
			}
			fanin.out <- v
			fanin.Unlock()
		}
		fmt.Printf("Stopping reader for channel %v because it is closed\n", c)
	}(c)
}
