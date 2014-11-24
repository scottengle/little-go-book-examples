package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	lock    sync.Mutex
)

type synchronizer struct{}

func (s *synchronizer) Incr() {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}

func (s *synchronizer) Increment() {
	for i := 0; i < 2; i++ {
		go s.Incr()
	}
	time.Sleep(time.Millisecond * 10)
}
