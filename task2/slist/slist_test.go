package slist

import (
	"math/rand"
	"testing"
)

func TestSortedList_GetMin(t *testing.T) {
	var err error
	sortedList := New()
	_, err = sortedList.GetMin()
	if err == nil {
		t.Errorf("Non nil error on empty list")
	}
	for i := 88; i <= 888; i++ {
		sortedList.Insert(i)
	}
	expectedMin := 88
	gotMin, err := sortedList.GetMin()
	if gotMin != expectedMin {
		t.Errorf("GetMin expected %d, got %d", expectedMin, gotMin)
	}
}

func TestSortedList_GetMax(t *testing.T) {
	var err error
	sortedList := New()
	_, err = sortedList.GetMax()
	if err == nil {
		t.Errorf("Non nil error on empty list")
	}
	for i := 88; i <= 888; i++ {
		sortedList.Insert(i)
	}
	expectedMax := 888
	gotMax, err := sortedList.GetMax()
	if gotMax != expectedMax {
		t.Errorf("GetMax expected %d, got %d", expectedMax, gotMax)
	}
}

func TestSortedList_Insert(t *testing.T) {
	var testCases = []struct {
		insertedValues, sortedValues []int
	}{
		{[]int{}, []int{}},
		{[]int{10}, []int{10}},
		{[]int{1, 10, 20, 30}, []int{1, 10, 20, 30}},
		{[]int{30, 20, 10, 1}, []int{1, 10, 20, 30}},
		{[]int{10, 20, 1, 30}, []int{1, 10, 20, 30}},
	}
	for _, tt := range testCases {
		sortedList := New()
		for i := 0; i < len(tt.insertedValues); i++ {
			sortedList.Insert(tt.insertedValues[i])
		}
		v := sortedList.Values()
		if len(v) != len(tt.sortedValues) {
			t.Errorf("Expect new array lenght %d, got %d for source array %v",
				len(tt.sortedValues), len(sortedList.Values()), tt.insertedValues)
		}
		for i, v := range sortedList.Values() {
			if v != tt.sortedValues[i] {
				t.Errorf("Expect array %v, got %v", tt.sortedValues, sortedList.Values())
			}
		}
	}
}

func TestSortedList_Delete(t *testing.T) {
	var testCases = []struct {
		insertedValues, sortedValues []int
		value                        int
	}{
		{[]int{}, []int{}, 10},
		{[]int{20}, []int{20}, 10},
		{[]int{10}, []int{}, 10},
		{[]int{1, 10}, []int{1}, 10},
		{[]int{10, 20}, []int{20}, 10},
		{[]int{1, 20}, []int{1, 20}, 10},
		{[]int{10, 10, 20, 10, 10}, []int{20}, 10},
		{[]int{1, 2, 4, 8, 10, 16, 32, 64}, []int{1, 2, 4, 8, 16, 32, 64}, 10},
	}
	for _, tt := range testCases {
		sortedList := New()
		for i := 0; i < len(tt.insertedValues); i++ {
			sortedList.Insert(tt.insertedValues[i])
		}
		sortedList.Delete(tt.value)
		v := sortedList.Values()
		if len(v) != len(tt.sortedValues) {
			t.Errorf("Expect new array lenght %d, got %d for insertedValues %v, value %d",
				len(tt.sortedValues), len(v), tt.insertedValues, tt.value)
		}
		for i, va := range v {
			if i >= len(tt.sortedValues) {
				t.Errorf("Index out of range")
				break
			}
			if va != tt.sortedValues[i] {
				t.Errorf("Expect array %v, got %v", tt.sortedValues, v)
			}
		}
	}
}

