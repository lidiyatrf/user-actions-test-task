package service

import "time"

type Action struct {
	Id         int       `json:"id"`
	Type       string    `json:"type"`
	UserId     int       `json:"userId"`
	TargetUser int       `json:"targetUser"`
	CreatedAt  time.Time `json:"createdAt"`
}
