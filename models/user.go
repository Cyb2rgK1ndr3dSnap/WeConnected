package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID `json:"id" gorm:"primary_key"`
	Name        string    `json:"username"`
	Lastname    string    `json:"lastname"`
	Email       string    `json:"email" gorm:"unique"`
	Password    string    `json:"password"`
	Phonenumber string    `json:"phoneNumber"`
	Role        uint      `json:"role"`
	Age         uint      `json:"age"`
	Created_At  time.Time `json:"created_at" gorm:"autoCreateTime"`
	Update_At   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
