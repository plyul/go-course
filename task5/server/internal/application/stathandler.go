package application

import (
	"fmt"
	"github.com/plyul/go-course/task5/topcnt/sortedmap"
	"github.com/plyul/go-course/task5/topcnt/textparser"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	multiplier               = 100000
	defaultSortedMapCapacity = 100
	minWordLen               = 3
)

func (a *Application) statHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("textHandler: %v", r)
	if r.Method != "" && r.Method != "GET" {
		a.defaultHandler(w, r)
		return
	}
	suri := strings.Split(r.RequestURI, "/")
	if len(suri) != 3 {
		a.defaultHandler(w, r)
		return
	}
	N, err := strconv.Atoi(suri[2])
	if err != nil {
		a.defaultHandler(w, r)
		return
	}
	countedWords := sortedmap.New(defaultSortedMapCapacity)
	a.text.Range(func(key, value interface{}) bool {
		p := textparser.New(strings.NewReader(value.(string)))
		bannedWords := make(map[string]struct{})
		var wd textparser.WordData
		for wd.Tag != textparser.EOF {
			wd = p.GetWord()
			wd.ChunkOffset = int64(key.(int)-1) * multiplier
			if wd.Tag == textparser.OnEdge {
				bannedWords[wd.Word] = struct{}{}
			}
			_, isBanned := bannedWords[wd.Word]
			if wd.Tag == textparser.Regular && len(wd.Word) >= minWordLen && !isBanned {
				countedWords.Increment(wd.Word, wd.ChunkOffset, wd.ChunkIdx)
			}
			if isBanned {
				countedWords.Delete(wd.Word)
			}
		}
		return true
	})
	topWords := countedWords.Top(N)
	for _, wd := range topWords {
		if _, err := fmt.Fprintf(w, "%s [%d] @%d\n", wd.Word, wd.Count, wd.InsertIndex); err != nil {
			log.Println("")
		}
	}
}
