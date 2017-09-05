package main

import "fmt"

func main() {
	var number1 int
	var number2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&number1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&number2)

	if number1 > number2 {
		fmt.Println("The remainder of", number1, "divided by", number2, "is", number1%number2)
	} else {
		fmt.Println("The remainder of", number2, "divided by", number1, "is", number2%number1)
	}
}
