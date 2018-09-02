package main

import (
	"fmt"
	"time"
)

func main() {
	// Конвейер из трёх горутин.

	naturals := make(chan int)
	squares := make(chan int)

	// Отправляем в канал naturals натуральные числа от 1 до 3 включительно с задержкой в 1 сек.
	go func() {
		for n := 1; n <= 3; n++ {
			naturals <- n
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	// Извлекаем значения из канала naturals, находим их квадрат и отправляем дальше в канал squares.
	go func() {
		// Цикл завершится, когда канал naturals будет закрыт и полностью высушен.
		for n := range naturals {
			squares <- n * n
		}
		close(squares)
	}()

	// Выводим значения из канала squares(основная горутина).
	for n := range squares {
		fmt.Println(n)
	}
}
