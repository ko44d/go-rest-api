package main

import (
	"github.com/ko44d/go-rest-api/controller"
	"github.com/ko44d/go-rest-api/db"
	"github.com/ko44d/go-rest-api/repository"
	"github.com/ko44d/go-rest-api/router"
	"github.com/ko44d/go-rest-api/usecase"
	"github.com/ko44d/go-rest-api/validator"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
