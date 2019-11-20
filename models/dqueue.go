package models

type PutMessageRequest struct {
	Message *string `json:"message"`
}

type PutMessageResponse struct {
	ID *string `json:"id"`
}

type GetVersionResponse struct {
	Version string `json:"version"`
}
