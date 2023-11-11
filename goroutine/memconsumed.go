package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//f, err := os.Create("myprogram.prof")
	//if err != nil {
	//
	//	fmt.Println(err)
	//	return
	//
	//}
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()


	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { x:= 1; fmt.Sprintf("%d", x); wg.Done(); <-c }
	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop() }
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
