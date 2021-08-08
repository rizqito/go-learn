package main

import (
	"fmt"
	"os" //untuk input input kali ya
)

func main(){
	var s string
	sep := ""
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}