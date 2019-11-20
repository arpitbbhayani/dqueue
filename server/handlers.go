package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/arpitbbhayani/dqueue/dqueue"
	"github.com/arpitbbhayani/dqueue/models"
)

func readHTTPRequestBody(reader io.Reader, obj interface{}) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		panic(err)
	}
}

func sendHTTPJSONResponse(obj interface{}, code int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	response, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	sendHTTPJSONResponse(models.HTTPVersionResponse{
		Version: "1.0.0",
	}, http.StatusOK, w)
}

func MessagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var request models.HTTPMessagePutRequest
		readHTTPRequestBody(r.Body, &request)
		dq := dqueue.GetInstance()
		response := dq.PutMessage(request.ToDqueueMessagePutRequest())
		sendHTTPJSONResponse(response, http.StatusOK, w)
	default:
		sendHTTPJSONResponse(models.HTTPError{
			Message: "method not allowed",
		}, http.StatusMethodNotAllowed, w)
	}
}
