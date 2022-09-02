package datastore

import (
	"basic-crud/internal/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (model *Product) List(dbConn *gorm.DB, pagination *models.Pagination) (products *[]Product, count int64, err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := dbConn.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	err = queryBuider.Model(&model).Where(model).Find(&products).Error
	dbConn.Model(&model).Count(&count)

	if err != nil {
		return nil, 0, err
	}

	return products, count, nil
}

func (model *Product) Get(dbConn *gorm.DB, id int) (product Product, err error) {
	err = dbConn.Model(&product).Where("id = ?", id).First(&product).Error

	return
}

func (model *Product) Save(dbConn *gorm.DB, product Product) (err error) {
	if dbConn == nil {
		return errors.New("database connection shoul not nil")
	}
	err = dbConn.Create(&product).Error

	if err != nil {
		return err
	}

	return
}

func (model *Product) Update(dbConn *gorm.DB, id int, product Product) (err error) {
	if dbConn == nil {
		return errors.New("database connection shoul not nil")
	}

	err = dbConn.Model(&model).Where("id = ?", id).Updates(product).Error

	if err != nil {
		return err
	}

	return
}

func (model *Product) Delete(dbConn *gorm.DB, id int) (err error) {
	if dbConn == nil {
		return errors.New("database connection shoul not nil")
	}

	err = dbConn.Delete(&model, id).Error

	if err != nil {
		return err
	}

	return
}
