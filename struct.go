package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//* struct function/method
func (customer Customer) sayHi(name string) {
	fmt.Println("hello", name, "namaku ", customer.Name)
}

func structBasic1() {
	qito := Customer{
		Name:    "qito",
		Address: "New York",
		Age:     30,
	}
	fmt.Println(qito)

	qito.sayHi("Joko")
}

type student struct {
	name  string
	grade int
}

func structBasic2() {
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

type person struct {
	name string
	age  int
}

type murid struct {
	grade  int
	person //* embedded structs
}

func structEmbedded() {
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
