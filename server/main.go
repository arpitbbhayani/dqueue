package server

import (
	"net/http"
	"sync"

	"github.com/arpitbbhayani/dqueue/dqueue"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func runHttpServer(wg *sync.WaitGroup) {
	defer wg.Done()

	r := mux.NewRouter()
	r.HandleFunc("/", VersionHandler)
	r.HandleFunc("/msg", MessagesHandler)
	http.Handle("/", r)

	logrus.Info("Dqueue started on port 4000")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		panic(err)
	}
}

func Run() {
	var wg sync.WaitGroup
	_ = dqueue.GetInstance()

	wg.Add(1)
	runHttpServer(&wg)

	wg.Wait()
}
