package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserSessionService struct {
	repo UserCreationRepo
}

func NewUserSessionService(repo UserCreationRepo) UserSessionService {
	return UserSessionService{repo}
}

type UserCreationRepo interface {
	Save(user User) error
}

func (s UserSessionService) CreateUser(user User, username string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("cannot hash password: %w", err)
	}

	user = user.SetUsernameAndPassword(username, string(hashedPassword))

	if err := s.repo.Save(user); err != nil {
		return fmt.Errorf("cannot save user: %w", err)
	}
	return nil
}

func (s UserSessionService) IsCredentialValid(inputPassword string, savedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(inputPassword)); err != nil {
		return false
	}

	return true
}
