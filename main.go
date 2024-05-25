package main

import "fmt"

type Person struct {
		Name string
		Age int
}


func main () {
		// var a int = 10

		// displayNumber(a)

		// var integer int = 20
		// var float float64 = 30.5
		// var text string = "Hello World"
		// var condition bool = true

		// fmt.Println("nilai integer adalah:", integer)
		// fmt.Println("nilai float adalah:", float)
		// fmt.Println("nilai text adalah:", text)
		// fmt.Println("nilai condition adalah:", condition)

		// // struct
		// john := Person{Name: "John", Age: 25}
		// fmt.Println("Name:", john.Name)
		// fmt.Println("Age:", john.Age)

		// // slice
		// var fruits = []string{"apple", "grape", "banana"}
		// fmt.Println(fruits)
		// fmt.Println(fruits[0])

		// // map
		// var person = map[string]string{"name": "John", "age": "25"}
		// fmt.Println(person)

		// if else
		// number := 5

		// if number%2 == 0 {
		// 	fmt.Println("Genap")
		// } else {
		// 	fmt.Println("Ganjil")
		// }

		// // for loop
		// fmt.Println("Bilangan pertama dari:", number);
		// for i, count := 2, 0; count < number; i++ {
		// 	if isPrime(i) {
		// 		fmt.Println(i)
		// 		count++
		// 	}
		// }

		// switch case
		// day := 7

		// switch day {
		// 	case 1:
		// 		fmt.Println("Senin")
		// 	case 2:
		// 		fmt.Println("Selasa")
		// 	case 3:
		// 		fmt.Println("Rabu")
		// 	case 4:
		// 		fmt.Println("Kamis")
		// 	case 5:
		// 		fmt.Println("Jumat")
		// 	case 6:
		// 		fmt.Println("Sabtu")
		// 	case 7:
		// 		fmt.Println("Minggu")
		// 	default:
		// 		fmt.Println("Tidak ada hari")
		// 	}

		// x := 20
		// fmt.Println("nilai awal x adalah:", x)

		// p := &x
		// *p = 10

		// fmt.Println("nilai x setelah diubah adalah:", x)

		// operator
		a := 10
		b := 3

		fmt.Println("a + b =", a+b)
		fmt.Println("a - b =", a-b)
		fmt.Println("a * b =", a*b)
		fmt.Println("a / b =", a/b)

		// Operator lain
		c := true
		d := false

		fmt.Println("c && d =", c && d)
		fmt.Println("c || d =", c || d)
		fmt.Println("!c =", !c)

		// Operator perbandingan
		fmt.Println("a == b =", a == b)
		fmt.Println("a != b =", a != b)
		fmt.Println("a > b =", a > b)
		fmt.Println("a < b =", a < b)
		fmt.Println("a >= b =", a >= b)
		fmt.Println("a <= b =", a <= b)
}

func displayNumber(number int){
	fmt.Println("nilai number adalah:", number)
}

func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}