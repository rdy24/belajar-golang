package main

import "fmt"

func getFullName() (string, string, string) {
	return "Yahya", "Kurnia", "Joko"
}

func main() {
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)
}
