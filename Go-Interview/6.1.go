package main

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name     string
	language string
}

func main() {
	// p := Programmer{"stefno", "go"}
	// fmt.Println(p)

	// name := (*string)(unsafe.Pointer(&p))
	// *name = "qcrao"

	// lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	// *lang = "Golang"

	// fmt.Println(p)

	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s))

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s))

}
