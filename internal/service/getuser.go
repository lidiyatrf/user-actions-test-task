package service

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
