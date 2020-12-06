// Пакет sortedmap реализует специальное отображение,
// в котором хранятся счётчики слов с сохранением порядка их появления в тексте
package sortedmap

import "sort"

type SortedMap struct {
	wordCounters         map[string]int   // слово -> кол-во
	wordInsertIndices    map[string]int64 // слово -> индекс вставки
	size                 int
}

func New(capacity int) *SortedMap {
	return &SortedMap{
		wordCounters:         make(map[string]int, capacity),
		wordInsertIndices:    make(map[string]int64, capacity),
		size:                 0,
	}
}

// Increment увеличивает значение по ключу key на единицу
// Если ключ key отсутствует, то он добавляется со значением 1
func (sm *SortedMap) Increment(key string, chunkOff int64, chunkIdx int) {
	uIdx := chunkOff + int64(chunkIdx)
	if _, ok := sm.wordCounters[key]; !ok {
		sm.wordCounters[key] = 0
		sm.size++
		sm.wordInsertIndices[key] = uIdx
		}
	sm.wordCounters[key]++
	if sm.wordInsertIndices[key] > uIdx {
		sm.wordInsertIndices[key] = uIdx
	}
}

// Delete удаляет ключ key, если он присутствует в отображении
func (sm *SortedMap) Delete(key string) {
	if _, ok := sm.wordCounters[key]; ok {
		delete(sm.wordCounters, key)
		delete(sm.wordInsertIndices, key)
		sm.size--
	}
}

type WordData struct {
	Word        string
	Count       int
	InsertIndex int64
}

// Top возвращает N слов с максимальным количеством вхождений и отсортированных в порядке,
// соответствующему индексу вставки.
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
