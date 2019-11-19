package server

import (
	"github.com/arpitbbhayani/dqueue/dqueue"
)

type HTTPError struct {
	Message string
}

type HTTPVersionRequest struct{}

type HTTPVersionResponse struct {
	Version string
}

type HTTPMessagePutRequest struct {
	Message string `json:"message"`
}

func (r *HTTPMessagePutRequest) ToDqueueMessagePutRequest() *dqueue.PutMessageRequest {
	return &dqueue.PutMessageRequest{
		Message: &r.Message,
	}
}

type HTTPMessagePutResponse struct {
	ID string
}
