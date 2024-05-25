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
	usecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(usecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
