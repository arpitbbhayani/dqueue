package server

import (
	"encoding/json"
	"net/http"
)

func sendHTTPJSONResponse(obj interface{}, code int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	response, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.WriteHeader(code)
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	sendHTTPJSONResponse(VersionResponse{
		Version: "1.0.0",
	}, http.StatusOK, w)
}

func MessagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
	default:
		sendHTTPJSONResponse(HTTPError{
			Message: "method not allowed",
		}, http.StatusMethodNotAllowed, w)
	}
}
