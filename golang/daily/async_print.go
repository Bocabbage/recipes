package main

import "fmt"

// 有三个函数，分别打印"cat", "fish","dog"要求每一个函数都用一个goroutine，按照顺序打印100次。

func printAnimal(in <-chan struct{}, out chan<- struct{}, s string) {
	for range in {
		fmt.Println(s)
		out <- struct{}{}
	}
	close(out)

}

func AsyncPrint() {
	mainChan := make(chan struct{})
	catChan := make(chan struct{})
	fishChan := make(chan struct{})
	dogChan := make(chan struct{})

	go printAnimal(catChan, fishChan, "Cat")
	go printAnimal(fishChan, dogChan, "Fish")
	go printAnimal(dogChan, mainChan, "Dog")

	go func(in chan<- struct{}) {
		// Start signal
		in <- struct{}{}
	}(mainChan)

	for i := 0; i < 10; i++ {
		<-mainChan
		fmt.Printf("Current: %d\n", i+1)
		catChan <- struct{}{}
	}
	close(catChan)
}
