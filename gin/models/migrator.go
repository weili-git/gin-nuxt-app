package models

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Migrate() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	// AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
	// It will change existing column’s type if its size, precision, nullable changed.
	// It WON’T delete unused columns to protect your data.

	var user_old []struct {
		ID   uint
		Name string
		Age  uint
	}
	db.Table("users").Select("id, name, age").Find(&user_old)

	for _, u := range user_old {
		names := strings.Split(u.Name, " ")
		last_name := names[0]
		first_name := ""
		if len(names) > 1 {
			first_name = names[1]
		}
		email := "Empty"
		birthday := time.Date(time.Now().Year()-int(u.Age), time.January, 1, 0, 0, 0, 0, time.UTC)

		db.Model(&User{}).Where("id=?", u.ID).Updates(User{
			FirstName: first_name,
			LastName:  last_name,
			Email:     email,
			Password:  "12345",
			Birthday:  birthday,
		})
	}
	db.Migrator().DropColumn(&User{}, "Name")
	db.Migrator().DropColumn(&User{}, "Age")
	fmt.Println("Migration is done.")
}
