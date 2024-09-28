package main

import "fmt"

const firstName string = "john"

const (
	square         = "kotak"
	isToday  bool  = true
	numeric  uint8 = 1
	floatNum       = 2.2
)

const (
	a = "konstanta"
	b
)

const (
	today string = "senin"
	sekarang
	isToday2 = true
)

const satu, dua = 1, 2
const three, four string = "tiga", "empat"

func main() {
	fmt.Print("halo ", firstName, "!\n")

	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")
}
