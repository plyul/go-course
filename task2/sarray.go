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
	insertIndex := getSortedIndex(a.array, value)
	a.array = append(a.array, 0)
	copy(a.array[insertIndex+1:], a.array[insertIndex:])
	a.array[insertIndex] = value
}

// Delete value from array
func (a *SortedArray) Delete(value int) {
	if value < 0 {
		panic("SortedArray.Delete: negative value passed")
	}
	if len(a.array) == 0 {
		return
	}
	for i := 0; i < len(a.array)-1; i++ {
		if a.array[i] == value {
			copy(a.array[i:], a.array[i+1:])
			a.array = a.array[:len(a.array)-1]
			i--
		}
	}
	if len(a.array) > 0 && a.array[len(a.array)-1] == value {
		a.array = a.array[:len(a.array)-1]
	}
}

// getSortedtIndex returns index for array, inserting value in which will retain array sorted.
func getSortedIndex(array []int, value int) int {
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
