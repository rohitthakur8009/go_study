package channels

import (
	"fmt"
	"math/rand"
	"testing"
)

func generator(done <-chan interface{}, ints []int) <-chan int {
	inputStream := make(chan int)
	go func() {
		for _, i := range ints {
			select {
			case <-done:
				return
			case inputStream <- i:
			}
		}
	}()
	return inputStream
}

func multiplierStage(done <-chan interface{}, inputStream <-chan int, multiplier int) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range inputStream {
			select {
			case multipliedStream <- i * multiplier:
			case <-done:
				return
			}
		}
	}()
	return multipliedStream
}

func adderStage(done <-chan interface{}, inputStream <-chan int, adder int) <-chan int {
	adderStream := make(chan int)
	go func() {
		defer close(adderStream)
		for i := range inputStream {
			select {
			case adderStream <- i + adder:
			case <-done:
				return
			}
		}
	}()
	return adderStream
}

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		for {
			for _, v := range values {
				select {
				case valueStream <- v:
				case <-done:
					return

				}
			}
		}
	}()
	return valueStream
}

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		for {
				select {
				case valueStream <- fn():
				case <-done:
					return

			}
		}
	}()
	return valueStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
				case takeStream <- <- valueStream:
			}
		}
	}()
	return takeStream
}

func toInt(done <-chan interface{}, valueStream <- chan interface{}) <- chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream{
			select {
			case intStream <- v.(int):
			case <- done:
				return
			}
		}
	}()
	return intStream
}

func TestPipeline(t *testing.T) {

	done := make(chan interface{})

	defer close(done)

	ints := []int{1, 2, 3, 4, 5, 6, 7}

	intStream := generator(done, ints)

	pipeLine := multiplierStage(done, adderStage(done, intStream, 2), 10)

	for v := range pipeLine {
		fmt.Println(v)
	}
}

func TestPipelineTake(t *testing.T) {

	done := make(chan interface{})

	defer close(done)


	pipeLine := take(done, repeat(done, 2, 3, 4), 10)

	for v := range pipeLine {
		fmt.Println(v)
	}
}

func TestPipelineRandomGen(t *testing.T) {

	done := make(chan interface{})

	defer close(done)

	randFunc := func() interface{} {
		return rand.Int()
	}

	pipeLine := take(done, repeatFn(done, randFunc), 10)

	for v := range pipeLine {
		fmt.Println(v)
	}
}


func TestPipelineRandomInt(t *testing.T) {

	done := make(chan interface{})

	defer close(done)

	randFunc := func() interface{} {
		return rand.Int()
	}

	pipeLine := toInt(done, take(done, repeatFn(done, randFunc), 10))

	for v := range pipeLine {
		fmt.Println(v)
	}
}