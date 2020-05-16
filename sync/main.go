package main

import (
	"fmt"
	"sync"
)

//锁

var x = 0
var lock sync.Mutex
var wg sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
