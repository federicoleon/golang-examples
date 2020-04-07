package services

import (
	"github.com/federicoleon/golang-examples/code_structure/mvc/domain"
)

var (
	UsersService = usersService{}
)

type usersService struct{}

func (service usersService) Get(id int64) (*domain.User, error) {
	user := domain.User{Id: id}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (service usersService) Save(user *domain.User) error {
	if err := user.Save(); err != nil {
		return err
	}
	return nil
}
