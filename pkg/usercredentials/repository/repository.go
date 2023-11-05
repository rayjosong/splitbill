package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rayjosong/splitbill/pkg/usercredentials"
)

type UserCredentialsRepository struct {
	db *gorm.DB
}

func (r UserCredentialsRepository) Save(c usercredentials.UserCredentials) error {
	result := r.db.Create(&c)

	if result.Error != nil {
		return fmt.Errorf("error saving to db: %w", result.Error)
	}

	return nil
}

func (r UserCredentialsRepository) FindCredentials(username string, password string) (*usercredentials.UserCredentials, error) {
	var credentials usercredentials.UserCredentials
	if err := r.db.Where("username = ?", username).First(&credentials).Error; err != nil {
		return nil, fmt.Errorf("error finding credentials: %w", err)
	}

	return &credentials, nil
}
