package tokenprovider

import (
	"200lab-project-1/common"
	"errors"
	"time"
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)
