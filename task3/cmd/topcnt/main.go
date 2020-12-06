package main

import (
	"fmt"
	"go-course/task3/internal/topcnt"
	"time"
)

func main() {
	app := topcnt.New()
	st := time.Now()
	if err := app.Run(); err != nil {
		fmt.Printf("Error executing program: %v\n", err)
	}
	et := time.Since(st)
	fmt.Printf("Time elapsed: %v", et.String())
}
