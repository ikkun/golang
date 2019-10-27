package main

import "fmt"

func foo() {
	defer fmt.Println("Done!")
	defer fmt.Println("Are we done?")
	fmt.Println("Doing some stuff, who knows what?")

}

func main() {
	foo()
}
