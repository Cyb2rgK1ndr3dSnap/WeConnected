package models

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key"`
	Name       string    `json:"name" gorm:"unique"`
	Created_At time.Time `json:"created_at" gorm:"autoCreateTime"`
}
