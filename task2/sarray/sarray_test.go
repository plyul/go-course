package sarray

import (
	"math/rand"
	"testing"
)

func TestInput(t *testing.T) {
	input := []int{10, 4, 3, 9, -5, -3}
	sa := New(10)
	sa.Input(input)

	expected := []int{4, 9, 10}
	if len(sa.array) != len(expected) {
		t.Errorf("Expected length is %d, got %d", len(expected), len(sa.array))
	}
	for i, v := range sa.array {
		if v != expected[i] {
			t.Errorf("Expect array %v, got %v", expected, sa.array)
		}
	}
}

func TestInsert(t *testing.T) {
	var testCases = []struct {
		arrayBefore, arrayAfter []int
		value                   int
	}{
		{[]int{10}, []int{10, 10}, 10},
		{[]int{}, []int{10}, 10},
		{[]int{1}, []int{1, 10}, 10},
		{[]int{20}, []int{10, 20}, 10},
		{[]int{1, 20}, []int{1, 10, 20}, 10},
		{[]int{1, 2, 4, 8, 16, 32, 64}, []int{1, 2, 4, 8, 10, 16, 32, 64}, 10},
	}
	for _, tt := range testCases {
		sortedArray := New(10)
		sortedArray.array = append(sortedArray.array, tt.arrayBefore...)
		sortedArray.Insert(tt.value)
		if len(sortedArray.array) != len(tt.arrayAfter) {
			t.Errorf("Expect new array lenght %d, got %d for source array %v, value %d",
				len(tt.arrayAfter), len(sortedArray.array), tt.arrayBefore, tt.value)
		}
		for i, v := range sortedArray.array {
			if v != tt.arrayAfter[i] {
				t.Errorf("Expect array %v, got %v", tt.arrayAfter, sortedArray.array)
			}
		}
	}
}

func TestDelete(t *testing.T) {
	var testCases = []struct {
		arrayBefore, arrayAfter []int
		value                   int
	}{
		{[]int{}, []int{}, 10},
		{[]int{20}, []int{20}, 10},
		{[]int{10}, []int{}, 10},
		{[]int{1, 10}, []int{1}, 10},
		{[]int{10, 20}, []int{20}, 10},
		{[]int{1, 20}, []int{1, 20}, 10},
		{[]int{10, 10, 10, 10, 10}, make([]int, 0, 1), 10},
		{[]int{1, 2, 4, 8, 10, 16, 32, 64}, []int{1, 2, 4, 8, 16, 32, 64}, 10},
	}
	for _, tt := range testCases {
		sortedArray := New(10)
		sortedArray.array = append(sortedArray.array, tt.arrayBefore...)
		sortedArray.Delete(tt.value)
		if len(sortedArray.array) != len(tt.arrayAfter) {
			t.Errorf("Expect new array lenght %d, got %d for source array %v, value %d",
				len(tt.arrayAfter), len(sortedArray.array), tt.arrayBefore, tt.value)
		}
		for i, v := range sortedArray.array {
			if v != tt.arrayAfter[i] {
				t.Errorf("Expect array %v, got %v", tt.arrayAfter, sortedArray.array)
			}
		}
	}
}

func TestGetIndex(t *testing.T) {
	var testCases = []struct {
		array       []int
		value       int
		insertIndex int
	}{
		{[]int{}, 10, 0},
		{[]int{1, 3, 5, 7, 9}, 4, 2},
		{[]int{1, 3, 5, 7, 9}, 10, 5},
		{[]int{1, 3, 5, 5, 5, 9}, 5, 2},
		{[]int{1, 3, 5, 7, 9}, 9, 4},
	}
	for _, tt := range testCases {
		gotIndex := getIndex(tt.array, tt.value)
		if gotIndex != tt.insertIndex {
			t.Errorf("Expect index %d for value %d in array %v, got %d", tt.insertIndex, tt.value, tt.array, gotIndex)
		}
	}
}

