package main

import "fmt"

func main() {
	//Similar with array, but size is dynamic
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

	var fruitsA = []string{"apple", "grape"}     // slice
	var fruitsB = [2]string{"banana", "melon"}   // array
	var fruitsC = [...]string{"papaya", "grape"} // array

	fmt.Println(fruitsA)
	fmt.Println(fruitsB)
	fmt.Println(fruitsC)

	var newFruits = fruits[0:2]

	fmt.Println(newFruits) // ["apple", "grape"]

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var fruits2 = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits2, "papaya")

	fmt.Println(fruits2) // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	var fruits3 = []string{"apple", "grape", "banana"}
	var aFruits3 = fruits[0:2]
	var bFruits3 = fruits[0:2:2]

	fmt.Println(fruits3)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits3)) // len: 3
	fmt.Println(cap(fruits3)) // cap: 3

	fmt.Println(aFruits3)      // ["apple", "grape"]
	fmt.Println(len(aFruits3)) // len: 2
	fmt.Println(cap(aFruits3)) // cap: 3

	fmt.Println(bFruits3)      // ["apple", "grape"]
	fmt.Println(len(bFruits3)) // len: 2
	fmt.Println(cap(bFruits3)) // cap: 2
}
