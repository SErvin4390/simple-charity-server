package main

import (
	"encoding/json"
	"net/http"
)

// APIReturn represents a standard API return object
type APIReturn struct {
	Data interface{} `json:"data,omitempty"`
}

// Send standardizes the return from the API for 20X actions
func Send(w http.ResponseWriter, code int, payload interface{}) {
	ret := APIReturn{}
	ret.Data = payload
	response, _ := json.Marshal(ret)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// SendError standardizes the return from the API for 40X actions
func SendError(w http.ResponseWriter, code int, message string) {
	if message == "" {
		switch code {
		case 400:
			message = "Client Error"
		case 401:
			message = "Not Authorized"
		case 403:
			message = "Not Authorized"
		case 404:
			message = "Not Found"
		}
	}

	ret := APIReturn{}
	ret.Data = map[string]string{
		"message": message,
	}
	Send(w, code, ret)
}
