package cookie

import (
	model "bluebell/Model"
	"math/rand"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	users := GenerateUsers(100)

	for _, user := range users {
		token, err := GetToken(&user)
		if err != nil {
			t.Error(err.Error())
			return
		}

		parse_user, err := ParseToken(token)
		if err != nil {
			t.Error(err.Error())
			return
		}

		if parse_user.ID != user.ID || parse_user.Username != user.Username {
			t.Errorf("ParseToken fail, parse_user: %v, user: %v", parse_user, user)
			return
		}
	}

}

func GenerateUsers(nuser int) (users []model.User) {
	// set random seed
	rand.Seed(time.Now().Unix())

	users = make([]model.User, 0)

	for i := 0; i < nuser; i++ {
		// random username and password
		name := ""
		for j := 0; j < 10; j++ {
			name += string(byte('a' + rand.Intn(26)))
		}

		password := ""
		for j := 0; j < 15; j++ {
			password += string(byte('a' + rand.Intn(26)))
		}

		users = append(users, model.User{
			ID:       int64(i),
			Username: name,
			Password: password,
		})
	}

	return
}
