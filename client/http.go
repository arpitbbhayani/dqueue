package client

import (
	"encoding/json"
	"fmt"

	"github.com/arpitbbhayani/dqueue/models"
)

type DqueueHTTPClient struct {
	URL string
}

func (q DqueueHTTPClient) PutMessage(request *models.PutMessageRequest) *models.PutMessageResponse {
	var response models.PutMessageResponse
	url := fmt.Sprintf("%s/msg", q.URL)
	payload, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(doPut(url, payload), &response)
	if err != nil {
		panic(err)
	}
	return &response
}

func (q DqueueHTTPClient) GetMessage(request *models.GetMessageRequest) *models.GetMessageResponse {
	var response models.GetMessageResponse
	url := fmt.Sprintf("%s/msg", q.URL)
	err := json.Unmarshal(doGet(url), &response)
	if err != nil {
		panic(err)
	}
	return &response
}
