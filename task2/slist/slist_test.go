package slist

import (
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
