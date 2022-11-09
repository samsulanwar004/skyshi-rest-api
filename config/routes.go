package config

import (
	"fmt"
	"skyshi-rest-api/controllers"
	"skyshi-rest-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	fmt.Println("Welcome my rest api")

	//migrate
	db.AutoMigrate(
		&models.Activity{},
		&models.Todo{},
	)

	PORT := ":3030"
	router := echo.New()

	// Activity
	handlerActivity := controllers.NewHandlerActivity(db)
	router.GET("/activity-groups", handlerActivity.GetAllActivity)
	router.GET("/activity-groups/:id", handlerActivity.GetActivity)
	router.POST("/activity-groups", handlerActivity.CreateActivity)
	router.PATCH("/activity-groups/:id", handlerActivity.UpdateActivity)
	router.DELETE("/activity-groups/:id", handlerActivity.DeleteActivity)

	// Activity
	handlerTodo := controllers.NewHandlerTodo(db)
	router.GET("/todo-items", handlerTodo.GetAllTodo)
	router.GET("/todo-items/:id", handlerTodo.GetTodo)
	router.POST("/todo-items", handlerTodo.CreateTodo)
	router.PATCH("/todo-items/:id", handlerTodo.UpdateTodo)
	router.DELETE("/todo-items/:id", handlerTodo.DeleteTodo)

	router.Logger.Fatal(router.Start(PORT))
}
