package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

func main() {

	// var rwlock sync.RWMutex
	// fmt.Println(rwlock)
	var a atomic.Value
	fmt.Println(a.Load())
	atomic.AddUint32(&a, ^uint32(3-1))
	runtime.Gosched()
	runtime.GC()
	sync.Once
	sync.WaitGroup
	var m sync.Map
	m.Load()
	// fmt.Println(mutexLocked)
	// fmt.Println(mutexWoken)
	// fmt.Println(mutexStarving)
	// fmt.Println(mutexWaiterShift)

	// a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(a)

	// s1 := a[:]
	// s2 := s1[:]
	// fmt.Println(s1)
	// fmt.Println(s2)

	// fmt.Println("")
	// fmt.Println("-------------------------------")
	// fmt.Println("")

	// a[9] = 10
	// s2[0] = 100
	// fmt.Println(a)
	// fmt.Println(s1)
	// fmt.Println(s2)

}
