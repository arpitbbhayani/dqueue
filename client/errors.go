package client

import (
	"github.com/arpitbbhayani/cerr"
)

type DqueueError struct {
	*cerr.Cerr
}

type DqueueTransportError struct {
	*cerr.Cerr
}

func NewDqueueTransportError(err error, isTimeout bool, isRetryable bool) *DqueueTransportError {
	return &DqueueTransportError{
		Cerr: cerr.NewCerrBuilder().
			SetErrorCode("transport_error").
			SetOriginalError(err).
			SetIsRetryable(isRetryable).
			SetIsTimeout(isTimeout).
			SetMessage("Unable to communicate with Dqueue server. Check if the server is up and running.").
			Build(),
	}
}
