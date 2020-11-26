package sortedmap

import (
	"testing"
)

func TestSortedMap_IncrementAndDelete(t *testing.T) {
	sm := New(10)
	sm.Increment("first")
	sm.Increment("second")
	sm.Increment("third")
	sm.Increment("second")
	sm.Increment("first")
	sm.Increment("second")

	if sm.wordCounters["first"] != 2 {
		t.Errorf("Failed at 'first' counter")
	}
	if sm.wordCounters["second"] != 3 {
		t.Errorf("Failed at 'second' counter")
	}
	if sm.wordCounters["third"] != 1 {
		t.Errorf("Failed at 'third' counter")
	}

	if sm.size != 3 {
		t.Errorf("Failed at size after insert")
	}

	sm.Delete("first")
	sm.Delete("second")
	sm.Delete("third")

	if len(sm.wordCounters) > 0 {
		t.Errorf("Failed at 'wordCounters' len after delete")
	}
	if len(sm.wordInsertIndices) > 0 {
		t.Errorf("Failed at 'wordCounters' len after delete")
	}
	if sm.size != 0 {
		t.Errorf("Failed at size after delete")
	}
}
