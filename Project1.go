package main

import "fmt"

//SCOPE
//GLOBAL VARAIBLE
var gVariable int = 500

/*
	การประกาศตัวแปร
*/
func main() {
	var myAge int = 37
	myAge2 := 38.9
	var something bool = true
	text := "Kreangkrai"
	age1, age2 := 35, 80
	fNumber := 60
	fSecond := 5
	fmt.Println("Value variable = ", myAge)
	fmt.Println("Value Variable = ", myAge2)
	fmt.Print("Value =", something)
	fmt.Print(text)
	fmt.Println(age1, age2)
	fmt.Println(fNumber + fSecond)

	//ชุดข้อความ
	p1 := "Kreangkrai"
	p2 := "Security Analyst"
	//Concatenation
	p3 := p1 + p2
	fmt.Println(p3)
	fmt.Println(p3[1:3])
	fmt.Println(p3[0])
	fmt.Println(p3[1:])

	//Boolean
	isEmpty := true
	isJumping := false
	fmt.Println(isEmpty)
	fmt.Println(isJumping)
	someBoolean := 5 == 3
	fmt.Println(someBoolean)

	lVariable := 40
	fmt.Println("Global Val=", gVariable)
	fmt.Println("Local Variable=", lVariable)
	anotherFunction()

	//Input Scanf
	fmt.Print("Input Your Number:")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println("Output = ", output)

	//Condition Statement
	condition := input > 2
	if condition {
		fmt.Println("Worked")
	} else {
		fmt.Println("Not worked.")
	}

	//&& เงื่อนไข
	//||
	if 6 > 3 && 8 > 5 {
		fmt.Println("Worked")
	}

	//switch
	fmt.Print("Enter Your Number:")
	var number int
	fmt.Scanf("%d", &number)
	switch number {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Unknown")

	}

	//for loop
	for i := 1; i <= 10; i++ {
		fmt.Println("Kreangkrai")
		if i%2 == 0 {
			fmt.Println(i, "--")
		} else {
			fmt.Println(i, "-")
		}
	}
	i := 1
	for i < 15 {
		fmt.Println("Kreangkrai")
		i = i + 1
	}

	dosomething("Golf Kreangkrai")
	addition(8, 2)
	result := addition2(5, 5)
	fmt.Println(result * 10)

	summation(1, 2, 3, 5)

	fmt.Println(factorial(5))

	//array
	x := [5]float64{6, 67, 74, 4, 6}
	var total float64
	for _, value := range x {
		total += value
	}
	fmt.Println(total)

	//slice
	y := make([]float64, 5)
	fmt.Println(y)
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 9, 10)
	fmt.Println(slice1)
	fmt.Println(slice2)

	arr := [5]float64{1, 2, 3, 4, 5}
	z := arr[0:4]

	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice3)

}

func anotherFunction() {
	lVariable := 50
	fmt.Println("Global Val=", gVariable)
	fmt.Println("Local Variable=", lVariable)

}

func dosomething(str string) {
	fmt.Println(str)
}

func addition(a int, b int) {
	fmt.Println(a + b)
}

func addition2(a int, b int) int {
	output := a + b
	return output
}

func summation(args ...int) {
	var total int
	for _, n := range args {
		total += n
	}
	fmt.Println(total)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
