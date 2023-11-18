package models

import "github.com/jinzhu/gorm"

type Share struct {
	gorm.Model
	ExpenseID  string  `json:"expense_id"`
	FromUserID string  `json:"from_user_id"`
	ToUserID   string  `json:"to_user_id"`
	Amount     float32 `json:"amount"`
}
