package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	lock := sync.RWMutex{}

	b := 0

	for i := 11; i < 20; i++ {
		go func(i int) {
			lock.RLock()
			fmt.Printf("RLock: from go routine %d: b = %d\n", i, b)
			time.Sleep(5 * time.Second)
			lock.RUnlock()
		}(i)
	}

	a := 0
	for b := 1; b < 10; b++ {
		go func(b int) {
			lock.Lock()
			fmt.Printf("Lock: from go routine %d: a = %d\n", b, a)
			time.Sleep(3 * time.Second)
			lock.Unlock()
		}(b)
	}

	<-time.After(time.Second * 60)
}

//Lock(): only one go routine read/write at a time by acquiring the lock
//RLock(): multiple go routine can read(not write) at a time by acquiring the lock
