package main

import (
	control "bluebell/Control"
	model "bluebell/Model"
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// listen on port 8080
var addr = "localhost:8080"

func startServer() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.Use(control.JwtAuthorization())
	v1.POST("/signup", control.SignUpHandler)
	v1.POST("/login", control.Login)
	// r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Run()
}

func startClient() {
	// set random seed
	rand.Seed(time.Now().Unix())

	nuser := 10
	paramSignups := make([]model.ParamSignUp, 0)
	paramLogins := make([]model.ParamLogin, 0)
	for i := 0; i < nuser; i++ {
		// random username and password
		name := ""
		for j := 0; j < 10; j++ {
			name += string('a' + rand.Intn(26))
		}
		password := ""
		for j := 0; j < 15; j++ {
			password += string('a' + rand.Intn(26))
		}

		paramSignups = append(paramSignups, model.ParamSignUp{
			Username:   name,
			Password:   password,
			RePassword: password,
		})
		paramLogins = append(paramLogins, model.ParamLogin{
			Username: name,
			Password: password,
		})
	}

	for _, param := range paramSignups {
		signup(param)
	}

	// user[0] login 4 times
	tokens := make([]string, 0)
	for i := 0; i < 3; i++ {
		token, ok := login(paramLogins[0])
		if !ok {
			panic("login fail")
		}
		tokens = append(tokens, token)
	}

	// now the first token is invalid
	if !loginByToken(tokens[0]) {
		panic("token should be invalid")
	}
}

func signup(user model.ParamSignUp) bool {
	c := &http.Client{}
	body, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/signup", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return false
	}

	// check statusok?
	if resp.StatusCode != 200 {
		return false
	}
	return true
}

func login(user model.ParamLogin) (token string, ok bool) {
	c := &http.Client{}
	body, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/login", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		return "", false
	}
	if resp.StatusCode != 200 {
		return "", false
	}

	// token 在body中的Data字段
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respbody := make(map[string]string)
	json.Unmarshal(buf.Bytes(), &respbody)
	token = respbody["Data"]

	return token, true
}

func loginByToken(token string) bool {
	c := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/login", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := c.Do(req)

	// check statusok?
	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func main() {
	go startServer()
	time.Sleep(3 * time.Second)

	startClient()
}
