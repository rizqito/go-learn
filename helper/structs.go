package helper

import "fmt"

//struct harus buat bagian seperti dibawah ini
type Person struct { // membuat struct person
	name    string // deklarasi variabel dengan tipe data
	age     int
	address string
}

type Employee struct { // membuat struct employee
	name     string
	address  string
	position string
}

func BelajarStruct() {
	fmt.Println("----------struct var array---------")
	var arrPerson []Person                // deklarasi variabel array pada structs
	arrPerson = append(arrPerson, Person{ // untuk mengisi data pada array struct
		name:    "Qito",
		age:     22,
		address: "Samarinda",
	})
	fmt.Println(arrPerson)

	fmt.Println("----------struct var biasa---------")
	var person Person //deklarasi variabel pada structs
	person.name = "Rizqi"
	person.age = 22
	person.address = "Samarinda"
	fmt.Println(person)

	fmt.Println("----------latihan---------")
	// buat structs employee data
	// employee struct isi atributnya ada name, address, position
	// buat array dari structs employee
	// isi array of struct nya
	// print array struct employee

	employees := []Employee{}               // deklarasi variabel array pada structs dengan short hand
	employees = append(employees, Employee{ // untuk mengisi data pada array struct
		name:     "Qito",
		address:  "Jl. Giri Makmur, RT.22, No.19",
		position: "Samarinda",
	})
	fmt.Println(employees)

	for _, value := range employees { // untuk menampilkan single data dengan foreach
		fmt.Println("Nama saya : ", value.name)
		fmt.Println("Alamat : ", value.address)
		fmt.Println("Posisi di : ", value.position)
	}
}
