package application

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

type Application struct {
	server       *http.Server
	Finished     chan struct{}
	shutdownOnce sync.Once
	text         sync.Map
}

const (
	serverTimeout     = 60 * time.Second
	pprofAddr         = ":6060"
	pprofCPUFileName  = "cpu.prof"
	pprofMemFileName  = "mem.prof"
	pprofSampleRateHz = 5000
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
		Finished: make(chan struct{}),
	}
	appMux.HandleFunc("/text/", app.textHandler)
	appMux.HandleFunc("/stat/", app.statHandler)
	appMux.HandleFunc("/stop/", app.stopHandler)
	appMux.HandleFunc("/", app.defaultHandler)
	return app
}

func (a *Application) Run() error {
	go func() { // Запуск горутины для контроля сигнала SIGINT. При получении сигнала инициирует завершение сервера
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Println("Server shutdown initiated via SIGINT")
		a.initShutdown()
	}()
	go func() { // Запуск HTTP-сервера подсистемы профайлинга
		log.Println(http.ListenAndServe(pprofAddr, nil))
	}()
	_, debug := os.LookupEnv("DEBUG")
	if debug {
		f, err := os.Create(pprofCPUFileName)
		if err != nil {
			log.Panic("could not create CPU profile: ", err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Println("could not close CPU profile file: ", err)
			}
			log.Println("CPU profile file closed successfully")
		}()
		runtime.SetCPUProfileRate(pprofSampleRateHz) // см. https://stackoverflow.com/questions/30871691/cant-get-golang-pprof-working/31366860#31366860
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Panic("could not start CPU profile: ", err)
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
	<-a.Finished
	return srverr
}

func (a *Application) defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("defaultHandler: %v", r)
	http.Error(w, "No such URL", http.StatusNotFound)
}
