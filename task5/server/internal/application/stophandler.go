package application

import (
	"context"
	"log"
	"net/http"
	"time"
)

func (a *Application) stopHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		a.defaultHandler(w, r)
		return
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		log.Printf("stopHandler:server.Shutdown error: %v", a.server.Shutdown(ctx))
		a.Finished <- true
	}()
	http.Error(w, "Server shutdown initiated!\n", http.StatusOK)
}
