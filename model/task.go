package model

import "time"

// タスクのモデル構造
type Task struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// user.goのstruct参照
	User User `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"` //user削除時に紐づくタスクを削除
	UserId uint `json:"user_id" form:"not null"` //user.goのUserのIDが格納される
}

// GETメソッドでリクエストがあったときに返すタスクのデータ構造
type TaskResponse struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
