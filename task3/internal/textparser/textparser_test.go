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
		tag WordTag
	}{
		{"chapter", OnEdge},
		{"one", OnEdge},
		{"news", OnEdge},
		{"of", Regular},
		{"a", Regular},
		{"wedding", OnEdge},
		{"scarlett", OnEdge},
		{"ohara", Regular},
		{"was", Regular},
		{"blahblah", Regular},
		{"twins", Regular},
		{"were", OnEdge},
		{"her", OnEdge},
		{"eyes", Regular},
		{"were", Regular},
		{"blahblah", Regular},
		{"and", Regular},
		{"gloves", OnEdge},
		{"on", OnEdge},
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
		wd := p.GetWord()
		if wd.Word != seq.word || wd.Tag != seq.tag {
			t.Errorf("Test sequence %d failed. Expected '%v', got '%s', %v", i, seq, wd.Word, wd.Tag)
		}
	}
}
