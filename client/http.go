package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/arpitbbhayani/dqueue/models"
)

func doPut(url string, payload []byte) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, bytes.NewReader(payload))
	request.ContentLength = int64(len(payload))
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		return contents
	}
}

func doGet(url string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		return contents
	}
}

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
