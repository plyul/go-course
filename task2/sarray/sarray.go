package sarray

type SortedArray struct {
	array []int
}

// New returns freshly initialised SortedArray with given capacity.
func New(capacity int) SortedArray {
	return SortedArray{make([]int, 0, capacity)}
}

// Input processes array of control values.
// Positive value will be inserted in array, negative value will result in deleting respective positive value from array (if exists).
func (a *SortedArray) Input(values []int) {
	for _, v := range values {
		if v >= 0 {
			a.Insert(v)
		} else {
			a.Delete(-v)
		}
	}
}

// Insert value to array
func (a *SortedArray) Insert(value int) {
	if value < 0 {
		panic("SortedArray.Insert: negative value passed")
	}
	insertIndex := getIndex(a.array, value)
	a.array = append(a.array, 0)
	copy(a.array[insertIndex+1:], a.array[insertIndex:])
	a.array[insertIndex] = value
}

// Delete value from array
func (a *SortedArray) Delete(value int) {
	if value < 0 {
		panic("SortedArray.Delete: negative value passed")
	}
	for {
		if len(a.array) == 0 {
			return
		}
		i := getIndex(a.array, value)
		if i == len(a.array) {
			return
		}
		if a.array[i] == value {
			copy(a.array[i:], a.array[i+1:])
			a.array = a.array[:len(a.array)-1]
		} else {
			return
		}
	}
}

func (a *SortedArray) IsEqual(s SortedArray) bool {
	if len(a.array) != len(s.array) {
		return false
	}
	for i, v := range a.array {
		if s.array[i] != v {
			return false
		}
	}
	return true
}

// getIndex returns index of given value in array.
// If value is not present in array, inserting it at given position will retain array sorted.
func getIndex(array []int, value int) int {
	var low int = 0
	var high = len(array)
	for low < high {
		mid := (low + high) >> 1
		if array[mid] < value {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
