package main

import "fmt"

func handleInput(input int) {
	if input == 42 {
		fmt.Println("Hello Universe!")
	} else {
		fmt.Println(input)
	}
}

func main()  {
	var input int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	handleInput(input)
}