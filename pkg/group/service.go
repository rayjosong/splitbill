package group

import (
	"fmt"

	"github.com/rayjosong/splitbill/internal/models"
)

type GroupRepository interface {
	InsertGroup(models.Group) error
}

type GroupService struct {
	repo GroupRepository
}

func NewGroupService(repo GroupRepository) GroupService {
	return GroupService{repo: repo}
}

func (s GroupService) CreateGroup(group models.Group) error {
	if err := s.repo.InsertGroup(group); err != nil {
		return fmt.Errorf("cannot insert group into repo: %w", err)
	}

	return nil
}
