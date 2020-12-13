package application

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

type Application struct {
	server   *http.Server
	Finished chan bool
	text     sync.Map
}

const (
	serverTimeout    = 60 * time.Second
	pprofAddr        = ":6060"
	pprofCPUFileName = "cpu.prof"
	pprofMemFileName = "mem.prof"
)

func New(addr string) *Application {
	appMux := http.NewServeMux()
	appSrv := http.Server{
		Addr:         addr,
		Handler:      appMux,
		ReadTimeout:  serverTimeout,
		WriteTimeout: serverTimeout,
	}
	app := &Application{
		server:   &appSrv,
	}
	appMux.HandleFunc("/text/", app.textHandler)
	appMux.HandleFunc("/stat/", app.statHandler)
	appMux.HandleFunc("/stop/", app.stopHandler)
	appMux.HandleFunc("/", app.defaultHandler)
	return app
}

func (a *Application) Run() error {
	go func() {
		log.Println(http.ListenAndServe(pprofAddr, nil))
	}()
	_, debug := os.LookupEnv("DEBUG")
	if debug {
		f, err := os.Create(pprofCPUFileName)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Println("could not close CPU profile file: ", err)
			}
		}()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
		log.Println("CPU profiling started")
	}

	srverr := a.server.ListenAndServe()

	if debug {
		f, err := os.Create(pprofMemFileName)
		if err != nil {
			log.Println("could not create memory profile: ", err)
			return srverr
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Println("could not close memory profile file: ", err)
			}
		}()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Println("could not write memory profile: ", err)
			return srverr
		}
		log.Println("Memory profile written successfully")
	}
	return srverr
}

func (a *Application) defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("defaultHandler: %v", r)
	http.Error(w, "No such URL", http.StatusNotFound)
}