// go test -run Length -coverprofile=coverage.out
// go tool cover -html=coverage.out
func TestSortedList_Length(t *testing.T) {
	type listOperation int
	const (
		opInsert listOperation = iota
		opDelete
	)
	var testSequence = []struct {
		value          int
		operation      listOperation
		expectedLength int
	}{
		{10, opInsert, 1},
		{20, opInsert, 2},
		{5, opInsert, 3},
		{15, opInsert, 4},
		{5, opDelete, 3},
		{20, opDelete, 2},
		{10, opDelete, 1},
	}
	sl := New()
	if sl.Length() != 0 {
		t.Error("New list has length not equal to zero")
	}
	for i, seq := range testSequence {
		if seq.operation == opInsert {
			sl.Insert(seq.value)
		} else {
			sl.Delete(seq.value)
		}
		if sl.Length() != seq.expectedLength {
			t.Errorf("Error in test sequence step %d", i)
		}
	}
}

func TestSortedList_IsEqual(t *testing.T) {
	var testCases = []struct {
		self, other []int
		isEqual     bool
	}{
		{[]int{}, []int{}, true},
		{[]int{10}, []int{}, false},
		{[]int{}, []int{10}, false},
		{[]int{10, 20}, []int{10}, false},
		{[]int{10}, []int{10, 20}, false},
		{[]int{10}, []int{10}, true},
		{[]int{10}, []int{20}, false},
	}
	for _, tt := range testCases {
		self := New()
		self.InsertValues(tt.other)
		other := New()
		other.InsertValues(tt.self)
		if self.IsEqual(other) != tt.isEqual {
			t.Errorf("Expect equality of %v and %v to be %v, but it is not", self, other, tt.isEqual)
		}
	}
}

// go test -bench='DeleteOne' -benchmem .
func BenchmarkSortedList_DeleteOne1k(b *testing.B)   { benchmarkSortedList_DeleteOne(1000, b) }
func BenchmarkSortedList_DeleteOne10k(b *testing.B)  { benchmarkSortedList_DeleteOne(10000, b) }
func BenchmarkSortedList_DeleteOne20k(b *testing.B)  { benchmarkSortedList_DeleteOne(20000, b) }
func BenchmarkSortedList_DeleteOne30k(b *testing.B)  { benchmarkSortedList_DeleteOne(30000, b) }
func BenchmarkSortedList_DeleteOne40k(b *testing.B)  { benchmarkSortedList_DeleteOne(40000, b) }
func BenchmarkSortedList_DeleteOne50k(b *testing.B)  { benchmarkSortedList_DeleteOne(50000, b) }
func BenchmarkSortedList_DeleteOne60k(b *testing.B)  { benchmarkSortedList_DeleteOne(60000, b) }
func BenchmarkSortedList_DeleteOne70k(b *testing.B)  { benchmarkSortedList_DeleteOne(70000, b) }
func BenchmarkSortedList_DeleteOne80k(b *testing.B)  { benchmarkSortedList_DeleteOne(80000, b) }
func BenchmarkSortedList_DeleteOne90k(b *testing.B)  { benchmarkSortedList_DeleteOne(90000, b) }
func BenchmarkSortedList_DeleteOne100k(b *testing.B) { benchmarkSortedList_DeleteOne(100000, b) }
func BenchmarkSortedList_DeleteOne110k(b *testing.B) { benchmarkSortedList_DeleteOne(110000, b) }
func BenchmarkSortedList_DeleteOne120k(b *testing.B) { benchmarkSortedList_DeleteOne(120000, b) }
func BenchmarkSortedList_DeleteOne130k(b *testing.B) { benchmarkSortedList_DeleteOne(130000, b) }
func BenchmarkSortedList_DeleteOne140k(b *testing.B) { benchmarkSortedList_DeleteOne(140000, b) }
func BenchmarkSortedList_DeleteOne150k(b *testing.B) { benchmarkSortedList_DeleteOne(150000, b) }

func benchmarkSortedList_DeleteOne(n int, b *testing.B) {
	rand.Seed(1)
	slist := New()
	for i := 0; i < n; i++ {
		v := rand.Intn(1<<63 - 1)
		slist.Insert(v)
	}
	v := rand.Intn(1<<63 - 1)
	b.ResetTimer()
	slist.Delete(v)
}
