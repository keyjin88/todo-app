package service

import "github.com/keyjin88/todo-app/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
