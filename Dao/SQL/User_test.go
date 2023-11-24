package sql

import (
	model "bluebell/Model"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func gnerateUsers(nuser int) (users []model.User) {
	users = make([]model.User, 0)
	for i := 0; i < nuser; i++ {
		users = append(users, model.User{
			Username: fmt.Sprintf("user%d", i),
			Password: fmt.Sprintf("password%d", i),
		})
	}
	return
}

func TestInsertUser(t *testing.T) {
	nuser := 100
	users := gnerateUsers(nuser)

	for _, user := range users {
		err := CreateUser(&user)
		if err != nil {
			t.Error(err.Error())
			return
		}
	}

	for _, user := range users {
		if _, err := GetUserByName(user.Username); err != nil {
			t.Errorf("User %v not exist", user)
			return
		}
	}

	for _, user := range users {
		if _, err := GetUserByName(user.Username); err != nil {
			t.Error(err.Error())
			return
		}
	}

	return
}

// try multi goroutine
func TestInsertUser2(t *testing.T) {
	nuser := 500
	users := gnerateUsers(nuser)

	ngo := 15
	wg := sync.WaitGroup{}
	for i := 0; i < ngo; i++ {
		wg.Add(1)
		go func(i int) {
			for k := i; k < nuser; k += ngo {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				err := CreateUser(&users[k])
				if err != nil {
					t.Error(err.Error())
					return
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	for _, user := range users {
		if _, err := GetUserByName(user.Username); err != nil {
			t.Error(err.Error())
		}
	}

	return
}
