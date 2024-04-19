package main

import "fmt"

func printNumString() {
	numTurn := make(chan struct{})
	alphaTurn := make(chan struct{})
	join := make(chan struct{})

	go func() {
		for i := 1; i <= 26; i++ {
			fmt.Printf("%d", i)
			if i%2 == 0 {
				alphaTurn <- struct{}{}
				if i != 26 {
					<-numTurn
				}
			}
		}
		join <- struct{}{}
	}()

	go func() {
		<-alphaTurn
		for j := 1; j <= 26; j++ {
			fmt.Printf("%c", j+96)

			if j%2 == 0 && j != 26 {
				numTurn <- struct{}{}
				<-alphaTurn
			}
		}
		join <- struct{}{}
	}()

	for i := 0; i < 2; i++ {
		<-join
	}
}

func main() {
	printNumString()
}
