package models

import (
	"time"
)

type User struct {
    ID             string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey;->" json:"id"`
    FirstName      string    `gorm:"not null" json:"first_name"`
    LastName       string    `gorm:"not null" json:"last_name"`
    Password       string    `gorm:"not null" json:"password"`
    Username       string    `gorm:"unique;not null" json:"username"`
    AccountCreated time.Time `gorm:"default:current_timestamp" json:"account_created"`
    AccountUpdated time.Time `gorm:"default:current_timestamp" json:"account_updated"`
}

type PublicUser struct {
    ID             string    `json:"id"`
    FirstName      string    `json:"first_name"`
    LastName       string    `json:"last_name"`
    Username       string    `json:"username"`
    AccountCreated time.Time `json:"account_created"`
    AccountUpdated time.Time `json:"account_updated"`
}