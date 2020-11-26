package wordscounter

import (
	"go-course/task3/internal/textparser"
	"strings"
	"testing"
)

const (
	minWordLength = 4
	numTopWords   = 10
)

func TestSortedMap_Top(t *testing.T) {
	input := "I like this very much.\nDo you like this?"
	p := textparser.New(strings.NewReader(input))
	wc := New(*p, minWordLength)
	output := wc.Top(numTopWords)
	if len(output) != 2 {
		t.Errorf("Unexpected length of output slice")
	}
	if output[0].Word != "like" && output[1].Word != "very" {
		t.Errorf("Unexpected words in output slice")
	}
	if output[0].InsertIndex != 1 && output[1].InsertIndex != 3 {
		t.Errorf("Unexpected insert index in output slice")
	}
}
