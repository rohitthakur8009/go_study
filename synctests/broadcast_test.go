package snyctests

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Button struct {
	Clicked *sync.Cond
}

func TestBroadCast(t *testing.T) {

	button := Button{ Clicked: sync.NewCond(&sync.Mutex{}) }
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()

			c.Wait()
			fn() }()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialogue box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})
	button.Clicked.Broadcast()
	clickRegistered.Wait()
}

func TestBroadCastNew(t *testing.T) {
	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	clickSubscribe := sync.WaitGroup{}
	clickSubscribe.Add(3)
	go func(c *sync.Cond) {
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
		fmt.Println("Button Clicked")
		clickSubscribe.Done()
	}(button.Clicked)

	go func(c *sync.Cond) {
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
		fmt.Println("Call API")
		clickSubscribe.Done()
	}(button.Clicked)

	go func(c *sync.Cond) {
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
		fmt.Println("Render Page")
		clickSubscribe.Done()
	}(button.Clicked)

	time.Sleep(time.Second)
	button.Clicked.Broadcast()

	clickSubscribe.Wait()
}