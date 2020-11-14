package sarray

import (
	"math/rand"
	"testing"
)

func TestInput(t *testing.T) {
	input := []int{10,4,3,9,-5,-3}
	sa := New(10)
	sa.Input(input)

	expected := []int{4,9,10}
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

func TestGetInsertIndex(t *testing.T) {
	var testCases = []struct {
		array       []int
		value       int
		insertIndex int
	}{
		{[]int{}, 10, 0},
		{[]int{1, 3, 5, 7, 9}, 4, 2},
		{[]int{1, 3, 5, 7, 9}, 10, 5},
	}
	for _, tt := range testCases {
		gotIndex := getInsertIndex(tt.array, tt.value)
		if gotIndex != tt.insertIndex {
			t.Errorf("Expect index %d for value %d in array %v, got %d", tt.insertIndex, tt.value, tt.array, gotIndex)
		}
	}
}

func BenchmarkSortedArray_Insert100(b *testing.B) {	benchmarkSortedArray_Insert(100, b) }
func BenchmarkSortedArray_Insert1k(b *testing.B) {	benchmarkSortedArray_Insert(1000, b) }
func BenchmarkSortedArray_Insert10k(b *testing.B) {	benchmarkSortedArray_Insert(10000, b) }
func BenchmarkSortedArray_Insert100k(b *testing.B) { benchmarkSortedArray_Insert(100000, b) }
func BenchmarkSortedArray_Insert110(b *testing.B) {	benchmarkSortedArray_Insert(110000, b) }
func BenchmarkSortedArray_Insert120(b *testing.B) {	benchmarkSortedArray_Insert(120000, b) }
func BenchmarkSortedArray_Insert130(b *testing.B) {	benchmarkSortedArray_Insert(130000, b) }
func BenchmarkSortedArray_Insert140(b *testing.B) {	benchmarkSortedArray_Insert(140000, b) }

func benchmarkSortedArray_Insert(n int, b *testing.B) {
	rand.Seed(1)
	sa := New(n)
	v := rand.Intn(65536)
	b.ResetTimer()
	for i := 0; i < n; i++ {
		sa.Insert(v)
	}

}