// go test -bench='Insert\d' .
// Evaluated complexity O(n^2)
func BenchmarkSortedArray_Insert100(b *testing.B)  { benchmarkSortedArray_Insert(100, b) }
func BenchmarkSortedArray_Insert1k(b *testing.B)   { benchmarkSortedArray_Insert(1000, b) }
func BenchmarkSortedArray_Insert10k(b *testing.B)  { benchmarkSortedArray_Insert(10000, b) }
func BenchmarkSortedArray_Insert20k(b *testing.B)  { benchmarkSortedArray_Insert(20000, b) }
func BenchmarkSortedArray_Insert30k(b *testing.B)  { benchmarkSortedArray_Insert(30000, b) }
func BenchmarkSortedArray_Insert40k(b *testing.B)  { benchmarkSortedArray_Insert(40000, b) }
func BenchmarkSortedArray_Insert50k(b *testing.B)  { benchmarkSortedArray_Insert(50000, b) }
func BenchmarkSortedArray_Insert60k(b *testing.B)  { benchmarkSortedArray_Insert(60000, b) }
func BenchmarkSortedArray_Insert70k(b *testing.B)  { benchmarkSortedArray_Insert(70000, b) }
func BenchmarkSortedArray_Insert80k(b *testing.B)  { benchmarkSortedArray_Insert(80000, b) }
func BenchmarkSortedArray_Insert90k(b *testing.B)  { benchmarkSortedArray_Insert(90000, b) }
func BenchmarkSortedArray_Insert100k(b *testing.B) { benchmarkSortedArray_Insert(100000, b) }
func BenchmarkSortedArray_Insert110k(b *testing.B) { benchmarkSortedArray_Insert(110000, b) }
func BenchmarkSortedArray_Insert120k(b *testing.B) { benchmarkSortedArray_Insert(120000, b) }
func BenchmarkSortedArray_Insert130k(b *testing.B) { benchmarkSortedArray_Insert(130000, b) }
func BenchmarkSortedArray_Insert140k(b *testing.B) { benchmarkSortedArray_Insert(140000, b) }
func BenchmarkSortedArray_Insert150k(b *testing.B) { benchmarkSortedArray_Insert(150000, b) }

func benchmarkSortedArray_Insert(n int, b *testing.B) {
	rand.Seed(1)
	sa := New(n)
	for i := 0; i < n; i++ {
		v := rand.Intn(1<<63 - 1)
		sa.Insert(v)
	}
}

// go test -bench='InsertOne' .
// Evaluated complexity O(n)
func BenchmarkSortedArray_InsertOne1k(b *testing.B)   { benchmarkSortedArray_InsertOne(1000, b) }
func BenchmarkSortedArray_InsertOne10k(b *testing.B)  { benchmarkSortedArray_InsertOne(10000, b) }
func BenchmarkSortedArray_InsertOne20k(b *testing.B)  { benchmarkSortedArray_InsertOne(20000, b) }
func BenchmarkSortedArray_InsertOne30k(b *testing.B)  { benchmarkSortedArray_InsertOne(30000, b) }
func BenchmarkSortedArray_InsertOne40k(b *testing.B)  { benchmarkSortedArray_InsertOne(40000, b) }
func BenchmarkSortedArray_InsertOne50k(b *testing.B)  { benchmarkSortedArray_InsertOne(50000, b) }
func BenchmarkSortedArray_InsertOne60k(b *testing.B)  { benchmarkSortedArray_InsertOne(60000, b) }
func BenchmarkSortedArray_InsertOne70k(b *testing.B)  { benchmarkSortedArray_InsertOne(70000, b) }
func BenchmarkSortedArray_InsertOne80k(b *testing.B)  { benchmarkSortedArray_InsertOne(80000, b) }
func BenchmarkSortedArray_InsertOne90k(b *testing.B)  { benchmarkSortedArray_InsertOne(90000, b) }
func BenchmarkSortedArray_InsertOne100k(b *testing.B) { benchmarkSortedArray_InsertOne(100000, b) }
func BenchmarkSortedArray_InsertOne110k(b *testing.B) { benchmarkSortedArray_InsertOne(110000, b) }
func BenchmarkSortedArray_InsertOne120k(b *testing.B) { benchmarkSortedArray_InsertOne(120000, b) }
func BenchmarkSortedArray_InsertOne130k(b *testing.B) { benchmarkSortedArray_InsertOne(130000, b) }
func BenchmarkSortedArray_InsertOne140k(b *testing.B) { benchmarkSortedArray_InsertOne(140000, b) }
func BenchmarkSortedArray_InsertOne150k(b *testing.B) { benchmarkSortedArray_InsertOne(150000, b) }

