package textparser

import (
	"fmt"
	"go-course/task3/internal/wordsprovider"
	"os"
	"testing"
)

// go test -run TestTextParser_Scan -coverprofile=coverage.out
// go tool cover -html=coverage.out
func TestTextParser_ScanScarlett(t *testing.T) {
	var testSequence = []struct {
		word   string
		status wordsprovider.WordTag
	}{
		{"Chapter", wordsprovider.OnEdge},
		{"one", wordsprovider.OnEdge},
		{"News", wordsprovider.OnEdge},
		{"of", wordsprovider.Regular},
		{"a", wordsprovider.Regular},
		{"Wedding", wordsprovider.OnEdge},
		{"Scarlett", wordsprovider.OnEdge},
		{"O'Hara", wordsprovider.Regular},
		{"was", wordsprovider.Regular},
		{"blah-blah", wordsprovider.Regular},
		{"twins", wordsprovider.Regular},
		{"were", wordsprovider.OnEdge},
		{"Her", wordsprovider.OnEdge},
		{"eyes", wordsprovider.Regular},
		{"were", wordsprovider.Regular},
		{"blah-blah", wordsprovider.Regular},
		{"and", wordsprovider.Regular},
		{"gloves", wordsprovider.OnEdge},
		{"On", wordsprovider.OnEdge},
		{"that", wordsprovider.Regular},
		{"bright", wordsprovider.OnEdge},
		{"", wordsprovider.EOF},
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
