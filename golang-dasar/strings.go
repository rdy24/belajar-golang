package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Yahya Kurnia", "Yahya"))
	fmt.Println(strings.Contains("Yahya Kurnia", "Budi"))

	fmt.Println(strings.Split("Yahya Kurniawna Joko", " "))

	fmt.Println(strings.ToLower("Yahya Kurnia Joko"))
	fmt.Println(strings.ToUpper("Yahya Kurnia Joko"))
	fmt.Println(strings.ToTitle("Yahya Kurnia Joko"))

	fmt.Println(strings.Trim("      Yahya Kurnia     ", " "))
	fmt.Println(strings.ReplaceAll("Yahya Joko Yahya", "Yahya", "Budi"))
}
