package model

import "time"

type Task struct {
	ID        uint      `JSON:"id" gorm:"primary_key"`
	Title     string    `JSON:"title" gorm:"not null"`
	CreatedAt time.Time `JSON:"created_at"`
	UpdatedAt time.Time `JSON:"updated_at"`
	User      User      `JSON:"user" gorm:"foreignkey:UserID; constraint:OnDelete:CASCADE"`
	UserID    uint      `JSON:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID        uint      `JSON:"id" gorm:"primary_key"`
	Title     string    `JSON:"title" gorm:"not null"`
	CreatedAt time.Time `JSON:"created_at"`
	UpdatedAt time.Time `JSON:"updated_at"`
}
