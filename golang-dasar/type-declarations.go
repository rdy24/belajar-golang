package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpYahya NoKTP = "18741982741897419874"
	var marriedStatus Married = true
	fmt.Println(noKtpYahya)
	fmt.Println(marriedStatus)
}
