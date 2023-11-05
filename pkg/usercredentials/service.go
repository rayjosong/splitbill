package usercredentials

import (
	"fmt"

	"github.com/rayjosong/splitbill/pkg/user"
	"golang.org/x/crypto/bcrypt"
)

type UserCredentialsSerivce struct {
	UserCredentialsRepo UserCredentialsRepo
}

type UserCredentialsRepo interface {
	Save(credentials UserCredentials) error
	FindCredentials(username string) (UserCredentials, error)
}

func NewUserCredentialsService(repo UserCredentialsRepo) UserCredentialsSerivce {
	return UserCredentialsSerivce{UserCredentialsRepo: repo}
}

func (s UserCredentialsSerivce) SaveCredentials(user user.User, username string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("cannot hash password: %w", err)
	}

	credentials := UserCredentials{UserID: user.ID, Username: username, Password: string(hashedPassword)}
	if err := s.UserCredentialsRepo.Save(credentials); err != nil {
		return fmt.Errorf("cannot save user credentials: %w", err)
	}

	return nil
}

func (s UserCredentialsSerivce) GetCredentialsForUser(username string) (*UserCredentials, error) {
	savedCredentials, err := s.UserCredentialsRepo.FindCredentials(username)
	if err != nil {
		return nil, fmt.Errorf("error finding credentials: %w", err)
	}

	return &savedCredentials, nil
}

func (s UserCredentialsSerivce) IsCredentialValid(username string, password string) (bool, error) {
	savedCredentials, err := s.UserCredentialsRepo.FindCredentials(username)
	if err != nil {
		return false, fmt.Errorf("error while verifying credentials: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(savedCredentials.Password), []byte(password)); err != nil {
		return false, fmt.Errorf("the password does not match")
	}

	return true, nil
}
