package models

import (
	"time"

	"github.com/google/uuid"
)

type Ofert struct {
	Created_By  uuid.UUID  `json:"created_by" gorm:"primary_key"`
	Update_By   uuid.UUID  `json:"updated_by" gorm:"primary_key"`
	Profession  uint       `json:"profession"`
	Company     uint       `json:"company"`
	Created_At  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	Update_At   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	Finished_At *time.Time `json:"finished_At"`
}
