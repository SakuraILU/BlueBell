package main

import (
	control "bluebell/Control"
	log "bluebell/Log"
	model "bluebell/Model"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var c = &http.Client{}

// listen on port 8080
var addr = "localhost:8080"

func startServer() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.Use(control.JwtAuthorization())
	v1.POST("/signup", control.SignUpHandler)
	v1.POST("/login", control.LoginHandler)
	v1.GET("/community", control.CommunityListHandler)
	v1.GET("/community/:id", control.CommunityDetail)
	// r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Run()
}

func startClient() {
	time.Sleep((time.Duration(rand.Intn(4)+1) * time.Second))

	param_signup := GenerateUserSignUp()
	// signup may fail because of duplicate username, just ignore it
	if !signup(param_signup) {
		return
	}

	time.Sleep((time.Duration(rand.Intn(4)+1) * time.Second))

	param_login := model.ParamLogin{
		Username: param_signup.Username,
		Password: param_signup.Password,
	}

	// get communities
	token, ok := login(param_login)
	if !ok {
		panic("login fail")
	}

	time.Sleep((time.Duration(rand.Intn(4)+1) * time.Second))

	communities, err := GetCommunities(token)
	if err != nil {
		panic(err)
	}

	time.Sleep((time.Duration(rand.Intn(4)+1) * time.Second))

	for _, community := range communities {
		c_detail, err := GetCommunityDetail(token, community.ID)
		if err != nil {
			panic(err)
		}
		log.Infof("community detail: %v", c_detail)
	}
}

func GenerateUserSignUp() model.ParamSignUp {
	// random username and password
	name := ""
	for j := 0; j < 8; j++ {
		name += string(byte('a' + rand.Intn(26)))
	}
	password := ""
	for j := 0; j < 12; j++ {
		password += string(byte('a' + rand.Intn(26)))
	}

	return model.ParamSignUp{
		Username:   name,
		Password:   password,
		RePassword: password,
	}
}

func signup(user model.ParamSignUp) bool {
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
	type ResponseData struct {
		Code int    `json:"Code"`
		Msg  string `json:"msg"`
		Data string `json:"Data"`
	}
	respbody := ResponseData{}
	json.Unmarshal(buf.Bytes(), &respbody)
	token = respbody.Data

	return token, true
}

func loginByToken(token string) bool {
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
	// set random seed
	rand.Seed(time.Now().Unix())

	go startServer()
	time.Sleep(3 * time.Second)

	wg := sync.WaitGroup{}
	// before run this test, delete the database file and flushall the redis data
	// and don't use too many goroutine! 5000 goroutine is enough
	// otherwise, database may be locked and cause reading fail -> panic
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			startClient()
			wg.Done()
		}()
	}
	wg.Wait()
}
