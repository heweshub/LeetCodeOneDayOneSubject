package main

import (
	"fmt"
	"reflect"
	"time"
)

// func main() {
// 	var mailbox uint8

// 	var lock sync.RWMutex

// 	sendCond := sync.NewCond(&lock)
// 	recvCond := sync.NewCond(lock.RLocker())

// 	sign := make(chan struct{}, 3)
// 	max := 5
// 	go func(max int) {
// 		defer func() {
// 			sign <- struct{}{}
// 		}()

// 		for i := 1; i <= max; i++ {
// 			time.Sleep(time.Millisecond * 500)
// 			lock.Lock()
// 			for mailbox == 1 {
// 				sendCond.Wait()
// 			}
// 			log.Printf("sender [%d]: the mailbox is empty!", i)
// 			mailbox = 0
// 			log.Printf("sender [%d]: the letter has been sent.", i)
// 			lock.Unlock()
// 			recvCond.Signal()
// 		}
// 	}(max)
// 	go func(max int) {
// 		defer func() {
// 			sign <- struct{}{}
// 		}()
// 		for j := 1; j <= max; j++ {
// 			time.Sleep(time.Millisecond * 500)
// 			lock.RLock()
// 			for mailbox == 0 {
// 				recvCond.Wait()
// 			}
// 			log.Printf("receiver [%d]: the mailbox is empty!", j)
// 			mailbox = 1
// 			log.Printf("receiver [%d]: the letter has been sent.", j)
// 			lock.Unlock()
// 			sendCond.Signal()
// 		}
// 	}(max)
// 	<-sign
// 	<-sign
// }

func bianli(obj interface{}) {
	t := reflect.TypeOf(obj)
	// v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%s\n", t.Field(i).Tag)

	}
	// rand.Intn(1)
}

type SomeStruct struct {
	a int
	s string
}

func test(obj interface{}) {
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	for i := 0; i < typ.NumField(); i++ {
		// tag := typ.Field(i).Tag
		fieldValue := val.Field(i).Interface()
		fmt.Println(fieldValue)
	}
}

func main() {
	// var ss SomeStruct
	// test(ss)
	// ss.a = 1
	// // ss.s = "22"
	// test(ss)
	Limiter()
}

func Limiter() {
	var capacity = 100
	var tokenBucket = make(chan struct{}, capacity)

	var fillInterval = time.Minute
	fileToken := func() {
		thicker := time.NewTicker(fillInterval)
		for {
			select {
			case <-thicker.C:
				tokenBucket <- struct{}{}
				fmt.Println("in time or in capacity")
			default:
				fmt.Println("current number of token is", len(tokenBucket), time.Now())
			}
		}
	}
	go fileToken()
	time.Sleep(time.Hour)
}
