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
	a.initShutdown()
	log.Println("Server shutdown initiated via HTTP")
	http.Error(w, "Command accepted\n", http.StatusOK)
}

func (a *Application) initShutdown() {
	go func() {
		a.shutdownOnce.Do(func() {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			log.Printf("initShutdown: a.server.Shutdown error: %v", a.server.Shutdown(ctx))
			close(a.Finished)
		})
	}()
}
