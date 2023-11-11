package channels

import (
	"fmt"
	"testing"
)

func TestChannelSimple(t *testing.T) {
	dataStream := make(chan int, 2)
	var readStream <- chan int
	var writeStream chan <- int

	readStream = dataStream
	writeStream = dataStream

	go func() {
		defer close(writeStream)
		for i:= 0 ; i < 5 ; i++ {
			writeStream <- i
		}
	}()
	
	for i := range readStream {
		fmt.Println(fmt.Sprintf("Reading: %d", i))
	}
	i, ok := <- readStream
	fmt.Println(fmt.Sprintf("Reading: %d; %v", i, ok))
}
