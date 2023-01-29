package main

import "fmt"

func main() {
	var name = "Kurnia"

	if name == "Yahya" {
		fmt.Println("Hello Yahya")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
