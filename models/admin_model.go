package models

import "gorm.io/gorm"

type Admins struct {
	gorm.Model
	Email    string `json:"username"`
	Password string `json:"password"`
}
	