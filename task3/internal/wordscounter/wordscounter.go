// Пакет wordscounter обрабатывает набор слов.
// Слова текста предоставляются интерфейсом WordsProvider
package wordscounter

import (
	"go-course/task3/internal/sortedmap"
	"go-course/task3/internal/wordsprovider"
)

const (
	defaultSortedMapCapacity = 1500
)

type WordsCounter struct {
	wp wordsprovider.WordsProvider
}

func New(wp wordsprovider.WordsProvider) *WordsCounter {
	return &WordsCounter{
		wp: wp,
	}
}

// Top получает слова от WordsProvider и возвращает N наиболее употребимых слов, соответствующих условиям:
// * Длина слова не меньше заданного minWordLen;
// * Слова с начала и конца предложения игнорируются и не участвуют в подсчёте;
// * Слова возвращаются в порядке, в котором они встретились в тексте.
func (wf WordsCounter) Top(N int, minWordLen int) []sortedmap.WordData {
	countedWords := sortedmap.New(defaultSortedMapCapacity)
	bannedWords := make(map[string]bool)
	for {
		word, tag := wf.wp.GetWord()
		if tag == wordsprovider.OnEdge {
			bannedWords[word] = true
		}
		_, isBanned := bannedWords[word]
		if tag == wordsprovider.Regular && len(word) >= minWordLen && !isBanned {
			countedWords.Increment(word)
		}
		if isBanned {
			countedWords.Delete(word)
		}
		if tag == wordsprovider.EOF {
			break
		}
	}
	return countedWords.Top(N)
}
