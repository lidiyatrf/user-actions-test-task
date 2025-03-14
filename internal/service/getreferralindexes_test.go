package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculateReferralIndexes(t *testing.T) {
	tests := []struct {
		name           string
		userCount      int
		actions        []Action
		expectedResult map[int]int
		expectedErr    error
	}{
		{
			name:      "simple_test",
			userCount: 5,
			actions: []Action{
				{
					Type:       "REFER_USER",
					UserId:     0,
					TargetUser: 1,
				},
				{
					Type:       "REFER_USER",
					UserId:     1,
					TargetUser: 3,
				},
				{
					Type:       "REFER_USER",
					UserId:     0,
					TargetUser: 2,
				},
			},
			expectedResult: map[int]int{
				0: 3,
				1: 1,
				2: 0,
				3: 0,
				4: 0,
			},
		},
		{
			name:      "self_invite",
			userCount: 2,
			actions: []Action{
				{
					Type:       "REFER_USER",
					UserId:     0,
					TargetUser: 1,
				},
				{
					Type:       "REFER_USER",
					UserId:     1,
					TargetUser: 1,
				},
			},
			expectedResult: map[int]int{
				0: 1,
				1: 0,
			},
		},
		{
			name:      "recusive_invites",
			userCount: 3,
			actions: []Action{
				{
					Type:       "REFER_USER",
					UserId:     0,
					TargetUser: 1,
				},
				{
					Type:       "REFER_USER",
					UserId:     1,
					TargetUser: 2,
				},
				{
					Type:       "REFER_USER",
					UserId:     2,
					TargetUser: 1,
				},
			},
			expectedErr: ErrRecursiveReferrals,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			users := make([]User, test.userCount)
			for i := 0; i < test.userCount; i++ {
				users[i] = User{Id: i}
			}
			service := Service{
				Users:   usersToMap(users),
				Actions: actionsToMap(test.actions),
			}

			referralIndexes, err := service.calculateReferralIndexes()
			require.Equal(t, test.expectedErr, err)
			if test.expectedErr == nil {
				assert.Len(t, referralIndexes, test.userCount)
				assert.Equal(t, referralIndexes, test.expectedResult)
			}
		})
	}
}
