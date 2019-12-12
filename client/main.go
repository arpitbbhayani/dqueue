package client

import (
	"github.com/arpitbbhayani/dqueue/models"
	"github.com/c-bata/go-prompt"
)

var (
	dqclient models.DqueueClient
)

func initDqueueClient(protocol string) {
	switch protocol {
	case "http":
		dqclient = DqueueHTTPClient{
			URL: "http://localhost:4096",
		}
	default:
		return
	}
}

func Run(protocol string) {
	initDqueueClient(protocol)
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("dqueue> "),
		prompt.OptionTitle("dqueue"),
	)
	p.Run()
}
