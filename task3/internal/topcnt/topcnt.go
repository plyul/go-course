package topcnt

import (
	"fmt"
	"os"
)

type Application struct {
	config *Configuration
}

func New() *Application {
	return &Application{
		config: configure(),
	}
}

func (app Application) Run() error {
	var err error
	f, err := os.Open(app.config.InputFileName)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	defer func() {
		if f != nil {
			if err := f.Close(); err != nil {
				fmt.Printf("error closing file %s: %v\n", app.config.InputFileName, err)
			}
		}
	}()
	err = app.computeConcurrent(f)
	return err
}
