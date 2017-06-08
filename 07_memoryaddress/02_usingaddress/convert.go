package main

import "fmt"

const milestokm float64 = 1.60934

func main() {
	var miles float64
	fmt.Print("Enter miles: ")
	fmt.Scan(&miles)
	km := miles * milestokm
	fmt.Println(miles, " miles in kilometres is ", km)

}
