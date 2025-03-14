package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"lidiyatrf/user-actions-test-task/internal/config"
)

type Service struct {
	Users   map[int]User
	Actions map[int][]Action
}

// New creates new service and populates actions and users from files specified in config.
func New(config config.Config) (*Service, error) {
	var users []User
	if err := readDataFromFile(config.UsersFilePath, &users); err != nil {
		return nil, fmt.Errorf("unable to read users file: %v", err)
	}

	var actions []Action
	if err := readDataFromFile(config.ActionsFilePath, &actions); err != nil {
		return nil, fmt.Errorf("unable to read actions file: %v", err)
	}

	return &Service{
		Users:   usersToMap(users),
		Actions: actionsToMap(actions),
	}, nil
}

func readDataFromFile(filePath string, result any) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf(" %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	if err := json.Unmarshal(bytes, result); err != nil {
		return fmt.Errorf("failed to unmarshal file: %v", err)
	}
	return nil
}

func usersToMap(users []User) map[int]User {
	usersMap := make(map[int]User, len(users))
	for _, user := range users {
		usersMap[user.Id] = user
	}
	return usersMap
}

func actionsToMap(actions []Action) map[int][]Action {
	actionsMap := make(map[int][]Action)
	for _, action := range actions {
		actionsMap[action.UserId] = append(actionsMap[action.UserId], action)
	}
	return actionsMap
}
