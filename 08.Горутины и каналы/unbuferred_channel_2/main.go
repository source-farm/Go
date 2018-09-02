package main

import (
	"fmt"
	"time"
)

func main() {
	// Конвейер из трёх горутин.

	naturals := make(chan int)
	squares := make(chan int)

	// Отправляем в канал naturals натуральные числа с задержкой в 1 сек.
	go func() {
		n := 1
		for {
			naturals <- n
			n++
			time.Sleep(1 * time.Second)
		}
	}()

	// Извлекаем значения из канала naturals, находим их квадрат и отправляем дальше в канал squares.
	go func() {
		for {
			n := <-naturals
			squares <- n * n
		}
	}()

	// Выводим значения из канала squares(основная горутина).
	for {
		fmt.Println(<-squares)
	}
}
