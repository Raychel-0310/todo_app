package model

import "time"

type User struct {
	ID uint `json: "id" gorm: "primaryKey"`
	// jsonに変換したときに自動的にIDをidにしてくれる
	Email string `json: "email" gorm: "unique"`
	Password string `json: "password"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

// サインアップ時の新しいユーザーの情報
type UserResponse struct {
	ID uint `json: "id" gorm: "primaryKey"`
	Email string `json: "email" gorm: "unique"`
}