package domain

import (
	"net/http"
	"testing"

	"github.com/sharmari/mvc/controllers"
	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := controllers.GetUser(0)
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found ", err.Message)
	/* Go way
	if user != nil {
		t.Error("We were not expecting a user with id 0")
	}

	if err == nil {
		t.Error("we were expecting an error when user id is 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("we were expecting 404 when user id is 0")
	}
	*/
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Rishabh", user.FirstName)
	assert.EqualValues(t, "Sharma", user.LastName)
	assert.EqualValues(t, "sharmarishabh148@gmail.com", user.Email)
}
