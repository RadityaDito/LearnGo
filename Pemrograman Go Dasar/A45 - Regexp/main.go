package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)
	// []string{"banana", "burger"}

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
	// []string{"banana", "burger", "soup"}

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)

	var str2 = regex.FindString(text)
	fmt.Println(str2)
	// "banana"

	var idx2 = regex.FindStringIndex(text)
	fmt.Println(idx2)
	// [0, 6]

	var str = text[0:6]
	fmt.Println(str)
	// "banana"

	var str3 = regex.FindAllString(text, -1)
	fmt.Println(str3)
	// ["banana", "burger", "soup"]

	var str4 = regex.FindAllString(text, 1)
	fmt.Println(str4)
	// ["banana"]

	var str5 = regex.ReplaceAllString(text, "potato")
	fmt.Println(str5)

	// "potato potato potato"

	var str6 = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str6)
	// "banana potato soup"

	var str7 = regex.Split(text, -1)
	fmt.Printf("%#v \n", str7)
	// []string{"", "n", "n", " ", "urger soup"}
}
