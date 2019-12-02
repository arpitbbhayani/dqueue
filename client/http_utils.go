package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func makeHTTPCall(url string, method string, payload []byte) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, bytes.NewReader(payload))
	request.ContentLength = int64(len(payload))
	response, err := client.Do(request)
	if err != nil {
		return nil, NewDqueueTransportError(err, false, false)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, NewDqueueTransportError(err, false, false)
	}
	return contents, nil

}

func doPut(url string, payload []byte) ([]byte, error) {
	return makeHTTPCall(url, "PUT", payload)
}

func doGet(url string) ([]byte, error) {
	return makeHTTPCall(url, "GET", []byte{})
}
