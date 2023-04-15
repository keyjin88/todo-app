package main

import (
	"github.com/keyjin88/todo-app"
	"github.com/keyjin88/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoute()); err != nil {
		log.Fatalf("error occurred while server running: %s", err.Error())
	}
}
