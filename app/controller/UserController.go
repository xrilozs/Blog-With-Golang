package controller

import (
	"app/helper"
	"app/models"
	"app/service"
	"encoding/json"
	"net/http"
	"time"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUserRequest models.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&loginUserRequest)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}

	user, err := service.GetUserByEmail(loginUserRequest.Email)
	if user == nil && err != nil {
		helper.SendResponse(w, 400, "Email not found", nil)
		return
	}

	errVerify := helper.VerifyPasswordHash(loginUserRequest.Password, user.PasswordHash)
	if errVerify != nil {
		helper.SendResponse(w, 400, "Invalid password", nil)
		return
	}

	token, err := helper.GenerateToken(loginUserRequest.Email)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}
	refresh, err := helper.GenerateRefresh(loginUserRequest.Email)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}
	tokenResponse := &models.TokenResponse{
		Token:   token,
		Refresh: refresh,
	}

	helper.SendResponse(w, 200, "User login is success", tokenResponse)
	return
}

func RefreshUser(w http.ResponseWriter, r *http.Request) {
	token, err := helper.GetAuthHeader(r)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}
	registeredClaim, err := helper.ValidateRefresh(token)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}

	user, err := service.GetUserByEmail(registeredClaim.Subject)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}

	refreshToken, err := helper.GenerateToken(user.Email)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}
	refresh, err := helper.GenerateRefresh(user.Email)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}
	tokenResponse := &models.TokenResponse{
		Token:   refreshToken,
		Refresh: refresh,
	}

	helper.SendResponse(w, 200, "User refresh is success", tokenResponse)
	return
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var reqUser models.UserRegister
	err := json.NewDecoder(r.Body).Decode(&reqUser)

	existingUser, _ := service.GetUserByEmail(reqUser.Email)
	if existingUser != nil {
		helper.SendResponse(w, 400, "Email already registered", nil)
		return
	}

	passwordHash, err := helper.HashPassword(reqUser.Password)
	if err != nil {
		helper.SendResponse(w, 400, err.Error(), nil)
		return
	}

	newUser := &models.User{
		PasswordHash: passwordHash,
		Name:         reqUser.Name,
		Email:        reqUser.Email,
		CreatedAt:    time.Now(),
	}
	user, err := service.Adduser(*newUser)
	if err != nil {
		helper.SendResponse(w, 400, "User register is failed", nil)
		return
	}

	helper.SendResponse(w, 200, "User register is success", user)
	return
}
