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
	channelBufferSize        = 100
	defaultSortedMapCapacity = 1500
)

func (app Application) computeConcurrent(f *os.File) error {
	chunkOffsets, err := chunkreader.ChunkOffsets(f, app.config.ChunkSizeBytes) // нарезка файла на границы чанков
	if err != nil {
		return errors.New(fmt.Sprintf("computeConcurrent error: %v", err))
	}
	numChunks := len(chunkOffsets) - 1 // количество чанков в файле
	numWorkers := app.config.NumWorkers
	if numChunks < numWorkers {
		numWorkers = numChunks
	}

	// запускаем заданное число воркеров для конкурентного чтения файла чанками и наливку слов в канал wordChan
	type jobStruct struct {
		begin int64
		end   int64
	}
	jobChan := make(chan jobStruct, numChunks)                    // воркеры будут брать задачи из этого канала
	wordChan := make(chan textparser.WordData, channelBufferSize) // воркеры будут складывать слова в этот канал
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go func(w int) {
			defer wg.Done()
			for job := range jobChan {
				//fmt.Printf("Worker %d start processing chunk [%d-%d)\n", w, job.begin, job.end)
				cr := chunkreader.New(f, job.begin, job.end)
				p := textparser.New(cr)
				var wd textparser.WordData
				for wd.Tag != textparser.EOF {
					wd = p.GetWord()
					wd.ChunkOffset = job.begin
					wordChan <- wd
				}
				//fmt.Printf("Worker %d finished\n", w)
			}
		}(w)
	}

	// наливаем задачи в канал jobChan
	for i := 0; i < numChunks; i++ {
		job := jobStruct{
			begin: chunkOffsets[i],
			end:   chunkOffsets[i+1],
		}
		jobChan <- job
	}
	close(jobChan)

	// Запускаем чтение слов из канала wordChan и их подсчёт
	topIsDone := make(chan bool)
	var topWords []sortedmap.WordData
	go func() {
		topWords = top(app.config.NumTopWords, app.config.MinWordLen, wordChan)
		topIsDone <- true
	}()
	wg.Wait()                     // ждём, пока все воркеры закончат наливку слов в канал
	close(wordChan)               // закрываем канал, больше слов не будет
	<-topIsDone                   // ждём, пока закончится подсчёт слов
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
