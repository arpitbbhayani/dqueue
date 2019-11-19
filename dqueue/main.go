package dqueue

import "fmt"

var (
	dq *Dqueue
)

type Dqueue struct{}

func initialize() {
	fmt.Println("initializing queue", dq)
}

func GetInstance() *Dqueue {
	if dq == nil {
		dq = &Dqueue{}
		initialize()
	}
	return dq
}

func (q *Dqueue) PutMessage(req *PutMessageRequest) *PutMessageResponse {
	return &PutMessageResponse{
		ID: "1",
	}
}
