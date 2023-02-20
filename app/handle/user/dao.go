package user

import (
	"time"

	"github.com/gin-qt-business/app/config"
)

const TABLE_NAME = "users"

type User struct {
	ID        uint `gorm:"primary_key"`
	Uid       string
	Username  string
	Password  string
	Phone     string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt uint
}

func (user *User) AddUser() error {
	err := config.DB.Table(TABLE_NAME).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *User) GetUser() error {
	err := config.DB.Table(TABLE_NAME).Take(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *User) CountUser() (uint, error) {
	var count uint
	err := config.DB.Table(TABLE_NAME).Where(user).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
