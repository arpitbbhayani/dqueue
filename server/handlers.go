package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/arpitbbhayani/dqueue/dqueue"
	"github.com/arpitbbhayani/dqueue/models"
	"github.com/sirupsen/logrus"
)

func readHTTPRequestBody(reader io.Reader, obj interface{}) error {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		return err
	}
	return nil
}

func sendHTTPJSONResponse(obj interface{}, code int, w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	response, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	sendHTTPJSONResponse(models.GetVersionResponse{
		Version: "1.0.0",
	}, http.StatusOK, w)
}

func MessagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var request models.PutMessageRequest
		if err := readHTTPRequestBody(r.Body, &request); err != nil {
			logrus.Error(err)
			return
		}

		dq := dqueue.GetInstance()
		response := dq.PutMessage(&request)
		if err := sendHTTPJSONResponse(response, http.StatusOK, w); err != nil {
			logrus.Error(err)
			return
		}
	case "GET":
		var request models.GetMessageRequest
		dq := dqueue.GetInstance()
		response := dq.GetMessage(&request)
		if err := sendHTTPJSONResponse(response, http.StatusOK, w); err != nil {
			logrus.Error(err)
			return
		}
	default:
		if err := sendHTTPJSONResponse(models.HTTPError{
			Message: "method not allowed",
		}, http.StatusMethodNotAllowed, w); err != nil {
			logrus.Error(err)
			return
		}
	}
}
