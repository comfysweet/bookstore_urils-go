package errors

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServiceError(t *testing.T) {
	err := NewInternalServiceError("this is the message", NewError("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "database error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
}
