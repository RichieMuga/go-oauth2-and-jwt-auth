// Package controllers defines the tests for controller
package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T){
	// Initialize the router
	router := SetupTestRouter()

  // Register routes from the router package
  req, err := http.NewRequest(http.MethodGet, "/api/v1/ping", nil)
  assert.NoError(t,err)

  // Create response recorder
  recorder := httptest.NewRecorder()

  // use existing router to serve the request
  router.ServeHTTP(recorder, req)

  // assertions to check responseassert.Equal(t, http.StatusOK, recorder.code)
  assert.Equal(t, http.StatusOK, recorder.Code)
  
  expectedResponse:=`{"message":"pong"}`
  assert.JSONEq(t, expectedResponse, recorder.Body.String(), "Response should be pong")
}
