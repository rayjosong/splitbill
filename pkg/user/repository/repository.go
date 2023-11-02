package repository

import "github.com/rayjosong/splitbill/pkg/user"

type UserRepo struct{}

func NewUserRepo() UserRepo {
	return UserRepo{}
}

func (r UserRepo) InsertUser(user user.User) error {
	return nil
}
