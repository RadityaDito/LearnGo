package main

import "fmt"

func main() {

	//Perulangan Menggunakan Keyword for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	//Perulangan Menggunakan Keyword for dengan argumen hanya kondisi
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	var i2 = 0

	//Perulangan Menggunakan Keyword for tanpa argumen
	for {
		fmt.Println("Angka", i2)

		i2++
		if i2 == 5 {
			break
		}
	}

	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Print(i) // 01234
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// Label
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
