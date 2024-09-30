package main

import "fmt"

type person struct {
	name string `tag1`
	age  int    `tag2`
}

type student struct {
	person
	age   int
	grade int
}

type Person struct {
	name string
	age  int
}

type People = Person

func main() {
	// var s1 student
	// s1.name = "john wick"
	// s1.grade = 2

	// fmt.Println("name  :", s1.name)
	// fmt.Println("grade :", s1.grade)

	// var s4 = student{}
	// s4.name = "wick"
	// s4.grade = 2

	// var s2 = student{"ethan", 2}

	// var s3 = student{name: "jason"}

	// fmt.Println("student 1 :", s1.name)
	// fmt.Println("student 2 :", s2.name)
	// fmt.Println("student 3 :", s3.name)

	// var s5 = student{name: "wick", grade: 2}

	// var s6 *student = &s5
	// fmt.Println("student 1, name :", s5.name)
	// fmt.Println("student 4, name :", s6.name)

	// s6.name = "ethan"
	// fmt.Println("student 1, name :", s5.name)
	// fmt.Println("student 4, name :", s6.name)

	//Embedded Struct
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.person.age = 22 // age of person
	s1.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("grade :", s1.grade)

	var p1 = person{name: "wick", age: 21}
	var s2 = student{person: p1, grade: 2}

	fmt.Println("name  :", s2.name)
	fmt.Println("age   :", s2.age)
	fmt.Println("grade :", s2.grade)

	// anonymous struct tanpa pengisian property
	var _ = struct {
		person
		grade int
	}{}

	// anonymous struct dengan pengisian property
	var _ = struct {
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	}

	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	var p3 = Person{"wick", 21}
	fmt.Println(p3)

	var p2 = People{"wick", 21}
	fmt.Println(p2)
}
