package entities

import "time"

type User struct {
	Id        string    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required"`
	Password  string    `json:"password" db:"password" binding:"required"`
	Avatar    string    `json:"avatar" db:"avatar"`
	Banned    bool      `json:"banned" db:"banned"`
	BanReason string    `json:"banReason" db:"ban_reason"`
	Status    string    `json:"status" db:"status"`
	Followers []string  `json:"followers" db:"followers"`
	Roles     []string  `json:"roles" db:"roles"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
