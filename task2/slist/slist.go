package slist

import "errors"

type node struct {
	value int
	next  *node
}

type SortedList struct {
	head *node
}

func New() SortedList {
	return SortedList{}
}

func (l *SortedList) Insert(v int) {
	if l.head == nil {
		l.head = &node{v, nil}
		return
	}
	var prev *node
	cur := l.head
	for cur.value < v {
		if cur.next == nil {
			cur.next = &node{v, nil}
			return
		}
		prev = cur
		cur = cur.next
	}
	if prev != nil {
		prev.next = &node{v, cur}
	} else {
		l.head = &node{v, cur}
	}
}

// Delete deletes all 'v' values from list
func (l *SortedList) Delete(v int) {
	for l.deleteOne(v) {
	}
}

// deleteOne tries to delete one value from list, returning true, if delete was successful (value was found in the list), else false
func (l *SortedList) deleteOne(v int) bool {
	if l.head == nil {
		return false
	}
	var prev *node
	cur := l.head
	for cur.value != v {
		prev = cur
		cur = cur.next
		if cur == nil {
			return false
		}
	}
	if prev != nil {
		prev.next = cur.next
	} else {
		l.head = cur.next
	}
	return true
}

func (l *SortedList) GetMax() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}
	cur := l.head
	for {
		if cur.next == nil {
			return cur.value, nil
		}
		cur = cur.next
	}
}

func (l *SortedList) GetMin() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}
	return l.head.value, nil
}

// Values returns all values from list as slice of ints
func (l *SortedList) Values() []int {
	var values []int
	if l.head == nil {
		return nil
	}
	cur := l.head
	for {
		values = append(values, cur.value)
		cur = cur.next
		if cur == nil {
			break
		}
	}
	return values
}
