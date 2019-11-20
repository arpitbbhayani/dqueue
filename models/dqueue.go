package models

import "encoding/json"

func toJSONString(x interface{}) string {
	b, _ := json.Marshal(x)
	return string(b)
}

type PutMessageRequest struct {
	Message string `json:"message"`
}

type PutMessageResponse struct {
	ID string `json:"id"`
}

func (x PutMessageResponse) ToString() string {
	return toJSONString(x)
}

type GetVersionResponse struct {
	Version string `json:"version"`
}

func (x GetVersionResponse) ToString() string {
	return toJSONString(x)
}

type GetMessageResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func (x GetMessageResponse) ToString() string {
	return toJSONString(x)
}

type GetMessageRequest struct{}
