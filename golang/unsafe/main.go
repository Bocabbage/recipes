package main

import (
	"fmt"
	"unsafe"
)

func sizeFunc() {
	fmt.Println(unsafe.Sizeof(float64(0))) // "8"

	var x struct {
		a bool
		b int16
		c []int
	}
	fmt.Printf("sizeof x: %v\n", unsafe.Sizeof(x))
}

func pointerFunc() {
	fmt.Printf("unsafe pointer: %#16x\n", func(f float64) uint64 {
		return *(*uint64)(unsafe.Pointer(&f))
	}(1.0))
}

func main() {
	sizeFunc()
	pointerFunc()
}
