package repository

import (
	"errors"
	"log"

	"github.com/eron97/login-authenticator.git/src/models"
	"gorm.io/gorm"
)

func NewDatabase(db *gorm.DB) Database {
	return &useDatabase{
		db,
	}
}

type Database interface {
	CreateUser(users []models.CreateUser) error
	VerificExist(email string) (bool, error)
	ReadAllUsers() ([]models.GetUsers, error)
}

type useDatabase struct {
	db *gorm.DB
}

func (con *useDatabase) VerificExist(email string) (bool, error) {
	var count int64
	if err := con.db.Model(&models.CreateUser{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, errors.New(err.Error())
	}
	return count != 0, nil
}

func (con *useDatabase) CreateUser(users []models.CreateUser) error {
	return con.db.Create(&users).Error
}

func (con *useDatabase) ReadAllUsers() ([]models.GetUsers, error) {
	var users []models.GetUsers
	if err := con.db.Table("users").Find(&users).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}
