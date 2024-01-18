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
	GetPasswordByEmail(email string) (string, error)
}

type useDatabase struct {
	db *gorm.DB
}

func (con *useDatabase) VerificExist(email string) (bool, error) {
	var count int64
	if err := con.db.Table("users").Where("email = ?", email).Count(&count).Error; err != nil {
		return false, errors.New(err.Error())
	}
	return count != 0, nil
}

func (con *useDatabase) GetPasswordByEmail(email string) (string, error) {
	var user models.LoginUser
	if err := con.db.Table("users").Select("email, password").Where("email = ?", email).Scan(&user).Error; err != nil {
		return "", errors.New(err.Error())
	}
	return user.Password, nil
}

func (con *useDatabase) CreateUser(users []models.CreateUser) error {
	if err := con.db.Table("users").Create(&users).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (con *useDatabase) ReadAllUsers() ([]models.GetUsers, error) {
	var users []models.GetUsers
	if err := con.db.Table("users").Find(&users).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}
