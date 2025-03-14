package service

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUser returns response with user data by user id.
// Example response: {"id":1,"name":"Ferdinande","createdAt":"2020-07-14T05:48:54.798Z"}
func (s *Service) GetUser(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	numId, err := strconv.Atoi(id)
	if err != nil {
		sendBadRequest(resp, ErrInvalidUserId)
		return
	}

	user, ok := s.Users[numId]
	if !ok {
		sendNotFound(resp, ErrUserNotFound)
		return
	}

	sendResponse(resp, user)
}
