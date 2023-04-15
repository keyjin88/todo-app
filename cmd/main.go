package main

import (
	"github.com/keyjin88/todo-app"
	"github.com/keyjin88/todo-app/pkg/handler"
	"github.com/keyjin88/todo-app/pkg/repository"
	"github.com/keyjin88/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	repos := repository.New()
	services := service.New(repos)
	handlers := handler.New(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoute()); err != nil {
		log.Fatalf("error occurred while server running: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
