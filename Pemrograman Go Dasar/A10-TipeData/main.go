package main

import "fmt"

func main() {

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var message2 = `Nama saya "John Wick".
Salam kenal.
Mari belajar "Golang".`

	fmt.Println(message2)

	// Demonstrating various data types in Go
	// Integer
	var integerNumber int = 10
	fmt.Println("Integer:", integerNumber)

	// Float
	var floatNumber float64 = 3.14
	fmt.Println("Float:", floatNumber)

	// String
	var text string = "Hello, Go!"
	fmt.Println("String:", text)

	// Boolean
	var isTrue bool = true
	fmt.Println("Boolean:", isTrue)

	// Array
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)

	// Slice
	slice := []int{1, 2, 3}
	fmt.Println("Slice:", slice)

	// Map
	person := map[string]string{
		"name": "John",
		"age":  "30",
	}
	fmt.Println("Map:", person)

	// Pointer
	x := 10
	pointerToX := &x
	fmt.Println("Pointer:", pointerToX, "Value:", *pointerToX)
}
