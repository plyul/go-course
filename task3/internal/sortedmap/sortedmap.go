// Пакет sortedmap реализует специальное отображение,
// в котором хранятся счётчики слов с сохранением порядка их вставки
package sortedmap

import "sort"

type SortedMap struct {
	wordCounters      map[string]int // слово -> кол-во
	wordInsertIndices map[string]int // слово -> индекс вставки
	inIdx             int            // монотонно возрастающее значение индекса вставки
	size              int
}

func New(capacity int) *SortedMap {
	return &SortedMap{
		wordCounters:      make(map[string]int, capacity),
		wordInsertIndices: make(map[string]int, capacity),
		inIdx:             1,
		size:              0,
	}
}

// Increment увеличивает значение по ключу key на единицу
// Если ключ key отсутствует, то он добавляется со значением 1
func (sm *SortedMap) Increment(key string) {
	if _, ok := sm.wordCounters[key]; !ok {
		sm.size++
		sm.wordInsertIndices[key] = sm.inIdx
		sm.inIdx++
		sm.wordCounters[key] = 0
	}
	sm.wordCounters[key]++
}

// Delete удаляет ключ key, если он присутствует в отображении
func (sm *SortedMap) Delete(key string) {
	_, keyIsPresent := sm.wordCounters[key]
	if keyIsPresent {
		delete(sm.wordCounters, key)
		delete(sm.wordInsertIndices, key)
		sm.size--
	}
}

type WordData struct {
	Word        string
	Count       int
	InsertIndex int
}

// Top возвращает N слов с максимальным количеством вхождений и отсортированных в порядке,
// в котором они вставлялись.
func (sm *SortedMap) Top(N int) []WordData {
	var words = make([]WordData, 0, len(sm.wordCounters))
	for word, cnt := range sm.wordCounters {
		words = append(words, WordData{
			Word:        word,
			Count:       cnt,
			InsertIndex: sm.wordInsertIndices[word],
		})
	}
	sort.Slice(words, func(i, j int) bool {
		// Нам нужно сортировать по убыванию, а не возрастанию, поэтому условие `>`
		return words[i].Count > words[j].Count
	})

	if len(words) < N {
		N = len(words)
	}
	words = words[:N]
	sort.Slice(words, func(i, j int) bool {
		return words[i].InsertIndex < words[j].InsertIndex
	})
	return words
}
