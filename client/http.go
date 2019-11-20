package client

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const URL = "http://localhost:4000"

func doPut(url string, payload string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, strings.NewReader(payload))
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
