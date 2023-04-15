package service

import (
	"github.com/keyjin88/todo-app/model"
	"github.com/keyjin88/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
