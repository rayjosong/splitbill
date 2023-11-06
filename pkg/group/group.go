package group

import (
	"github.com/jinzhu/gorm"
	"github.com/rayjosong/splitbill/pkg/user"
)

type Group struct {
	gorm.Model
	GroupID  string         `gorm:"primary_key" json:"group_id"`
	Name     string         `json:"name"`
	Members  []user.User    `gorm:"many2many:user_groups;"`
	Expenses []user.Expense `gorm:"foreignkey:GroupID"`
}
