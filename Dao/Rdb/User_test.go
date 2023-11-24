package rdb

import (
	config "bluebell/Config"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// test SetTokens and GetToken
func TestRedis1(t *testing.T) {
	token_rdb.FlushAll()

	rand.Seed(time.Now().Unix())

	user_id := int64(0)
	ntoken := 100
	tokenstrs := make([]string, 0)
	// random generate num tokens
	for i := 0; i < ntoken; i++ {
		token := ""
		for i := 0; i < 32; i++ {
			token += string(rune(rand.Intn(26) + 97))
		}
		tokenstrs = append(tokenstrs, token)
	}

	// set tokens
	for _, token := range tokenstrs {
		SetToken(user_id, token)
	}

	// get tokens, only NDUPLICATE tokens
	tokens, err := GetTokenStrs(user_id)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(tokens) != config.Cfg.Logic.NDuplicateLogin {
		t.Errorf("GetTokenStrs fail, tokens: %v", tokens)
		return
	}
}

// test multi goroutine SetTokens and GetToken
func TestRedis2(t *testing.T) {
	token_rdb.FlushAll()

	rand.Seed(time.Now().Unix())

	user_id := int64(0)
	ntoken := 100

	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			tokenstrs := make([]string, 0)
			// random generate num tokens
			for i := 0; i < ntoken; i++ {
				token := ""
				for i := 0; i < 32; i++ {
					token += string(rune(rand.Intn(26) + 97))
				}
				tokenstrs = append(tokenstrs, token)
			}

			// set tokens
			for _, token := range tokenstrs {
				SetToken(user_id, token)
			}

			wg.Done()
		}()
	}
	wg.Wait()

	// get tokens, only NDUPLICATE tokens
	tokens, err := GetTokenStrs(user_id)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(tokens) != config.Cfg.Logic.NDuplicateLogin {
		t.Errorf("GetTokenStrs fail, tokens: %v", tokens)
		return
	}
}
