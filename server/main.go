package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/arpitbbhayani/dqueue/dqueue"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func runHttpServer(wg *sync.WaitGroup) {
	defer wg.Done()

	r := mux.NewRouter()
	r.HandleFunc("/", VersionHandler)
	r.HandleFunc("/msg", MessagesHandler)
	http.Handle("/", r)

	port := viper.GetInt("port")
	logrus.Infof("dqueue server started on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}

func Run(configPath string) {
	InitializeConfig(configPath)

	var wg sync.WaitGroup
	_ = dqueue.GetInstance()

	wg.Add(1)
	runHttpServer(&wg)

	wg.Wait()
}
