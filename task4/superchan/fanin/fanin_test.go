package fanin

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestFanIn(t *testing.T) {
	const numberOfValuesToSend = 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fanin := New(ctx)

	in1 := make(chan string)
	in2 := make(chan string)
	in3 := make(chan string)
	fanin.Add(in1)
	fanin.Add(in2)
	fanin.Add(in3)
	close(in3) // Закрытие канала исключает его из FanIn
	go func() {
		for i := 0; i < numberOfValuesToSend; i++ {
			in1 <- "channel1"
			in2 <- "channel2"
		}
	}()

	var result []string
	for s := range fanin.out { // Будет закрыт по таймауту контекста
		result = append(result, s)
	}
	if len(result) != numberOfValuesToSend*2 {
		t.Errorf("Unexpected number of values in result")
	}
	fmt.Println(result)
}
