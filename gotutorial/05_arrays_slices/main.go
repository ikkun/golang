package main

import "fmt"

func main() {
	//Arrays
	// val:=[5]{1,2,3,4,5}
	// var fruitArra [2]string

	// Assign values
	// fruitArra[0]="Apple"
	// fruitArra[1]="Orange"
	fruitArra := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitArra)
	fmt.Println(fruitArra[1])
	fmt.Println(len(fruitArra))
	fmt.Println(fruitArra[1:3])
}
