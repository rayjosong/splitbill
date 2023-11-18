package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `gorm:"type:varchar(100);unique_index" json:"email"`
	Phone    string    `json:"phone,omitempty"`
	Balance  float32   `json:"balance"`
	Expenses []Expense `gorm:"foreignkey:CreatorID"`
	Groups   []Group   `gorm:"many2many:user_groups;"`
	Shares   []Share   `gorm:"foreignkey:FromUserID"`
	Username string    `gorm:"uniquelnotnull"`
	Password string    `gorm:"not null"`
}

func (u User) SetUsernameAndPassword(username string, password string) User {
	u.Username = username
	u.Password = password
	return u
}
