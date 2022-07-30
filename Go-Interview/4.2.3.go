package main

import (
	"fmt"
)

// func main() {
// 	taskCh := make(chan int, 100)
// 	go worker(taskCh)
// 	for i := 0; i < 10; i++ {
// 		taskCh <- i
// 	}
// 	// time.Sleep(time.Second * 10)
// }

func worker(taskCh <-chan int) {
	const N = 5

	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("task %d by worker %d\n", task, id)
			}
		}(i)
	}
}
