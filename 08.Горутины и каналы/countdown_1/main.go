package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 3; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	// launch
}
