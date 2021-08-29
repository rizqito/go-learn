package main

import "fmt"

type Blacklist func(string) bool

func registeredUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func blacklistAdmin(name string) bool{
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool{
// 	return name == "root"
// }

func anonymousFunc() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registeredUser("admin", blacklist)
	registeredUser("eko", blacklist)
}
