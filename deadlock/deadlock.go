package main

import (
	"fmt"
	"sync"
	"time"
)

type Number struct {
	val int
	mutex sync.Mutex
}

func (n *Number) Lock() {
	fmt.Printf("Locking %d\n", n.val)
	n.mutex.Lock()
}

func (n *Number) Unlock() {
	fmt.Printf("Releasing %d\n", n.val)
	n.mutex.Unlock()
}

func (n *Number) UpdateVal(val int) {
	fmt.Printf("Updating from %d to %d\n", n.val, val)
	n.val = val
}


func main() {

	a := &Number{
		val: 1,
		mutex: sync.Mutex{},
	}

	b := &Number{
		val: 2,
		mutex: sync.Mutex{},
	}

	wg := sync.WaitGroup{}

	add := func(num1, num2 *Number)  {
		fmt.Printf("Extecuting with %d, %d \n", num1.val, num2.val)
		num1.Lock()
		val1 := num1.val
		defer num1.Unlock()

		time.Sleep(1 * time.Second)

		num2.Lock()
		val2 := num2.val
		defer num2.Unlock()

		fmt.Printf("Addition %d", val1 + val2)

		defer func(){
			fmt.Printf("Exiting go routine")
			wg.Done()
		}()
	}

	wg.Add(2)

	go add(a, b)
	go add(b, a)

	wg.Wait()

}