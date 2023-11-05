package usercredentials

import (
	"github.com/rayjosong/splitbill/pkg/user"

	"gorm.io/gorm"
)

type UserCredentials struct {
	gorm.Model
	UserID   uint      `gorm:"unique;not null"`
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	User     user.User `gorm:"foreignkey:UserID"`
}
