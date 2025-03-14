package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateUserNextActions(t *testing.T) {
	tests := []struct {
		name           string
		userActions    []Action
		actionType     string
		expectedResult map[string]float64
		expectedErr    error
	}{
		{
			name: "1/2_test",
			userActions: []Action{
				{
					Type: "ADD_CONTACT",
				},
				{
					Type: "EDIT_CONTACT",
				},
				{
					Type: "ADD_CONTACT",
				},
				{
					Type: "VIEW_CONTACTS",
				},
			},
			actionType: "ADD_CONTACT",
			expectedResult: map[string]float64{
				"EDIT_CONTACT":  0.5,
				"VIEW_CONTACTS": 0.5,
			},
		},
		{
			name: "1/3_test",
			userActions: []Action{
				{
					Type: "WELCOME",
				},
				{
					Type: "CONNECT_CRM",
				},
				{
					Type: "EDIT_CONTACT",
				},
				{
					Type: "EDIT_CONTACT",
				},
				{
					Type: "EDIT_CONTACT",
				},
				{
					Type: "ADD_CONTACT",
				},
				{
					Type: "VIEW_CONTACTS",
				},
			},
			actionType: "EDIT_CONTACT",
			expectedResult: map[string]float64{
				"ADD_CONTACT":  0.33,
				"EDIT_CONTACT": 0.67,
			},
		},
		{
			name: "1/7_test",
			userActions: []Action{
				{
					Type: "EDIT_CONTACT",
				},
				{
					Type: "ADD_CONTACT",
				},
				{
					Type: "EDIT_CONTACT",
				},
				{
					Type: "VIEW_CONTACTS",
				},
				{
					Type: "EDIT_CONTACT",
				},
				{
					Type: "VIEW_CONTACTS",
				},
				{
					Type: "EDIT_CONTACT",
				},
				{
					Type: "VIEW_CONTACTS",
				},
				{
					Type: "EDIT_CONTACT",
				},
				{
					Type: "EDIT_CONTACT",
				},
				{
					Type: "EDIT_CONTACT",
				},
				{
					Type: "ADD_CONTACT",
				},
			},
			actionType: "EDIT_CONTACT",
			expectedResult: map[string]float64{
				"ADD_CONTACT":   0.29,
				"EDIT_CONTACT":  0.29,
				"VIEW_CONTACTS": 0.42,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := Service{}
			result := service.calculateUserNextActions(test.userActions, test.actionType)
			require.Equal(t, test.expectedResult, result)
		})
	}
}
