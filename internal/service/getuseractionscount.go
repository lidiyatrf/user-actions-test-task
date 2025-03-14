package service

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CountResponse struct {
	Count int `json:"count"`
}

// GetUserActionsCount returns response with count of user actions by user id.
// Example response: {"count":49}
func (s *Service) GetUserActionsCount(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	numId, err := strconv.Atoi(id)
	if err != nil {
		sendBadRequest(resp, ErrInvalidUserId)
		return
	}

	userActions, ok := s.Actions[numId]
	if !ok {
		sendNotFound(resp, ErrUserNotFound)
		return
	}

	sendResponse(resp, CountResponse{Count: len(userActions)})
}
