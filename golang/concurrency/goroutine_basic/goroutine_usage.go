package goroutine_basic

import (
	"fmt"
	"time"
)

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// basic-usage for Goroutine:
func SpinnerTest() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

// [todo] basic-usage for unbuff-Channel:

// definition: 定义 write channel
func counterWorker(out chan<- int) {
	defer close(out)
	for x := 0; x <= 20; x++ {
		out <- x
	}
}

// definition: 定义 read channel + write channel
func squarerWorker(in <-chan int, out chan<- int) {
	defer close(out)
	for x := range in {
		out <- x * x
	}
}

// basic-usage for Channel:
func PipelineTest() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		defer close(naturals)
		for x := 0; x <= 20; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		defer close(squares)
		// 不停执行 x <- squares，直到chan-closed
		for x := range naturals {
			squares <- x * x
		}
	}()

	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}

func PipelineTestV2() {
	naturals := make(chan int)
	squares := make(chan int)

	go counterWorker(naturals)
	go squarerWorker(naturals, squares)

	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}
