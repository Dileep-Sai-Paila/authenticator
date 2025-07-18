package services

import (
	"authenticator/internal/models"
	"authenticator/internal/repositories"
	"authenticator/pkg/utils"
	"errors"
)

// handles the business logic for user registration
func RegisterUser(email, password string) (*models.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	newUser := &models.User{
		Email:    email,
		Password: hashedPassword,
	}
	createdUser, err := repositories.CreateUser(newUser)
	if err != nil {
		return nil, errors.New("failed to create user")
	}
	return createdUser, nil
}

// handles the business logic for user login
func LoginUser(email, password string) (string, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}
	tokenString, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", errors.New("could not generate token")
	}
	return tokenString, nil
}

// retrieves a user's profile by their email
func GetUserProfile(email string) (*models.User, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
