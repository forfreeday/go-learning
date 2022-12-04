package main

import (
	"fmt"
	"sync"
)

var (
	testLock = sync.Mutex{}
	count    int
)

func getLock() int {
	testLock.Lock()
	defer testLock.Unlock()
	fmt.Println("getLock: ", count)
	return count
}

func setLock(c int) {
	testLock.Lock()
	defer testLock.Unlock()
	fmt.Println("setLock: ", c)
	count = c
}

func main() {

	for i := 0; i < 100; i++ {
		go setLock(i)
	}
	for i := 0; i < 100; i++ {
		go getLock()
	}
}
