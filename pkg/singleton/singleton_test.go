package singleton

import (
	"testing"
	"time"
)

func Test_getInstance(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			if i := getInstance(); i == nil {
				t.Error("instance is nil")
			}
		}()
	}
	time.Sleep(time.Second)
}
