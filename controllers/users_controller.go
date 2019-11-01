package controllers

import (
	"encoding/json"
	"github.com/x777/mvc_pattern/services"
	"github.com/x777/mvc_pattern/utils"
	"log"
	"net/http"
	"strconv"
)

// User controller...
func GetUser(resp http.ResponseWriter, req *http.Request){
	userIdParam := req.URL.Query().Get("user_id")
	log.Printf("About to process user_id %v", userIdParam)

	userId, err := (strconv.ParseInt(userIdParam, 10, 64))
	if err != nil {
		// Just return the Bad Request to the cliend
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		userErr, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(userErr)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		//	Handler the error and return to the client
		userErr, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(userErr))
		return

	}

//	return user to client
	jsonValue, _:= json.Marshal(user)
	resp.Write(jsonValue)

}
