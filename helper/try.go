package helper

import (
	"fmt"
	"os" //untuk input input kali ya
)

func Try() {
	var s string
	sep := ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
