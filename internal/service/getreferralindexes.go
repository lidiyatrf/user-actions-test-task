package service

import (
	"net/http"
)

func (s *Service) GetReferralIndexes(resp http.ResponseWriter, req *http.Request) {
	result, err := s.calculateReferralIndexes()
	if err != nil {
		sendInternalError(resp, err)
		return
	}
	sendResponse(resp, result)
}

func (s *Service) calculateReferralIndexes() (map[int]int, error) {
	result := make(map[int]int, len(s.Users))

	for _, user := range s.Users {
		_, err := s.calculateReferrals(user.Id, result)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (s *Service) calculateReferrals(userId int, result map[int]int) (int, error) {
	if count, ok := result[userId]; ok {
		if count < 0 {
			return 0, ErrRecursiveReferrals
		}
		return count, nil
	}

	result[userId] = -1

	var count int
	for _, action := range s.Actions[userId] {
		if action.Type == "REFER_USER" && action.UserId != action.TargetUser {
			res, err := s.calculateReferrals(action.TargetUser, result)
			if err != nil {
				return 0, err
			}
			count += 1 + res
		}
	}
	result[userId] = count
	return count, nil
}
