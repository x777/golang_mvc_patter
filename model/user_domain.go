package model

import (
	"fmt"
	"github.com/x777/mvc_pattern/utils"
	"net/http"
)
// Fake DB data
var (
	users = map[int64]*User {
		123: {
			Id:        1,
			FirstName: "Yas",
			LastName:  "Dimon",
			Email:     "diaysrb@gmail.com",
		},
	}
)

// User domain for interacting with DB
func GetUser(userId int64) (*User, *utils.ApplicationError){
	if user := users[userId]; user != nil{
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code:"not_found",
	}

}
