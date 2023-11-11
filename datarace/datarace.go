package main

import (
	"fmt"
	"sync"
)

//Run with the -race flag to find the data race condition.

func main() {
	data := make(map[string]string)
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	wg.Add(1)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		data["a"] = "a"
		defer wg.Done()
	}()
	lock.Lock()
	data["a"] = "b"
	lock.Unlock()


	fmt.Printf("%v", data)
	wg.Wait()
}

