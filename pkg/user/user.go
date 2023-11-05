package user

import (
	"time"

	"gorm.io/gorm"
)

type UserCredentials struct {
	gorm.Model
	UserID   uint   `gorm:"unique;not null"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	User     User   `gorm:"foreignkey:UserID"`
}

type User struct {
	gorm.Model
	UserID   uint      `gorm:"primary_key" json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `gorm:"type:varchar(100);unique_index" json:"email"`
	Phone    string    `json:"phone,omitempty"`
	Balance  float32   `json:"balance"`
	Expenses []Expense `gorm:"foreignkey:CreatorID"`
	Groups   []Group   `gorm:"many2many:user_groups;"`
	Shares   []Share   `gorm:"foreignkey:FromUserID"`
}

type Expense struct {
	gorm.Model
	ExpenseID   string    `gorm:"primary_key" json:"expense_id"`
	Description string    `json:"description"`
	Amount      float32   `json:"amount"`
	Date        time.Time `json:"date"`
	CreatorID   string    `json:"creator_id"`
	GroupID     string    `json:"group_id"`
	Shares      []Share   `gorm:"foreignkey:ExpenseID"`
}

type Group struct {
	gorm.Model
	GroupID  string    `gorm:"primary_key" json:"group_id"`
	Name     string    `json:"name"`
	Members  []User    `gorm:"many2many:user_groups;"`
	Expenses []Expense `gorm:"foreignkey:GroupID"`
}

type Share struct {
	gorm.Model
	ShareID    string  `gorm:"primary_key" json:"share_id"`
	ExpenseID  string  `json:"expense_id"`
	FromUserID string  `json:"from_user_id"`
	ToUserID   string  `json:"to_user_id"`
	Amount     float32 `json:"amount"`
}
