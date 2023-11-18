package repository

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rayjosong/splitbill/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (r UserRepository) Save(user models.User) error {
	result := r.db.Create(&user)

	if result.Error != nil {
		return fmt.Errorf("error saving to db: %w", result.Error)
	}

	return nil
}

func (r UserRepository) GetUserFromID(userID string) (*models.User, error) {
	var user models.User

	result := r.db.First(&user, "id = ?", userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}
