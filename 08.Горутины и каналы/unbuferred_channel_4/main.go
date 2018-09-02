package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for n := 1; n <= 3; n++ {
		out <- n
		time.Sleep(1 * time.Second)
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}

func printer(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}
