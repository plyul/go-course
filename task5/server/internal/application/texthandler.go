package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Chapter struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

func (a *Application) textHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("textHandler: %v", r)
	if r.Method != "POST" {
		a.defaultHandler(w, r)
		return
	}
	chapter := Chapter{}
	err := json.NewDecoder(r.Body).Decode(&chapter)
	if err != nil {
		http.Error(w, fmt.Sprintf("Decode request failed: %v", err.Error()), http.StatusBadRequest)
		return
	}
	if chapter.Number < 1 {
		http.Error(w, "Wrong chapter number", http.StatusBadRequest)
		return
	}
	a.text.Store(chapter.Number, chapter.Text)
	http.Error(w, "Chapter accepted", http.StatusOK)
}
