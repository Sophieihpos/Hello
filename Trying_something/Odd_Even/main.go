package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a whole number: ")
	fmt.Scan(&number)
	outcome := number % 2

	if outcome == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	fmt.Print("Yeah motherfucker!")
}
