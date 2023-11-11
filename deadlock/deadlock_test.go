package main

import (
	"sync"
	"testing"
	"time"
)

func TestCannotUpdateInLockedState(t *testing.T) {
	num := Number{
		val: 1,
		mutex: sync.Mutex{},
	}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		num.Lock()
		defer func() {
			wg.Done()
		}()
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(2 * time.Second)
	num.UpdateVal(2)

	wg.Wait()

	if num.val != 1 {
		t.Errorf("error_locking_number|expected: %d|actual :%d", 1, num.val)
	}
}
