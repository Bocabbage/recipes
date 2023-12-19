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

// ----- Basic-usage for unbuff-Channel -----

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

// ---- Wait for finish
func WaitRoutineTest() {
	// 用于同步的无语义channel:
	ch := make(chan struct{})

	for i := 0; i < 5; i++ {
		go func(ivar int) {
			// 闭包获取ch
			time.Sleep(1 * time.Second)
			fmt.Printf("%d\n", ivar)
			ch <- struct{}{}
		}(i)
	}

	// Join
	for i := 0; i < 5; i++ {
		<-ch
	}

	fmt.Println("Finish!")
}

// 带 error-handle 的 Join
type item struct {
	param int
	err   error
}

func stringWorker(s string) (int, error) {
	var errorResult error
	fmt.Println(s)
	return 0, errorResult
}

func WaitRoutineWithErrorTest() ([]int, error) {
	// 在 channel 中带入 error 信息，并使用 buff-channel（否则一个error routine就可能导致goroutinue leak）
	ch := make(chan item, 5)

	for i := 0; i < 5; i++ {
		go func() {
			var it item
			it.param, it.err = stringWorker("Hey!")
			ch <- it
		}()
	}

	// Join
	result := make([]int, 0)
	for i := 0; i < 5; i++ {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		result = append(result, it.param)
	}

	fmt.Println(result)
	return result, nil
}
