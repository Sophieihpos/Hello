package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	half := func(n int) (float64, bool) {
		return float64(n) / 2, n%2 == 0
	}
	h, even := half(number)
	fmt.Println(h, even)
}
