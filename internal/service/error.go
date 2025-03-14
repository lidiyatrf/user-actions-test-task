package service

import "fmt"

var (
	ErrRecursiveReferrals = fmt.Errorf("recursive referrals")
	ErrInvalidUserId      = fmt.Errorf("invalid user id")
	ErrUserNotFound       = fmt.Errorf("user not found")
	ErrBlankActionType    = fmt.Errorf("action type can't be blank")
)
