package main

import "fmt"

func main() {
	var revenue float64
	var expense float64
	var rate float64
	fmt.Print("Enter your revenue ")
	fmt.Scan(&revenue)
	fmt.Print("Enter your expense ")
	fmt.Scan(&expense)
	fmt.Print("Enter your rate ")
	fmt.Scan(&rate)

	earningBeforeTax := revenue - expense
	earningAfterTax := earningBeforeTax - ((earningBeforeTax) * rate / 100)
	ratio := earningBeforeTax / earningAfterTax

	fmt.Println("EBT:", earningBeforeTax)
	fmt.Println("EAT", earningAfterTax)
	fmt.Println("Ratio:", ratio)
}
