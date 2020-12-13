package fanout

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestFanOut(t *testing.T) {
	input := make(chan string)
	defer close(input)
	fanout := New(input)

	out1 := fanout.NewOutChannel()
	out2 := fanout.NewOutChannel()

	go func() {
		input <- "Uno"
		input <- "Dos"
		input <- "Tres"
		input <- "Cuatro"
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fanout.Start(ctx)
	funIsOver := false
	var result []string
	for {
		select {
		case s, ok := <-out1:
			if !ok {
				funIsOver = true
				break
			}
			result = append(result, s)
		case s, ok := <-out2:
			if !ok {
				funIsOver = true
				break
			}
			result = append(result, s)
		}
		if funIsOver {
			break
		}
	}
	for i, r := range result {
		fmt.Printf("[%d] `%s`\n", i, r)
	}
	if len(result) != 8 {
		t.Errorf("Unexpected result len")
	}
}
