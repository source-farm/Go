package main

import (
	"fish"
	"fmt"
)

func main() {
	if fish.Kind("Salmon") == fish.Sea {
		fmt.Println("Salmon is sea fish.")
	}
}
