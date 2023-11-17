package main

import (
	control "bluebell/Control"
	model "bluebell/Model"
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// listen on port 8080
var addr = "localhost:8080"

func startServer() {
	r := gin.Default()
	r.POST("/signup", control.SignUpHandler)
	r.POST("/login", control.Login)
	// r.ForwardedByClientIP = true
	// r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Run()
}

func startClient() {
	// register users
	users := []model.ParamSignUp{
		{
			Username:   "user1",
			Password:   "123456",
			RePassword: "123456",
		},
		{
			Username:   "user2",
			Password:   "32142546",
			RePassword: "32142546",
		},
		// user2 already exist
		{
			Username:   "user2",
			Password:   "321rg46",
			RePassword: "321rg46",
		},
		// invalid password
		{
			Username:   "user3",
			Password:   "2546",
			RePassword: "g3132142546",
		},
	}

	// start a client and send requests
	c := &http.Client{}

	// signup
	for _, user := range users {
		body, _ := json.Marshal(user)
		req, err := http.NewRequest("POST", "http://localhost:8080/signup", bytes.NewReader(body))
		if err != nil {
			panic(err)
		}
		resp, err := c.Do(req)
		if err != nil {
			panic(err)
		}
		// print response
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		respBody := buf.String()
		println(respBody)
	}

	// login
	loginUsers := []model.ParamLogin{
		{
			Username: users[0].Username,
			Password: users[0].Password,
		},
		{
			Username: users[1].Username,
			Password: users[1].Password,
		},
		// invalid password
		{
			Username: users[0].Username,
			Password: users[0].Password + "123",
		},
		// invalid username
		{
			Username: "user1000",
			Password: "12345656543342t",
		},
	}

	for i, user := range loginUsers {
		body, _ := json.Marshal(user)
		req, err := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewReader(body))
		// 0,1 success, 2,3 fail
		if i < 2 {
			if err != nil {
				panic(err)
			}
			resp, err := c.Do(req)
			if err != nil {
				panic(err)
			}
			// print response
			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			respBody := buf.String()
			println(respBody)
		} else {
			if err != nil {
				panic(err)
			}
			resp, err := c.Do(req)
			if err != nil {
				panic(err)
			}
			// print response
			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			respBody := buf.String()
			println(respBody)
		}
	}

}

func main() {
	go startServer()
	time.Sleep(1 * time.Second)
	startClient()
}
