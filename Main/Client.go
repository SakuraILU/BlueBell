package main

import (
	log "bluebell/Log"
	model "bluebell/Model"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var httpaddr = "http://" + serveraddr
var c = &http.Client{}

func startClient() {
	time.Sleep((time.Duration(rand.Intn(2)) * time.Second))

	param_signup := GenerateUserSignUp()
	// signup may fail because of duplicate username, just ignore it
	if !signup(param_signup) {
		return
	}

	time.Sleep((time.Duration(rand.Intn(2)) * time.Second))

	param_login := model.ParamLogin{
		Username: param_signup.Username,
		Password: param_signup.Password,
	}

	// get communities
	token, ok := login(param_login)
	if !ok {
		panic("login fail")
	}

	time.Sleep((time.Duration(rand.Intn(5)) * time.Second))

	communities, err := GetCommunities(token)
	if err != nil {
		panic(err)
	}

	time.Sleep((time.Duration(rand.Intn(2)) * time.Second))

	for _, community := range communities {
		c_detail, err := GetCommunityDetail(token, community.ID)
		if err != nil {
			panic(err)
		}
		log.Infof("community detail: %v", c_detail)
	}

	// time.Sleep((time.Duration(rand.Intn(1)) * time.Second))

	// create post
	for i := 0; i < 20; i++ {
		post := GeneratePost()
		if !CreatePost(token, post) {
			panic("create post fail")
		}
	}

	time.Sleep((time.Duration(rand.Intn(3)) * time.Second))

	// get posts
	choice := model.TIME
	if rand.Intn(2) == 0 {
		choice = model.SCORE
	}
	post_details, err := GetPosts(token, 1, 1, 5, choice)
	if err != nil {
		panic(err)
	}

	time.Sleep((time.Duration(rand.Intn(2)) * time.Second))

	// vote for posts
	for i := 0; i < 5; i++ {
		vote := model.ParamVote{
			PostID: post_details[i].Post.ID,
			Choice: 1,
		}
		if !VoteForPost(token, vote) {
			panic("vote fail")
		}
	}

	// signup may fail because of duplicate username, just ignore it
	param_signup = GenerateUserSignUp()
	if !signup(param_signup) {
		return
	}

	time.Sleep((time.Duration(rand.Intn(4)) * time.Second))

	param_login = model.ParamLogin{
		Username: param_signup.Username,
		Password: param_signup.Password,
	}

	// get communities
	token, ok = login(param_login)
	if !ok {
		panic("login fail")
	}

	for i := 0; i < 3; i++ {
		vote := model.ParamVote{
			PostID: post_details[i].Post.ID,
			Choice: -1,
		}
		if !VoteForPost(token, vote) {
			panic("vote fail")
		}
	}

	// signup may fail because of duplicate username, just ignore it
	param_signup = GenerateUserSignUp()
	if !signup(param_signup) {
		return
	}

	time.Sleep((time.Duration(rand.Intn(4)) * time.Second))

	param_login = model.ParamLogin{
		Username: param_signup.Username,
		Password: param_signup.Password,
	}

	// get communities
	token, ok = login(param_login)
	if !ok {
		panic("login fail")
	}

	for i := 0; i < 2; i++ {
		vote := model.ParamVote{
			PostID: post_details[i].Post.ID,
			Choice: 1,
		}
		if !VoteForPost(token, vote) {
			panic("vote fail")
		}
	}

	for i := 4; i < 5; i++ {
		vote := model.ParamVote{
			PostID: post_details[i].Post.ID,
			Choice: 1,
		}
		if !VoteForPost(token, vote) {
			panic("vote fail")
		}
	}

	time.Sleep((time.Duration(rand.Intn(4)) * time.Second))

	// get posts
	post_details, err = GetPosts(token, 1, 1, 6, model.SCORE)
	if err != nil {
		panic(err)
	}
	for _, post_detail := range post_details {
		log.Infof("[%d] post: title %v, community %v, vote %v", post_detail.Post.ID, post_detail.Post.Title, post_detail.CommunityDetail.Name, post_detail.NVote)
	}

	time.Sleep((time.Duration(rand.Intn(3)) * time.Second))

	// get post
	post_detail, err := GetPost(token, post_details[len(post_details)-1].Post.ID)
	log.Infof("post: name %v, title %v, %d, content %s", post_detail.AuthorName, post_detail.Post.Title, post_detail.NVote, post_detail.Post.Content)

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

func GeneratePost() model.ParamPost {
	// random username and password
	title := ""
	for j := 0; j < 1+rand.Intn(10); j++ {
		title += string(byte('a' + rand.Intn(26)))
	}
	content := ""
	for j := 0; j < 524; j++ {
		content += string(byte('a' + rand.Intn(26)))
	}

	return model.ParamPost{
		Title:       title,
		Content:     content,
		CommunityID: int64(1 + rand.Intn(4)),
	}
}

func signup(user model.ParamSignUp) bool {
	body, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/signup", httpaddr), bytes.NewReader(body))
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
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/login", httpaddr), bytes.NewReader(body))
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
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/login", httpaddr), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := c.Do(req)

	// check statusok?
	return resp.StatusCode == http.StatusOK
}

func GetCommunities(token string) ([]model.ParamCommunity, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/community", httpaddr), nil)
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
		Code int                    `json:"Code"`
		Msg  string                 `json:"msg"`
		Data []model.ParamCommunity `json:"Data"`
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

func GetCommunityDetail(token string, id int64) (model.ParamCommunityDetail, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/community/%d", httpaddr, id), nil)
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
		Code int                        `json:"Code"`
		Msg  string                     `json:"msg"`
		Data model.ParamCommunityDetail `json:"Data"`
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

func CreatePost(token string, post model.ParamPost) bool {
	data, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/post", httpaddr), bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := c.Do(req)

	// check statusok?
	return resp.StatusCode == http.StatusOK
}

func GetPosts(taken string, cid int, page, size int, order string) ([]*model.ParamPostDetail, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/posts?cid=%d&page=%d&size=%d&order=%s", httpaddr, cid, page, size, order), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+taken)

	resp, err := c.Do(req)
	// get communities
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	type ResponseData struct {
		Code int                      `json:"Code"`
		Msg  string                   `json:"msg"`
		Data []*model.ParamPostDetail `json:"Data"`
	}

	respbody := ResponseData{}
	json.Unmarshal(buf.Bytes(), &respbody)
	posts := respbody.Data

	// check statusok?
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return posts, nil
}

func GetPost(token string, id int64) (post *model.ParamPostDetail, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/post/%d", httpaddr, id), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.Do(req)
	if err != nil {
		return
	}

	// get post
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	type ResponseData struct {
		Code int                    `json:"Code"`
		Msg  string                 `json:"msg"`
		Data *model.ParamPostDetail `json:"Data"`
	}
	respbody := ResponseData{}
	json.Unmarshal(buf.Bytes(), &respbody)

	// check statusok?
	if resp.StatusCode != http.StatusOK {
		return
	}

	post = respbody.Data
	return
}

func VoteForPost(token string, vote model.ParamVote) bool {
	data, err := json.Marshal(vote)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/vote", httpaddr), bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := c.Do(req)

	// check statusok?
	return resp.StatusCode == http.StatusOK
}
