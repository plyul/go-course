package bufferedchan

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	input := make(chan string)
	bc := New(input, 2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go bc.Start(ctx)

	input <- "Uno"
	input <- "Dos"
	input <- "Tres"
	input <- "Cuatro"

	fmt.Printf("Got '%s' from output\n", <-bc.OutChannel())
	fmt.Printf("Got '%s' from output\n", <-bc.OutChannel())
	fmt.Printf("Got '%s' from output\n", <-bc.OutChannel())
	fmt.Printf("Got '%s' from output\n", <-bc.OutChannel())
}
