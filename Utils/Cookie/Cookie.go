package cookie

import (
	log "bluebell/Log"
	model "bluebell/Model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	ExpireTime = 2 * time.Hour
)

var (
	salt string = "Genshin Impact"
)

type Claims struct {
	jwt.StandardClaims
	UserID   int64
	Username string
}

func GetToken(user *model.User) (string, error) {
	c := Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Bluebell",
			ExpiresAt: time.Now().Add(ExpireTime).Unix(),
		},
		UserID:   user.ID,
		Username: user.Password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// note: should pass []byte to SignedString, not string...
	return token.SignedString([]byte(salt))
}

func ParseToken(token_str string) (user *model.User, err error) {
	c := &Claims{}

	log.Errorf("parse token: %v", token_str)
	if _, err = jwt.ParseWithClaims(token_str, c, func(t *jwt.Token) (interface{}, error) {
		// note: salt is []byte, not string... horrible bug...
		return []byte(salt), nil
	}); err != nil {
		return
	}

	user = &model.User{
		ID:       c.UserID,
		Username: c.Username,
	}

	return user, nil
}
