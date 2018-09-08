package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	cancel := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		cancel <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)

	fmt.Println("Commencing countdown.")
	for countdown := 3; countdown >= 0; countdown-- {
		select {
		case <-tick:
			fmt.Println(countdown)
		case <-cancel:
			fmt.Println("Launch canceled!")
			return
		}
	}

	fmt.Println("Launch rocket!")
}
