package datastore

import (
	"basic-crud/internal/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	RoleID    int            `json:"role_id"`
	Role      Role           `gorm:"foreignKey:role_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (model *User) List(dbConn *gorm.DB, pagination *models.Pagination) (users *[]User, count int64, err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := dbConn.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	err = queryBuider.Model(&model).Preload("Role").Where(model).Find(&users).Error
	dbConn.Model(&model).Count(&count)

	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (model *User) Get(dbConn *gorm.DB, id int) (user User, err error) {
	err = dbConn.Model(&user).Preload("Role").Where("id = ?", id).First(&user).Error

	return
}

func (model *User) Save(dbConn *gorm.DB, user User) (err error) {
	if dbConn == nil {
		return errors.New("database connection shoul not nil")
	}
	err = dbConn.Create(&user).Error

	if err != nil {
		return err
	}

	return
}

func (model *User) Update(dbConn *gorm.DB, id int, user User) (err error) {
	if dbConn == nil {
		return errors.New("database connection shoul not nil")
	}

	err = dbConn.Model(&model).Where("id = ?", id).Updates(user).Error

	if err != nil {
		return err
	}

	return
}

func (model *User) Delete(dbConn *gorm.DB, id int) (err error) {
	if dbConn == nil {
		return errors.New("database connection shoul not nil")
	}

	err = dbConn.Delete(&model, id).Error

	if err != nil {
		return err
	}

	return
}
