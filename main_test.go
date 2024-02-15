package main

import (
	"bytes"
	"fmt"

	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jicodes/webapp/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	// Initialize your Gin router
	router := setupRouter()

	// Create an user 
	createUserPayload := []byte(`{"first_name" : "First", "last_name" : "Last", "password" : "TestPassword", "username": "testuser@example.com"}`)
	createUserReq, _ := http.NewRequest("POST", "/v1/user", bytes.NewBuffer(createUserPayload))
	createUserReq.Header.Set("Content-Type", "application/json")

	createUserResp := httptest.NewRecorder()

	router.ServeHTTP(createUserResp, createUserReq)

	assert.Equal(t, http.StatusCreated, createUserResp.Code, "Should return a 201 status code for a successful user creation")
	// Validate the user was created 
	getUserReq, _ := http.NewRequest("GET", "/v1/user/self", nil)
	getUserReq.Header.Set("Content-Type", "application/json")

	username := "testuser@example.com"
	password := "TestPassword"
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	getUserReq.Header.Set("Authorization", basicAuth)

	getUserResp := httptest.NewRecorder()
	router.ServeHTTP(getUserResp, getUserReq)

	assert.Equal(t, http.StatusOK, getUserResp.Code)

	var createdUser models.User
	err := json.Unmarshal(getUserResp.Body.Bytes(), &createdUser)
	if err != nil {
    t.Fatalf("Failed to parse response body: %v", err)
	}

	expectedUser := models.User{
		FirstName: "First",
		LastName:  "Last",
		Username:  "testuser@example.com",
	}

	assert.Equal(t, expectedUser.FirstName, createdUser.FirstName)
	assert.Equal(t, expectedUser.LastName, createdUser.LastName)
	assert.Equal(t, expectedUser.Username, createdUser.Username)
}

func TestUpdateUser(t *testing.T) {
	router := setupRouter()
	// Test 2 - Update the account and validate the account was updated
	updateUserPayload := []byte(`{"first_name" : "UpdatedFirst", "last_name" : "UpdatedLast", "password" : "UpdatedPassword"}`)
	updateUserReq, _ := http.NewRequest("PUT", "/v1/user/self", bytes.NewBuffer(updateUserPayload))
	updateUserReq.Header.Set("Content-Type", "application/json")

	username := "testuser@example.com"
	password := "TestPassword"
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	updateUserReq.Header.Set("Authorization", basicAuth)

	updateUserResp := httptest.NewRecorder()
	router.ServeHTTP(updateUserResp, updateUserReq)

	assert.Equal(t, http.StatusOK, updateUserResp.Code)

	fmt.Println("updatedUser response body: ", updateUserResp.Body.String())

	// Validate the updated user information
	var updatedUser models.User
    err := json.Unmarshal(updateUserResp.Body.Bytes(), &updatedUser)
    if err != nil {
        t.Fatalf("Failed to parse response body: %v", err)
    }

	fmt.Println("updatedUser:", updatedUser)

	assert.Equal(t, "UpdatedFirst", updatedUser.FirstName)
	assert.Equal(t, "UpdatedLast", updatedUser.LastName)

	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte("UpdatedPassword"))
	assert.NoError(t, comparePasswordErr)

}