func benchmarkSortedArray_InsertOne(n int, b *testing.B) {
	rand.Seed(1)
	sa := New(n)
	for i := 0; i < n; i++ {
		v := rand.Intn(1<<63 - 1)
		sa.Insert(v)
	}
	v := rand.Intn(1<<63 - 1)
	b.ResetTimer()
	sa.Insert(v)
}

// go test -bench='DeleteOne' .
func BenchmarkSortedArray_DeleteOne1k(b *testing.B)   { benchmarkSortedArray_DeleteOne(1000, b) }
func BenchmarkSortedArray_DeleteOne10k(b *testing.B)  { benchmarkSortedArray_DeleteOne(10000, b) }
func BenchmarkSortedArray_DeleteOne20k(b *testing.B)  { benchmarkSortedArray_DeleteOne(20000, b) }
func BenchmarkSortedArray_DeleteOne30k(b *testing.B)  { benchmarkSortedArray_DeleteOne(30000, b) }
func BenchmarkSortedArray_DeleteOne40k(b *testing.B)  { benchmarkSortedArray_DeleteOne(40000, b) }
func BenchmarkSortedArray_DeleteOne50k(b *testing.B)  { benchmarkSortedArray_DeleteOne(50000, b) }
func BenchmarkSortedArray_DeleteOne60k(b *testing.B)  { benchmarkSortedArray_DeleteOne(60000, b) }
func BenchmarkSortedArray_DeleteOne70k(b *testing.B)  { benchmarkSortedArray_DeleteOne(70000, b) }
func BenchmarkSortedArray_DeleteOne80k(b *testing.B)  { benchmarkSortedArray_DeleteOne(80000, b) }
func BenchmarkSortedArray_DeleteOne90k(b *testing.B)  { benchmarkSortedArray_DeleteOne(90000, b) }
func BenchmarkSortedArray_DeleteOne100k(b *testing.B) { benchmarkSortedArray_DeleteOne(100000, b) }
func BenchmarkSortedArray_DeleteOne110k(b *testing.B) { benchmarkSortedArray_DeleteOne(110000, b) }
func BenchmarkSortedArray_DeleteOne120k(b *testing.B) { benchmarkSortedArray_DeleteOne(120000, b) }
func BenchmarkSortedArray_DeleteOne130k(b *testing.B) { benchmarkSortedArray_DeleteOne(130000, b) }
func BenchmarkSortedArray_DeleteOne140k(b *testing.B) { benchmarkSortedArray_DeleteOne(140000, b) }
func BenchmarkSortedArray_DeleteOne150k(b *testing.B) { benchmarkSortedArray_DeleteOne(150000, b) }

func benchmarkSortedArray_DeleteOne(n int, b *testing.B) {
	rand.Seed(1)
	sa := New(n)
	for i := 0; i < n; i++ {
		v := rand.Intn(1<<63 - 1)
		sa.Insert(v)
	}
	v := rand.Intn(1<<63 - 1)
	b.ResetTimer()
	sa.Delete(v)
}
