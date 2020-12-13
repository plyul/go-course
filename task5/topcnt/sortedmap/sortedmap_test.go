package sortedmap

import (
	"testing"
)

func TestSortedMap_IncrementAndDelete(t *testing.T) {
	sm := New(10)
	sm.Increment("first", 0, 1)
	sm.Increment("second", 0, 2)
	sm.Increment("third", 0, 3)
	sm.Increment("second", 0, 4)
	sm.Increment("first", 0, 5)
	sm.Increment("second", 0, 6)

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

	if sm.wordInsertIndices["first"] != 1 {
		t.Errorf("Failed at 'first' insert index")
	}
	if sm.wordInsertIndices["second"] != 2 {
		t.Errorf("Failed at 'first' insert index")
	}
	if sm.wordInsertIndices["third"] != 3 {
		t.Errorf("Failed at 'first' insert index")
	}

	sm.Delete("first")
	sm.Delete("second")
	sm.Delete("third")

	if len(sm.wordCounters) > 0 {
		t.Errorf("Failed at 'wordCounters' len after delete")
	}
	if len(sm.wordInsertIndices) > 0 {
		t.Errorf("Failed at 'wordInsertIndices' len after delete")
	}
	if sm.size != 0 {
		t.Errorf("Failed at size after delete")
	}
}
