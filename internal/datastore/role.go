package datastore

import (
	"basic-crud/internal/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (model *Role) List(dbConn *gorm.DB, pagination *models.Pagination) (roles *[]Role, count int64, err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := dbConn.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	err = queryBuider.Model(&model).Where(model).Find(&roles).Error
	dbConn.Model(&model).Count(&count)

	if err != nil {
		return nil, 0, err
	}

	return roles, count, nil
}

func (model *Role) Get(dbConn *gorm.DB, id int) (role Role, err error) {
	err = dbConn.Model(&role).Where("id = ?", id).First(&role).Error

	return
}

func (model *Role) Save(dbConn *gorm.DB, role Role) (err error) {
	if dbConn == nil {
		return errors.New("database connection shoul not nil")
	}
	err = dbConn.Create(&role).Error

	if err != nil {
		return err
	}

	return
}

func (model *Role) Update(dbConn *gorm.DB, id int, role Role) (err error) {
	if dbConn == nil {
		return errors.New("database connection shoul not nil")
	}

	err = dbConn.Model(&model).Where("id = ?", id).Updates(role).Error

	if err != nil {
		return err
	}

	return
}

func (model *Role) Delete(dbConn *gorm.DB, id int) (err error) {
	if dbConn == nil {
		return errors.New("database connection shoul not nil")
	}

	err = dbConn.Delete(&model, id).Error

	if err != nil {
		return err
	}

	return
}
