package main

import "fmt"

type student struct {
	name  string
	grade int
}

func basic() {
	var s1 student
	s1.name = "Rizqi"
	s1.grade = 22
	var s2 = student{"ethan", 2}
	var s3 = student{name: "jason"}
	var s4 = student{name: "jason", grade: 2}
	var s5 = student{grade: 2, name: "bruce"}

	// variabel objek pointer
	var s6 *student = &s1

	fmt.Printf("student s1 Nama : %s, grade : %x \n", s1.name, s1.grade)
	fmt.Printf("student s2 Nama : %s, grade : %x \n", s2.name, s2.grade)
	fmt.Printf("student s3 Nama : %s, grade : %x \n", s3.name, s3.grade)
	fmt.Printf("student s4 Nama : %s, grade : %x \n", s4.name, s4.grade)
	fmt.Printf("student s5 Nama : %s, grade : %x \n", s5.name, s5.grade)
	fmt.Printf("student s6 Nama : %s, grade : %x \n", s6.name, s6.grade)
}
