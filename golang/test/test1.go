package main

import "fmt"

var (
	firstName, lastName string
)

func main() {
}
func getName() {
	fmt.Println("Please enter your name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
}
