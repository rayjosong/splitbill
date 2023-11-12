package repository

import (
	"fmt"

	"github.com/rayjosong/splitbill/pkg/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (r UserRepository) Save(user user.User) error {
	result := r.db.Create(&user)

	if result.Error != nil {
		return fmt.Errorf("error saving to db: %w", result.Error)
	}

	return nil
}
