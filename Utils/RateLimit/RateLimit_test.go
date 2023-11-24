package ratelimit

import (
	"testing"
	"time"
)

func TestRateLimit(t *testing.T) {
	rate := 15
	nbucket := 100
	limiter := NewRateLimit(rate, nbucket)

	for i := 0; i < 100; i++ {
		for !limiter.Allow() {
			if i%rate != 0 {
				t.Error("TestRateLimit fail")
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
