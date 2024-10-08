package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

func calculate(d float64) (area float64, circumference float64) {
	// hitung luas
	area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}
