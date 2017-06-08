package main

import "fmt"

const kmToMiles float64 = 0.621371

func main() {
	var km float64
	fmt.Print("Enter kilometres run: ")
	fmt.Scan(&km)
	miles := km * kmToMiles
	fmt.Println(km, "kilometres is", miles, "miles.")
}
