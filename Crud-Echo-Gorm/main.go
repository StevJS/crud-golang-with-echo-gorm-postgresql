package main

import (
	"Crud-Echo-Gorm/configuration"
	"Crud-Echo-Gorm/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//e.Group("/user")
	e.POST("/createUser", controller.CreateUser)
	e.GET("/users", controller.GetAllUsers)
	e.GET("/user/:id", controller.GetUserById)
	e.PUT("updateUser/:id", controller.UpdateUser)
	e.DELETE("/deleteUser/:id", controller.DeleteUserById)

	//connection database
	configuration.DatabaseInit()
	gorm := configuration.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	//Start server
	e.Logger.Fatal(e.Start(":8080"))
}
