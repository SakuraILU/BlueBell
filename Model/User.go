package model

type User struct {
	ID       int64  `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

type ParamSignUp struct {
	Username   string
	Password   string
	RePassword string
}

type ParamLogin struct {
	Username string
	Password string
}
