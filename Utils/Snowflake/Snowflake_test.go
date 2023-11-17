package snowflake

import (
	"fmt"
	"testing"
	"time"
)

// single goroutine
func TestSnowflake1(t *testing.T) {
	start, err := time.Parse("2006-01-02 15:04:05", "2020-01-01 00:00:00")
	if err != nil {
		t.Error(err)
		return
	}

	id_generator, err := NewSnowflake(start, 1, 1)
	if err != nil {
		t.Error(err)
		return
	}

	ids := make([]int64, 0)
	for i := 0; i < 10; i++ {
		ids = append(ids, id_generator.NextID())
	}

	// no duplicate id
	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			if ids[i] == ids[j] {
				t.Errorf("duplicate id: %d", ids[i])
				return
			}
		}
	}
}

// multiple goroutine
func TestSnowflake2(t *testing.T) {
	start, err := time.Parse("2006-01-02 15:04:05", "2020-01-01 00:00:00")
	if err != nil {
		t.Error(err)
		return
	}

	id_generator, err := NewSnowflake(start, 1, 1)
	if err != nil {
		t.Error(err)
		return
	}

	ids := make([]int64, 0)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				id := id_generator.NextID()
				ids = append(ids, id)
				// print in binary
				fmt.Printf("%064b\n", id)
			}
		}()
	}

	// no duplicate id
	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			if ids[i] == ids[j] {
				t.Errorf("duplicate id: %d", ids[i])
				return
			}
		}
	}
}
