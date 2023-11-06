package group

import "fmt"

type GroupRepository interface {
	InsertGroup(Group) error
}

type GroupService struct {
	repo GroupRepository
}

func NewGroupService(repo GroupRepository) GroupService {
	return GroupService{repo: repo}
}

func (s GroupService) CreateGroup(group Group) error {
	if err := s.repo.InsertGroup(group); err != nil {
		return fmt.Errorf("cannot insert group into repo: %w", err)
	}

	return nil
}
