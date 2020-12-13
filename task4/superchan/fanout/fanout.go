package fanout

import (
	"context"
	"fmt"
	"sync"
)

type FanOut struct {
	sync.Mutex
	ctx      context.Context
	inChan   <-chan string
	outChans []chan<- string
}

func New(input <-chan string) *FanOut {
	return &FanOut{
		ctx:    context.Background(),
		inChan: input,
	}
}

func (fanout *FanOut) NewOutChannel() <-chan string {
	c := make(chan string)
	fanout.Lock()
	fanout.outChans = append(fanout.outChans, c)
	fanout.Unlock()
	return c
}

func (fanout *FanOut) Start(ctx context.Context) {
	fanout.ctx = ctx
	go func() {
		for {
			select {
			case s := <-fanout.inChan:
				fanout.Lock()
				for _, c := range fanout.outChans {
					c <- s
				}
				fanout.Unlock()
			case <-fanout.ctx.Done():
				fmt.Println("Closing output channels due cancel signal")
				fanout.Lock()
				for _, c := range fanout.outChans {
					close(c)
				}
				fanout.Unlock()
				return
			}
		}
	}()
}
