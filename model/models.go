package model

import (
	"time"
)

// User ユーザTBLの構造体
type User struct {
	ID         string `json:"id" gorm:"column:id" v-post:"required" validate:"required"`
	Name       string `json:"username" gorm:"column:name" v-post:"required" validate:"required"`
	FirstName  string `json:"firstName" gorm:"column:first_name"`
	LastName   string `json:"lastName" gorm:"column:last_name"`
	Email      string `json:"Email" gorm:"column:email"`
	Password   string `json:"Password" gorm:"column:password"`
	Phone      string `json:"Phone" gorm:"column:phone"`
	UserStatus int    `json:"UserStatus" gorm:"column:userStatus"`
	Version    int    `json:"Version" gorm:"column:version"`
}

// Todos Todoのリスト
type Todos struct {
	Todo []Todo `validate:"dive"`
}

// Todo TodoTBLの構造体
type Todo struct {
	ID          string     `json:"id" gorm:"column:id" v-post:"required" validate:"required"`
	UserID      string     `json:"owner" gorm:"column:user_id" v-post:"required" validate:"required"`
	Title       string     `json:"title" gorm:"column:title"`
	Contents    string     `json:"contents" gorm:"column:contents"`
	Start       *time.Time `json:"start" gorm:"column:start"`
	Due         *time.Time `json:"due" gorm:"column:due"`
	ActualStart *time.Time `json:"actualStart" gorm:"column:actualStart"`
	ActualEnd   *time.Time `json:"actualEnd" gorm:"column:actualEnd"`
	Status      *int       `json:"status" gorm:"column:status"`
	Version     int        `json:"version" gorm:"column:version"`
}
