package slist

import "errors"

type node struct {
	value int
	next  *node
}

type SortedList struct {
	head *node
	len  int
}

func New() SortedList {
	return SortedList{}
}

func (l *SortedList) Insert(v int) {
	if l.head == nil {
		l.head = &node{v, nil}
		l.len++
		return
	}
	var prev *node
	cur := l.head
	for cur.value < v {
		if cur.next == nil {
			cur.next = &node{v, nil}
			l.len++
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
	l.len++
}

func (l *SortedList) InsertValues(values []int) {
	for _, v := range values {
		l.Insert(v)
	}
}

func (l *SortedList) Delete(v int) {
	var prev *node
	cur := l.head
	for cur != nil {
		if cur.value == v {
			if prev != nil {
				prev.next = cur.next
				cur = cur.next
			} else {
				l.head = cur.next
				cur = l.head
			}
			l.len--
		} else {
			prev = cur
			cur = cur.next
		}
	}
}

func (l *SortedList) Length() int {
	return l.len
}

func (l *SortedList) IsEqual(other SortedList) bool {
	if l.Length() != other.Length() {
		return false
	}
	iterSelf := l.head
	iterOther := other.head
	for iterSelf != nil {
		if iterOther.value != iterSelf.value {
			return false
		}
		iterSelf = iterSelf.next
		iterOther = iterOther.next
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
