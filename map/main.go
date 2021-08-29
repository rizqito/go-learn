package main

import "fmt"

func main() {
	//              Key   value
	//              v       v
	person := map[string]string{
		"name":    "Rizqi",
		"address": "Melbourne",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//* membuat map baru
	book := make(map[string]string)
	book["title"] = "buku novel"
	book["author"] = "rizqi"
	book["wrong"] = "ups"

	delete(book, "wrong")

	fmt.Println(len(book))
	fmt.Println(book)
}
