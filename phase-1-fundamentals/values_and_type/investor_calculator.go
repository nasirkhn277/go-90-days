package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 5.0
	var investmentAmount float64
	years, expectedReturn := 10.0, 5.5

	fmt.Print("Enter Investmtnt Amount: ")
	fmt.Scan(&investmentAmount)

	futureValue := (investmentAmount * math.Pow((1+expectedReturn/100), years))
	futureRealvalue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Println(futureValue)
	fmt.Println(futureRealvalue)
}
