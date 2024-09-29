package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	var data map[string]int
	// data["one"] = 1
	// akan muncul error!

	data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	var newChicken = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(newChicken)) // 2
	fmt.Println(newChicken)

	delete(newChicken, "januari")

	fmt.Println(len(newChicken)) // 1
	fmt.Println(newChicken)

	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
