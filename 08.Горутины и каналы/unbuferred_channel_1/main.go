package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	fmt.Println("waiting for goroutine to complete")
	// В этом месте происходит блокировка, пока горутина запущенная чуть
	// выше не отправит значение в канал done. Блокировка происходит,
	// т.к. канал done является небуферизированным.
	<-done
	fmt.Println("gouroutine completed")
}
