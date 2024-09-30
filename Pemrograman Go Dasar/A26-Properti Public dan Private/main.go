package main

import (
	"belajar-golang-level-akses/library"
	f "fmt"
)

// ."belajar-golang-level-akses/library"

func main() {
	library.SayHello("ethan")

	var s1 = library.Student{Name: "ethan", Grade: 21}
	f.Println("name ", s1.Name)
	f.Println("grade", s1.Grade)
}
