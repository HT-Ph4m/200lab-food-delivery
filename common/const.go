package common

import "log"

func AppRecover() {
	if err := recover(); err != nil {

		log.Println("Recover", err)
	}
}

const (
	DbTypeRestaurant = 1
	DbTypeFood       = 2
	DbTypeCategory   = 3
	DbTypeUser       = 4
)

// const CurrentUser = "user"

// type Requester interface {
// 	GetUserId() int
// 	GetEmail() string
// 	GetRole() string
// }
