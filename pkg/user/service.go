package user

import (
	"errors"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rayjosong/splitbill/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserSessionService struct {
	repo UserCreationRepo
}

func NewUserSessionService(repo UserCreationRepo) UserSessionService {
	return UserSessionService{repo}
}

type UserCreationRepo interface {
	Save(user models.User) error
	GetUserFromID(userID string) (*models.User, error)
}

func (s UserSessionService) CreateUser(user models.User, username string, password string) error {
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

func (s UserSessionService) GetCurrentUser(c *gin.Context) (*models.User, error) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID == nil {
		return nil, errors.New("user not logged in")
	}

	id, ok := userID.(string)
	if !ok {
		return nil, errors.New("cannot cast userid to string")
	}

	user, err := s.repo.GetUserFromID(id)
	if err != nil {
		return nil, fmt.Errorf("cannot get user from ID: %w", err)
	}

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return &user, nil

