package model

type User struct {
	ID       int64 `gorm:"primaryKey"`
	Username string
	Password string
}
