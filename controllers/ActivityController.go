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

type HandlerActivity struct {
	db *gorm.DB
}

func NewHandlerActivity(db *gorm.DB) *HandlerActivity {
	return &HandlerActivity{db}
}

func (h *HandlerActivity) GetAllActivity(c echo.Context) (err error) {
	var req models.Activity
	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	handler := repositories.NewHandlerActivity(h.db)
	result, err := handler.GetAllActivity(req)

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

func (h *HandlerActivity) CreateActivity(c echo.Context) (err error) {
	var req models.Activity
	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	handler := repositories.NewHandlerActivity(h.db)
	result, err := handler.CreateActivity(req)

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

func (h *HandlerActivity) GetActivity(c echo.Context) (err error) {
	var req models.Activity
	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	req.ID, _ = strconv.Atoi(c.Param("id"))
	handler := repositories.NewHandlerActivity(h.db)
	result, err := handler.GetActivity(req)

	if err != nil {
		msg := fmt.Sprintf("Activity with ID %d Not Found", req.ID)
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

func (h *HandlerActivity) UpdateActivity(c echo.Context) (err error) {
	var req models.Activity
	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "Bad Request",
			"message": "title cannot be null",
			"data":    nil,
		})
	}

	req.ID, _ = strconv.Atoi(c.Param("id"))
	handler := repositories.NewHandlerActivity(h.db)
	result, err := handler.UpdateActivity(req)

	if err != nil {
		msg := fmt.Sprintf("Activity with ID %d Not Found", req.ID)
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

func (h *HandlerActivity) DeleteActivity(c echo.Context) (err error) {
	var req models.Activity
	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	req.ID, _ = strconv.Atoi(c.Param("id"))
	handler := repositories.NewHandlerActivity(h.db)
	result, err := handler.DeleteActivity(req)

	if err != nil {
		msg := fmt.Sprintf("Activity with ID %d Not Found", req.ID)
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
