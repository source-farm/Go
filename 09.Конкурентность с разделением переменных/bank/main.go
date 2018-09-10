package main

import "fmt"

var balance int

func Deposit(amount int) { balance = balance + amount }
func Balance() int       { return balance }

func main() {
	// Alice:
	go func() {
		Deposit(200)
		fmt.Println("= ", balance)
	}()

	// Bob:
	go Deposit(100)
}
