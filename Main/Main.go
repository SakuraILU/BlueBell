package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	// set random seed
	rand.Seed(time.Now().Unix())

	go startServer()
	time.Sleep(3 * time.Second)

	wg := sync.WaitGroup{}
	// before run this test, delete the database file and flushall the redis data
	// and don't use too many goroutine! 5000 goroutine is enough
	// otherwise, database may be locked and cause reading fail -> panic
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			startClient()
			wg.Done()
		}()
	}
	wg.Wait()
}
