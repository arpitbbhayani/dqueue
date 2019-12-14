package client

import (
	"fmt"

	"github.com/arpitbbhayani/dqueue/models"
	"github.com/c-bata/go-prompt"
)

var (
	dqclient models.DqueueClient
)

func initDqueueClient(protocol string, host string, port int) {
	switch protocol {
	case "http":
		dqclient = DqueueHTTPClient{
			URL: fmt.Sprintf("http://%s:%d", host, port),
		}
	default:
		return
	}
}

func Run(protocol string, host string, port int) {
	initDqueueClient(protocol, host, port)
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("dqueue> "),
		prompt.OptionTitle("dqueue"),
	)
	p.Run()
}
