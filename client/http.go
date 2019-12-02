package client

import (
	"encoding/json"
	"fmt"

	"github.com/arpitbbhayani/dqueue/models"
)

type DqueueHTTPClient struct {
	URL string
}

func (q DqueueHTTPClient) PutMessage(request *models.PutMessageRequest) (*models.PutMessageResponse, error) {
	var response models.PutMessageResponse
	url := fmt.Sprintf("%s/msg", q.URL)
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	b, err := doPut(url, payload)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (q DqueueHTTPClient) GetMessage(request *models.GetMessageRequest) (*models.GetMessageResponse, error) {
	var response models.GetMessageResponse
	url := fmt.Sprintf("%s/msg", q.URL)
	b, err := doGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
