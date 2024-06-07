package models

import (
	"time"

	"example.com/gin/utils"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Birthday  time.Time `json:"birthday"`
}

type User_old struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func (m *User) BeforeCreate(orm *gorm.DB) error { // hook ?
	hash, err := utils.Encrypt(m.Password)
	if err == nil {
		m.Password = hash
	}
	return err
}

func GetUserInfoByEmail(email string) (User, error) {
	var user User
	if err := DB.Where("email=?", email).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
