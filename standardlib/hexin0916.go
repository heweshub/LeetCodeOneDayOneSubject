package main

import (
	"fmt"
	"testing"
)

// func main() {
// 	// var count uint32

// 	// trigger := func(i uint32, fn func()) {
// 	// 	for {
// 	// 		if n := atomic.LoadUint32(&count); n == i {
// 	// 			fn()
// 	// 			atomic.AddUint32(&count, 1)
// 	// 			break
// 	// 		}
// 	// 		time.Sleep(time.Nanosecond)
// 	// 	}
// 	// }
// 	// for i := uint32(0); i < 10; i++ {
// 	// 	go func(i uint32) {
// 	// 		fn := func() {
// 	// 			fmt.Println(i)
// 	// 		}
// 	// 		trigger(i, fn)
// 	// 	}(i)
// 	// }
// 	// trigger(10, func() {})
// 	// a := make([]byte, 100000)
// 	// fmt.Println(a)

// 	// numbers1 := []int{1, 2, 3, 4, 5, 6}
// 	// for i := range numbers1 {
// 	// 	if i == 3 {
// 	// 		numbers1[i] |= i // 4 和 3进行按位或操作得到的结果是 7
// 	// 	}
// 	// }
// 	// fmt.Println(numbers1)
// 	// fmt.Println()

// 	// numbers3 := []int{1, 2, 3, 4, 5, 6} //切片 引用类型 number3[i]的结果值一直在变
// 	// maxIndex3 := len(numbers3) - 1
// 	// for i, e := range numbers3 {
// 	// 	if i == maxIndex3 {
// 	// 		numbers3[0] += e
// 	// 	} else {
// 	// 		numbers3[i+1] += e
// 	// 	}
// 	// 	fmt.Println(numbers3)
// 	// }
// 	// syscall.Errno
// 	// fmt.Println(numbers3)

// 	f()
// 	fmt.Println("next")
// }

func f() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("solve panic")
		}
	}()
	panic("panic")
}

func Testf(t *testing.B) {
	// t.N
}
