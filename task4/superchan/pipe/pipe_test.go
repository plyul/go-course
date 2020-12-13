package pipe

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	input := make(chan string)
	output := make(chan string)
	defer close(input)
	defer close(output)

	go New(input, output, func(s string) string {
		return strings.ReplaceAll(s, "a", "$")
	})

	inputData := []string{"Yavin", "Tatooine", "Kashyyyk", "Endor"}
	expectedOutput := []string{"Y$vin", "T$tooine", "K$shyyyk", "Endor"}
	for i, s := range inputData {
		input <- s
		o := <-output
		if expectedOutput[i] != o {
			t.Errorf("Error on index %d", i)
		}
	}
}
