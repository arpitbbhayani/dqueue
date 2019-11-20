package dqueue

import (
	"fmt"

	"github.com/arpitbbhayani/dqueue/models"
)

var (
	dq *Dqueue
)

type Dqueue struct{}

func initialize() {
	fmt.Println("initializing queue ...")
}

func GetInstance() *Dqueue {
	if dq == nil {
		dq = &Dqueue{}
		initialize()
	}
	return dq
}

func (q *Dqueue) PutMessage(req *models.PutMessageRequest) *models.PutMessageResponse {
	id := "1"
	return &models.PutMessageResponse{
		ID: id,
	}
}

func (q *Dqueue) GetMessage(req *models.GetMessageRequest) *models.GetMessageResponse {
	id := "1"
	return &models.GetMessageResponse{
		ID:      id,
		Message: "message-dummy",
	}
}
