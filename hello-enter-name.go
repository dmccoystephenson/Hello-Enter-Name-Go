package main

import "fmt"

func main() {
	fmt.Println("Enter your first name.")
	var firstName string
	fmt.Scanln(&firstName)
	fmt.Println("Hello " + firstName + "!")
}