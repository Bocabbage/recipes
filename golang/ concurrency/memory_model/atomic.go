package memorymodel

// (1) sync.WaitGroup: 信号量
// (2) atomic 库提供了涵盖基础类型到复杂类型的原子操作APIs
//		- atomic.Value (包含Load和Store方法)

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total struct {
	sync.Mutex
	value int
}

var totalValue uint64

func atomicTestMutexWorker(wg *sync.WaitGroup) {
	// 使用 Mutex 实现原子操作
	defer wg.Done() // 等价于: wg.Add(-1)

	for i := 0; i < 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func atomicTestAtomicWorker(wg *sync.WaitGroup) {
	// [Better] 使用 Atomic 实现原子操作
	defer wg.Done()

	var i uint64
	for i = 0; i < 100; i++ {
		atomic.AddUint64(&totalValue, i)
	}
}

func AtomicTestMutex() {
	// WaitGroup: 可以理解为一种 sem
	var wg sync.WaitGroup
	wg.Add(2)

	go atomicTestMutexWorker(&wg)
	go atomicTestMutexWorker(&wg)
	wg.Wait()
	fmt.Printf("[Mutex]total value: %d\n", total.value)

	wg.Add(2)

	go atomicTestAtomicWorker(&wg)
	go atomicTestAtomicWorker(&wg)
	wg.Wait()
	fmt.Printf("[Atomic]total value: %d\n", totalValue)
}
