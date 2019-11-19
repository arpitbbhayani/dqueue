package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/arpitbbhayani/dqueue/dqueue"
	"github.com/gorilla/mux"
)

func runHttpServer(wg *sync.WaitGroup) {
	defer wg.Done()

	r := mux.NewRouter()
	r.HandleFunc("/", VersionHandler)
	r.HandleFunc("/msg", MessagesHandler)
	http.Handle("/", r)

	fmt.Println("Listening to :4000")
	http.ListenAndServe(":4000", nil)
}

func Run() {
	var wg sync.WaitGroup
	_ = dqueue.GetInstance()

	wg.Add(1)
	go runHttpServer(&wg)

	wg.Wait()
}
