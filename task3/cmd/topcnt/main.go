package main

import (
	"fmt"
	"go-course/task3/internal/topcnt"
)

func main() {
	app := topcnt.New()
	if err := app.Run(); err != nil {
		fmt.Printf("Error executing program: %v\n", err)
	}
}
