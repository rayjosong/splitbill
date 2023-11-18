package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Expense struct {
	gorm.Model
	Description string    `json:"description"`
	Amount      float32   `json:"amount"`
	Date        time.Time `json:"date"`
	CreatorID   string    `json:"creator_id"`
	GroupID     string    `json:"group_id"`
	Shares      []Share   `gorm:"foreignkey:ExpenseID"`
}
