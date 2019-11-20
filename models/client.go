package models

type DqueueClient interface {
	PutMessage(*PutMessageRequest) *PutMessageResponse
	GetMessage(*GetMessageRequest) *GetMessageResponse
}
