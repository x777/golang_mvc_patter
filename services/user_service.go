package services

import (
	"github.com/x777/mvc_pattern/model"
	"github.com/x777/mvc_pattern/utils"
)

// User service...
// Call domain for GetUser
func GetUser(userId int64) (*model.User, *utils.ApplicationError){
	return model.GetUser(userId)
}
