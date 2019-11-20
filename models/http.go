package models

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

func (r *HTTPMessagePutRequest) ToDqueueMessagePutRequest() *PutMessageRequest {
	return &PutMessageRequest{
		Message: &r.Message,
	}
}

type HTTPMessagePutResponse struct {
	ID string
}
