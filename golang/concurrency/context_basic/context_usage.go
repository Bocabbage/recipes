package context_basic

import (
	"context"
	"fmt"
	"time"
)

func mockWriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// Done channel close 时走这条路 (cancel)
			fmt.Println("Write Redis finish, stop running")
			return
		default:
			// [!] 实际场景中需要需要将可能会一直阻塞的部分抽出到select选择中，
			// 防止在此处阻塞导致无法执行到ctx.Done()，goroutine无法正常退出
			// example:
			// case [redis io channel]:
			// 	if [some-condition to check redis io result] {
			// 		select {
			// 		case <-ctx.Done():
			// 			return
			// 		case [some-case]:
			//		...
			// 		}
			// 	}
		}

		fmt.Println("Write Redis running")
		// mock: redis io
		// ...
		time.Sleep(2 * time.Second)
	}
}

func mockWriteMySQL(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// Done channel close 时走这条路 (cancel)
			fmt.Println("Write MySQL finish, stop running")
			return
		default:
			fmt.Println("Write MySQL running")
			// mock: mysql io
			// ...
			time.Sleep(2 * time.Second)
		}
	}
}

func handleRequest(ctx context.Context) {
	go mockWriteMySQL(ctx)
	go mockWriteRedis(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandleRequest Done.")
			return
		default:
			fmt.Println("HandleRequest running.")
			time.Sleep(2 * time.Second)
		}
	}
}

func CancelContextExample() {

	ctx, cancel := context.WithCancel(context.Background())
	go handleRequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines")
	cancel()

	time.Sleep(5 * time.Second)
}

func TimeoutContextExample() {

}

func ValueContextExample() {

}
