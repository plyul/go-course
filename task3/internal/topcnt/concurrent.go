package topcnt

import (
	"errors"
	"fmt"
	"go-course/task3/internal/chunkreader"
	"go-course/task3/internal/sortedmap"
	"go-course/task3/internal/textparser"
	"os"
	"sync"
)

const (
	channelBufferSize         = 100
	defaultSortedMapCapacity = 1500
)

func (app Application) computeConcurrent(f *os.File) error {
	chunkOffsets, err := chunkreader.ChunkOffsets(f, app.config.ChunkSizeBytes)
	if err != nil {
		return errors.New(fmt.Sprintf("computeConcurrent error: %v", err))
	}
	fmt.Println(chunkOffsets)

	var wg sync.WaitGroup
	numChunks := len(chunkOffsets) - 1
	// TODO: Сделать ограничение количества воркеров задаваемым параметром N
	wg.Add(numChunks) // Количество воркеров по количеству чанков
	wordChan := make(chan textparser.WordData, channelBufferSize) // все читатели чанков будут складывать слова в этот канал
	for i := 0; i < numChunks; i++ {
		go func(begin, end int64) { // запускаем параллельное/конкурентное чтение файла чанками и наливку слов в канал
			defer wg.Done()
			fmt.Printf("Processing chunk [%d-%d)\n", begin, end)
			cr := chunkreader.New(f, begin, end)
			p := textparser.New(cr)
			var wd textparser.WordData
			for wd.Tag != textparser.EOF {
				wd = p.GetWord()
				wd.ChunkOffset = begin
				wordChan <- wd
			}
			fmt.Printf("Chunk %d finished\n", begin)
		}(chunkOffsets[i], chunkOffsets[i+1])
	}

	topConcurrentIsDone := make(chan bool)
	var topWords []sortedmap.WordData
	go func() { // Запускаем чтение слов из канала wordChan и их подсчет
		topWords = top(app.config.NumTopWords, app.config.MinWordLen, wordChan)
		topConcurrentIsDone <- true
	}()
	wg.Wait() // ждём, пока все воркеры закончат работу
	close(wordChan) // закрываем канал, больше слов не будет
	<-topConcurrentIsDone // ждём, пока закончится подсчёт слов
	for _, wd := range topWords { // и выводим результат
		fmt.Printf("%s [%d] @%d\n", wd.Word, wd.Count, wd.InsertIndex)
	}
	return nil
}

// top получает слова из канала wordChan и возвращает N наиболее употребимых слов, соответствующих условиям:
// * Длина слова не меньше заданного minWordLen;
// * Слова с начала и конца предложения игнорируются и не участвуют в подсчёте;
// * Слова возвращаются в порядке, в котором они встретились в тексте.
func top(N int, minWordLen int, wordChan <-chan textparser.WordData) []sortedmap.WordData {
	countedWords := sortedmap.New(defaultSortedMapCapacity)
	bannedWords := make(map[string]struct{})
	for w := range wordChan {
		if w.Tag == textparser.OnEdge {
			bannedWords[w.Word] = struct{}{}
		}
		_, isBanned := bannedWords[w.Word]
		if w.Tag == textparser.Regular && len(w.Word) >= minWordLen && !isBanned {
			countedWords.Increment(w.Word, w.ChunkOffset, w.ChunkIdx)
		}
		if isBanned {
			countedWords.Delete(w.Word)
		}
	}
	return countedWords.Top(N)
}
