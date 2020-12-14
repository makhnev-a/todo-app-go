package main

import (
	"github.com/makhnev-a/todo-app-go"
	"github.com/makhnev-a/todo-app-go/pkg/handler"
	"github.com/makhnev-a/todo-app-go/pkg/repository"
	"github.com/makhnev-a/todo-app-go/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s\n", err.Error())
	}
}
