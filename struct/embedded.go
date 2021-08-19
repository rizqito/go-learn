package main

import "fmt"

type person struct {
	name string
	age  int
}

type murid struct {
	grade  int
	person //* embedded structs
}

func embedded() {
	var s1 = murid{}
	s1.name = "john"
	s1.age = 21
	s1.grade = 2
	fmt.Printf("student s1 Nama : %s, age : %x, age : %x, grade : %x\n", s1.name, s1.age, s1.person.age, s1.grade)

	// --------------
	var p1 = person{name: "john", age: 22}
	var m1 = murid{person: p1, grade: 2}
	fmt.Printf("student s1 Nama : %s, age : %x, age : %x, grade : %x\n", m1.name, m1.age, m1.person.age, m1.grade)

	//? anonymous struct
	var n1 = struct {
		person
		grade int
	}{person: p1, grade: 2}
	fmt.Printf("student n1 Nama : %s, age : %x, age : %x, grade : %x\n", n1.name, n1.age, n1.person.age, n1.grade)

	//? slice & struct combination
	var allMurid = []person{
		{name: "jonathan", age: 23},
		{name: "glenn", age: 26},
		{name: "beth", age: 22},
	}
	for _, siswa := range allMurid {
		fmt.Println(siswa.name, " berumur ", siswa.age)
	}

	//? slice anonymous struct
	var allStudent = []struct {
		person
		grade int
	}{
		{person: person{"wick", 30}, grade: 2},
		{person: person{"magie", 25}, grade: 3},
		{person: person{"rick", 35}, grade: 1},
	}
	for _, student := range allStudent {
		fmt.Println(student.person.name, "berumur", student.person.age, " mendapat grade", student.grade)
	}

	//? nested struct
	type nestedStruct struct {
		orang struct {
			name string
			umur int
		}
		grade int
		hobi  []string
	}

	//? alias
	type people = person
	var q1 = person{"wick", 21}
	fmt.Println(q1)
}
