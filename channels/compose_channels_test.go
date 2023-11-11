package channels

import (
	"fmt"
	"testing"
	"time"
)

type A struct {
	PropA string
}

type B struct {
	PropB string
}

func processA(a string) ([]A, error) {
	time.Sleep(time.Second * 2)
	return []A{
		A{
			PropA: a,
		},
	}, nil
}

func processB(a, b string) (map[string]B, error) {
	time.Sleep(time.Second * 4)
	return map[string]B{
		a: {
			PropB: b,
		},
	}, nil
}

func composeChannels() error {
	var a []A
	var b map[string]B
	errorA := make(chan error)
	errorB := make(chan error)

	go func() {
		defer close(errorA)
		var err error
		a, err = processA("A")
		if err != nil {
			errorA <- err
		}
	}()

	go func() {
		defer close(errorB)
		var err error
		b, err = processB("a", "b")
		if err != nil {
			errorB <- err
		}
	}()


	for i := 0; i < 2; i++ {
		select {
		case errA := <-errorA:
			if errA != nil {
				return errA
			}
		case errB := <-errorB:
			if errB != nil {
				return errB
			}
		}
	}

	fmt.Println(a)
	fmt.Println(b)
	return nil
}

func TestComposeChannels(t *testing.T) {
	err := composeChannels()
	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}
}
