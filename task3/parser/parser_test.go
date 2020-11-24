package parser

import (
	"os"
	"testing"
)

// go test -run TestTextParser_Scan -coverprofile=coverage.out
// go tool cover -html=coverage.out
func TestTextParser_Scan(t *testing.T) {
	var testSequence = []struct {
		value string
	}{
		{"O'Hara"}, {"beautiful"}, {"realize"}, {"this"}, {"when"}, {"caught"},
		{"charm"}, {"Tarleton"}, {"twins"}, {"eyes"}, {"were"}, {"green"},
		{"skin"}, {"that"}, {"soft"}, {"white"}, {"skin"}, {"which"},
		{"Southern"}, {"women"}, {"valued"}, {"highly"}, {"covered"}, {"carefully"},
		{"from"}, {"Georgia"}, {"with"}, {"hats"}, {"that"},
	}

	f, err := os.Open("test_cut.txt")
	defer f.Close()
	if err != nil {
		t.Errorf("Error opening file: %v", err)
		return
	}
	p := New(f)
	for _, seq := range testSequence {
		w := p.Scan()
		if w != seq.value {
			t.Errorf("Test sequence 1 failed. Expected '%s', got '%s'", seq.value, w)
		}
	}
	w := p.Scan()
	if w != "" {
		t.Errorf("EOF case failed")
	}
}
