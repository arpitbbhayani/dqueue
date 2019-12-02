package models

type DqueueClient interface {
	PutMessage(*PutMessageRequest) (*PutMessageResponse, error)
	GetMessage(*GetMessageRequest) (*GetMessageResponse, error)
}
