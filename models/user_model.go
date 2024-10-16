// package models

//	type Users struct {
//		ID       uint   `gorm:"primaryKey"`
//		Username string `gorm:"unique;not null"`
//		Email    string `gorm:"unique;not null"`
//		Password string `gorm:"not null"`
//		Phone    int    `gorm:"not null"`
//	}
package models

import (
	"gorm.io/gorm"
)

type Users struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
}

// This function can be used to check if the Users table exists
func MigrateUsers(db *gorm.DB) {
	err := db.AutoMigrate(&Users{})
	if err != nil {
		panic("Failed to migrate Users table: " + err.Error())
	}
}
