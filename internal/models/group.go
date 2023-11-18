package models

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	GroupID  string    `gorm:"primary_key" json:"group_id"`
	Name     string    `json:"name"`
	Members  []User    `gorm:"many2many:user_groups;"`
	Expenses []Expense `gorm:"foreignkey:GroupID"`
}
