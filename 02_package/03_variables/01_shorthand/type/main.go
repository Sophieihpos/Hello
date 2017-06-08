package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 3.14
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \t %T \t %T \n", e, f, g)

}
