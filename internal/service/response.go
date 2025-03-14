package service

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func sendResponse(resp http.ResponseWriter, result any) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	if result == nil {
		return
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		log.Println("cannot marshal:", err)
		return
	}
	if _, err := resp.Write(bytes); err != nil {
		log.Println("cannot write bytes to response:", err)
		return
	}
}

func sendError(resp http.ResponseWriter, httpStatus int, err error) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(httpStatus)
	if err == nil {
		return
	}

	response := ErrorResponse{Error: err.Error()}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Println("cannot marshal:", err)
		return
	}
	if _, err := resp.Write(bytes); err != nil {
		log.Println("cannot write bytes to response:", err)
		return
	}
}

func sendBadRequest(resp http.ResponseWriter, err error) {
	sendError(resp, http.StatusBadRequest, err)
}

func sendNotFound(resp http.ResponseWriter, err error) {
	sendError(resp, http.StatusNotFound, err)
}

func sendInternalError(resp http.ResponseWriter, err error) {
	sendError(resp, http.StatusInternalServerError, err)
}
