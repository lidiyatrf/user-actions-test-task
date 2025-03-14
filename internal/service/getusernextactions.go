package service

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Service) GetUserNextActions(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	actionType := vars["actionType"]

	if actionType == "" {
		sendBadRequest(resp, ErrBlankActionType)
		return
	}

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

	sendResponse(resp, s.calculateUserNextActions(userActions, actionType))
}

func (s *Service) calculateUserNextActions(userActions []Action, actionType string) map[string]float64 {
	nextActions := make(map[string]int)
	var totalActionsCounter int
	for i, currentAction := range userActions {
		if currentAction.Type == actionType && i+1 < len(userActions) {
			nextAction := userActions[i+1].Type
			nextActions[nextAction]++
			totalActionsCounter++
		}
	}

	probabilities := make(map[string]float64, len(nextActions))
	for k, v := range nextActions {
		probabilities[k] = math.Round(float64(v)/float64(totalActionsCounter)*100) / 100
	}
	adjustProbabilities(probabilities)

	return probabilities
}

// adjustProbabilities adjusts probabilities so that sum is equal to 1
func adjustProbabilities(nextActions map[string]float64) {
	var keyToAdjust string
	var currentSum float64
	for k, v := range nextActions {
		if k > keyToAdjust {
			keyToAdjust = k
		}
		currentSum += v
	}
	if currentSum != 1 && keyToAdjust != "" {
		nextActions[keyToAdjust] = nextActions[keyToAdjust] + (1 - currentSum)
	}
}
