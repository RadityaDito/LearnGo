package main

import (
	"fmt"
	"os"
)

func main() {
	orderSomeFood("pizza")
	orderSomeFood("burger")

	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		func() {
			defer fmt.Println("halo 3")
		}()
	}

	fmt.Println("halo 2")

	defer fmt.Println("halo")
	os.Exit(1)
	fmt.Println("selamat datang")

}

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silakan tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda:", menu)
}
