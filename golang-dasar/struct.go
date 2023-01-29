package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu(){
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var Yahya Customer
	Yahya.Name = "Yahya"
	Yahya.Address = "Indonesia"
	Yahya.Age = 30

	Yahya.sayHi("Joko")
	Yahya.sayHuuu()

	//fmt.Println(Yahya.Name)
	//fmt.Println(Yahya.Address)
	//fmt.Println(Yahya.Age)
	//
	//joko := Customer{
	//	Name:    "Joko",
	//	Address: "Cirebon",
	//	Age:     35,
	//}
	//fmt.Println(joko)
	//
	//budi := Customer{"Budi", "Jakarta", 35}
	//fmt.Println(budi)
}
