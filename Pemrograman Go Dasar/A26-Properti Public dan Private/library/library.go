package library

import "fmt"

func SayHello(name string) {
	fmt.Println("hello")
	introduce(name)
}

func introduce(name string) {
	fmt.Println("nama saya", name)
}

type Student struct {
	Name  string
	Grade int
}

// var Student = struct {
// 	Name  string
// 	Grade int
// }{}

// func init() {
// 	Student.Name = "John Wick"
// 	Student.Grade = 2

// 	fmt.Println("--> library/library.go imported")
// }
