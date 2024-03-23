package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

const Pi = 3.14

func main() {
	var x int = 10
	displayNumber(x)

	var integer int = 10
	var decimal float64 = 10.1
	var text string = "hello"
	var condition bool = true

	fmt.Println("nilai integer adalah:", integer)
	fmt.Println("nilai decimal adalah:", decimal)
	fmt.Println("nilai text adalah:", text)
	fmt.Println("nilai condition adalah:", condition)

	// struct
	john := Person{Name: "John Doe", Age: 25}
	fmt.Println("Struct Person:", john)

	// Slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice numbers:", numbers)

	// Map
	colors := map[string]string{"red": "#ff0000", "green": "#00ff00", "blue": "#0000ff"}
	fmt.Println("Map colors:", colors)

	// If Else
	num := 20

	if(num % 2 == 0) {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}

	// Looping
	for i, count :=  2, 0; count < 2; i++ {
		if isPrime(i){
			fmt.Println(i)
			count++
		}
	}

	// switch case
	day := 10

	switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		default:
			fmt.Println("Invalid day")
	}

	//pointer
	xy := 10
	fmt.Println("nilai xy adalah:", xy)

	p := &xy // pointer to xy
	*p = 20

	fmt.Println("nilai xy setelah diubah adalah:", xy)

	// aritmatika
	a := 10
	b := 20

	fmt.Println("a + b =", a + b)
	fmt.Println("a - b =", a - b)
	fmt.Println("a * b =", a * b)
	fmt.Println("a / b =", a / b)
	fmt.Println("a % b =", a % b)

	// operator logika
	c := true
	d := false
	fmt.Println("c && d =", c && d)
	fmt.Println("c || d =", c || d)
	fmt.Println("!c =", !c)

	// operator perbandingan
	fmt.Println("a == b", a == b)
	fmt.Println("a != b", a != b)
	fmt.Println("a > b", a > b)
	fmt.Println("a < b", a < b)
	fmt.Println("a >= b", a >= b)
	fmt.Println("a <= b", a <= b)
}

func displayNumber (number int) {
	fmt.Println("nilai number adalah:", number)
}

func isPrime (number int) bool {
	for i := 2; i < number; i++ {
		if number % i == 0 {
			return false
		}
	}

	return true
}