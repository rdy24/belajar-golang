package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	helper.SayHello("Yahya")
	// helper.sayGoodbye("Yahya") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
