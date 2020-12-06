package chunkreader

import (
	"io"
	"os"
	"testing"
)

func TestChunkReader_Read(t *testing.T) {
	f, _ := os.Open("testdata/test.txt")
	defer func() {
		_ = f.Close()
	}()
	cr := New(f, 5, 15)
	var buf = make([]byte, 3)
	n, err := cr.Read(buf)
	if n != 3 || err != nil || string(buf[0:n]) != "ing" {
		t.Error("First read failed")
	}
	n, err = cr.Read(buf)
	if n != 3 || err != nil || string(buf[0:n]) != " st" {
		t.Error("Second read failed")
	}
	n, err = cr.Read(buf)
	if n != 3 || err != nil || string(buf[0:n]) != "ora" {
		t.Error("Third read failed")
	}
	n, err = cr.Read(buf)
	if n != 1 || err != nil || string(buf[0:n]) != "g" {
		t.Error("Fourth read failed")
	}
	n, err = cr.Read(buf)
	if n != 0 || err != io.EOF {
		t.Error("Fifth read failed")
	}
}
