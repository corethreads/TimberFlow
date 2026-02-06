package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Business_ID   string `gorm:"type:uuid;unique"`
	Business_Name string `gorm:"uniqueIndex;not null" `
	Username      string `gorm:"uniqueIndex;not null"`
	Email         string `gorm:"unique" `
	Password      string `gorm:"not null" `
	Role          string `gorm:"default:admin" `
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}
