package dqueue

type PutMessageRequest struct {
	Message *string
}

type PutMessageResponse struct {
	ID *string
}
