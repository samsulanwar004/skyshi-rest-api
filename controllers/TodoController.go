package controllers

import (
	"fmt"
	"net/http"
	"skyshi-rest-api/models"
	"skyshi-rest-api/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HandlerTodo struct {
	db *gorm.DB
}

func NewHandlerTodo(db *gorm.DB) *HandlerTodo {
	return &HandlerTodo{db}
}

func (h *HandlerTodo) GetAllTodo(c echo.Context) (err error) {
	var req models.Todo

	req.ActivityGroupID, _ = strconv.Atoi(c.QueryParam("activity_group_id"))

	handler := repositories.NewHandlerTodo(h.db)
	result, err := handler.GetAllTodo(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "Success",
		"message": "Success",
		"data":    result,
	})
}

func (h *HandlerTodo) CreateTodo(c echo.Context) (err error) {
	var req models.Todo

	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "Bad Request",
			"message": "title cannot be null",
			"data":    nil,
		})
	}

	if req.ActivityGroupID == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "Bad Request",
			"message": "activity_group_id cannot be null",
			"data":    nil,
		})
	}

	if req.Priority == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "Bad Request",
			"message": "priority cannot be null",
			"data":    nil,
		})
	}

	handler := repositories.NewHandlerTodo(h.db)
	result, err := handler.CreateTodo(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status":  "Success",
		"message": "Success",
		"data":    result,
	})
}

func (h *HandlerTodo) GetTodo(c echo.Context) (err error) {
	var req models.Todo

	req.ID, _ = strconv.Atoi(c.Param("id"))
	handler := repositories.NewHandlerTodo(h.db)
	result, err := handler.GetTodo(req)

	if err != nil {
		msg := fmt.Sprintf("Todo with ID %d Not Found", req.ID)
		return c.JSON(http.StatusNotFound, echo.Map{
			"status":  "Not Found",
			"message": msg,
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "Success",
		"message": "Success",
		"data":    result,
	})
}

func (h *HandlerTodo) UpdateTodo(c echo.Context) (err error) {
	var req models.Todo

	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "Bad Request",
			"message": "title cannot be null",
			"data":    nil,
		})
	}

	req.ID, _ = strconv.Atoi(c.Param("id"))
	handler := repositories.NewHandlerTodo(h.db)
	result, err := handler.UpdateTodo(req)

	if err != nil {
		msg := fmt.Sprintf("Todo with ID %d Not Found", req.ID)
		return c.JSON(http.StatusNotFound, echo.Map{
			"status":  "Not Found",
			"message": msg,
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "Success",
		"message": "Success",
		"data":    result,
	})
}

func (h *HandlerTodo) DeleteTodo(c echo.Context) (err error) {
	var req models.Todo

	req.ID, _ = strconv.Atoi(c.Param("id"))
	handler := repositories.NewHandlerTodo(h.db)
	result, err := handler.DeleteTodo(req)

	if err != nil {
		msg := fmt.Sprintf("Todo with ID %d Not Found", req.ID)
		return c.JSON(http.StatusNotFound, echo.Map{
			"status":  "Not Found",
			"message": msg,
			"data":    result,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "Success",
		"message": "Success",
		"data":    result,
	})
}
