package main

import "fmt"

func main() {
	// Константа PI экспортируется из файла "tools.go", который принадлежит тому же
	// пакету, что и этот файл. Поэтому мы и можем обращаться к этой константе.
	fmt.Printf("π = %f\n", PI)

	// Вызов circleArea проходит, т.к. это объявление уровня пакета, т.е. эта функция
	// видна только файлам пакета "main".
	r := 2.0
	fmt.Printf("Area of circle with radius %.2f: %.2f", r, circleArea(r))
}
