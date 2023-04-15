package main

import (
	"github.com/keyjin88/todo-app"
	"github.com/keyjin88/todo-app/pkg/handler"
	"github.com/keyjin88/todo-app/pkg/repository"
	"github.com/keyjin88/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.New()
	services := service.New(repos)
	handlers := handler.New(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoute()); err != nil {
		log.Fatalf("error occurred while server running: %s", err.Error())
	}
}
