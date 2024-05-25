package main

import (
	"github.com/ko44d/go-rest-api/controller"
	"github.com/ko44d/go-rest-api/db"
	"github.com/ko44d/go-rest-api/repository"
	"github.com/ko44d/go-rest-api/router"
	"github.com/ko44d/go-rest-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
