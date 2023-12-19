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
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer
	for {
		fmt.Println(<-squares)
	}
}
