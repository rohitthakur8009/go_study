package main

import (
"fmt"
"os"
"syscall"
	"time"
)

func lockFile(filename string, wait time.Duration) (*os.File, error) {
		file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}


	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		file.Close()
		return nil, err
	}
	fmt.Println("File locked. Other processes cannot write to it.")
	time.Sleep(wait)
	return file, nil

}

func main() {
	filename := "test.log"


	lockedFile, err := lockFile(filename, time.Minute * 2)
	if err != nil {
		fmt.Println("Could not lock the file:", err)
		return
	}

	defer func() {
		syscall.Flock(int(lockedFile.Fd()), syscall.LOCK_UN)
		lockedFile.Close()
	}()


}

