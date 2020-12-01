package topcnt

import (
	"errors"
	"fmt"
	"go-course/task3/internal/textparser"
	"go-course/task3/internal/wordscounter"
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
	f, err := os.Open(app.config.InputFileName)
	defer func() {
		if f != nil {
			if err := f.Close(); err != nil {
				fmt.Printf("error closing file %s: %v\n", app.config.InputFileName, err)
			}
		}
	}()
	if err != nil {
		return errors.New(fmt.Sprintf("error opening file: %v", err))
	}
	p := textparser.New(f)
	wc := wordscounter.New(p)
	output := wc.Top(app.config.NumTopWords, app.config.MinWordLen)
	for _, wd := range output {
		fmt.Printf("%s [%d] @%d\n", wd.Word, wd.Count, wd.InsertIndex)
	}
	return nil
}
