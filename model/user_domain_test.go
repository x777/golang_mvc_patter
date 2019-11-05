package model

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization:
	//Execution:
	user, err := GetUser(0)
	// Validation:
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exists", err.Message)

}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Yas", user.FirstName)
	assert.EqualValues(t, "Dimon", user.LastName)
	assert.EqualValues(t, "diaysrb@gmail.com", user.Email)
}
