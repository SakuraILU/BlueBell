package main

import (
	control "bluebell/Control"
	model "bluebell/Model"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
	v1.GET("/community", control.CommunityListHandler)
	v1.GET("/community/:id", control.CommunityDetail)
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

	// test multiple login
	// // user[0] login 4 times
	// tokens := make([]string, 0)
	// for i := 0; i < 4; i++ {
	// 	token, ok := login(paramLogins[0])
	// 	if !ok {
	// 		panic("login fail")
	// 	}
	// 	time.Sleep(5 * time.Second)
	// 	tokens = append(tokens, token)
	// }

	// // now the first token is invalid
	// if loginByToken(tokens[0]) == true {
	// 	panic("token should be invalid...")
	// }

	// get communities
	token, ok := login(paramLogins[0])
	if !ok {
		panic("login fail")
	}

	communities, err := GetCommunities(token)
	if err != nil {
		panic(err)
	}
	for _, community := range communities {
		c_detail, err := GetCommunityDetail(token, community.ID)
		if err != nil {
			panic(err)
		}

		log.Printf("community %v detail: %v", community, c_detail)
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
	return resp.StatusCode == http.StatusOK
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
	return resp.StatusCode == http.StatusOK
}

func GetCommunities(token string) ([]model.ParamCommity, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/community", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := c.Do(req)
	// get communities
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	type ResponseData struct {
		Code int                  `json:"Code"`
		Msg  string               `json:"msg"`
		Data []model.ParamCommity `json:"Data"`
	}

	respbody := ResponseData{}
	json.Unmarshal(buf.Bytes(), &respbody)
	communities := respbody.Data

	// check statusok?
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return communities, nil
}

func GetCommunityDetail(token string, id int64) (model.ParamCommityDetail, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/community/"+fmt.Sprint(id), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := c.Do(req)
	// get communities
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	type ResponseData struct {
		Code int                      `json:"Code"`
		Msg  string                   `json:"msg"`
		Data model.ParamCommityDetail `json:"Data"`
	}

	respbody := ResponseData{}
	json.Unmarshal(buf.Bytes(), &respbody)
	community := respbody.Data

	// check statusok?
	if resp.StatusCode != http.StatusOK {
		return community, err
	}

	return community, nil
}

func main() {
	go startServer()
	time.Sleep(3 * time.Second)

	startClient()
}
