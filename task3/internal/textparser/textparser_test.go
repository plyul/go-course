package textparser

import (
	"fmt"
	"os"
	"testing"
)

// go test -run TestTextParser_Scan -coverprofile=coverage.out
// go tool cover -html=coverage.out
func TestTextParser_ScanScarlett(t *testing.T) {
	var testSequence = []struct {
		word   string
		status WordTag
	}{
		{"Chapter", OnEdge},
		{"one", OnEdge},
		{"News", OnEdge},
		{"of", Regular},
		{"a", Regular},
		{"Wedding", OnEdge},
		{"Scarlett", OnEdge},
		{"O'Hara", Regular},
		{"was", Regular},
		{"blah-blah", Regular},
		{"twins", Regular},
		{"were", OnEdge},
		{"Her", OnEdge},
		{"eyes", Regular},
		{"were", Regular},
		{"blah-blah", Regular},
		{"and", Regular},
		{"gloves", OnEdge},
		{"On", OnEdge},
		{"that", Regular},
		{"bright", OnEdge},
		{"", EOF},
	}
	f, err := os.Open("testdata/test_cut.txt")
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("Error closing file %s\n", f.Name())
		}
	}()
	if err != nil {
		t.Errorf("Error opening file: %v", err)
		return
	}
	p := New(f)
	for i, seq := range testSequence {
		w, s := p.GetWord()
		if w != seq.word || s != seq.status {
			t.Errorf("Test sequence %d failed. Expected '%v', got '%s', %v", i, seq, w, s)
		}
	}
}
