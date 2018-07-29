package main

import "fmt"

var fullName = firstName + " " + secondName
var firstName = "John"
var secondName = "Smith"

func init() {
	fmt.Println("Calling init()")
}

func main() {
	fmt.Println(fullName)
}
