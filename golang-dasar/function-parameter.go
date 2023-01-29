package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Yahya"
	sayHelloTo(firstName, "Kurnia")
	sayHelloTo("Budi", "Nugraha")
}
