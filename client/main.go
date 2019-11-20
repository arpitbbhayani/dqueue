package client

import (
	"github.com/arpitbbhayani/dqueue/models"
	"github.com/c-bata/go-prompt"
)

type DqueueClient interface {
	PutMessage(*models.PutMessageRequest) *models.PutMessageResponse
	GetMessage(*models.GetMessageRequest) *models.GetMessageResponse
}

var (
	dqclient DqueueClient
)

func initDqueueClient(protocol string) {
	switch protocol {
	case "http":
		dqclient = DqueueHTTPClient{
			URL: "http://localhost:4000",
		}
	default:
		panic("invalid protocol. allowed protocol: http")
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
