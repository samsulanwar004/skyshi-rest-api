package repositories

import (
	"skyshi-rest-api/models"

	"gorm.io/gorm"
)

type HandlerActivity struct {
	db *gorm.DB
}

func NewHandlerActivity(db *gorm.DB) *HandlerActivity {
	return &HandlerActivity{db}
}

func (h *HandlerActivity) GetAllActivity(req models.Activity) ([]models.Activity, error) {

	var err error

	var result []models.Activity
	if err := h.db.Find(&result).Error; err != nil {
		return []models.Activity{}, err
	}
	return result, err
}

func (h *HandlerActivity) CreateActivity(req models.Activity) (models.Activity, error) {

	var err error
	err = h.db.Create(&req).Error
	if err != nil {
		return models.Activity{}, err
	}
	return req, nil
}

func (h *HandlerActivity) GetActivity(req models.Activity) (models.Activity, error) {
	var err error
	err = h.db.Model(&req).Where("id = ?", req.ID).Take(&req).Error

	if err != nil {
		return models.Activity{}, err
	}
	return req, err
}

func (h *HandlerActivity) UpdateActivity(req models.Activity) (models.Activity, error) {
	var err error
	err = h.db.Model(&req).Where("id = ?", req.ID).Updates(&req).Take(&req).Error

	if err != nil {
		return models.Activity{}, err
	}
	return req, err
}

func (h *HandlerActivity) DeleteActivity(req models.Activity) (models.ActivityNull, error) {

	db := h.db.Model(&req).Where("id = ?", req.ID).Take(&req).Delete(&req)

	if db.Error != nil {
		return models.ActivityNull{}, db.Error
	}
	return models.ActivityNull{}, nil
}
