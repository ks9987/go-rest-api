package model

import "time"

type User struct {
	ID        uint      `JSON:"id" gorm:"primary_key"`
	Email     string    `JSON:"email" gorm:"unique"`
	Password  string    `JSON:"password"`
	CreatedAt time.Time `JSON:"created_at"`
	UpdatedAt time.Time `JSON:"updated_at"`
}

type UserResponse struct {
	ID    uint   `JSON:"id" gorm:"primary_key"`
	Email string `JSON:"email" gorm:"unique"`
}
