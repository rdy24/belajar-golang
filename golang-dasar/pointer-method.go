package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	Yahya := Man{"Yahya"}
	Yahya.Married()

	fmt.Println(Yahya.Name)
}
