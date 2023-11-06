package repository

import (
	"fmt"

	"github.com/rayjosong/splitbill/pkg/group"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return GroupRepository{db: db}
}

func (r GroupRepository) CreateGroup(group group.Group) error {
	if err := r.db.Create(&group).Error; err != nil {
		return fmt.Errorf("error inserting into db: %w", err)
	}

	return nil
}
