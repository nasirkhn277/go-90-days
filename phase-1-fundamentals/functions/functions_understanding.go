package main

import (
	"fmt"
	"math"
)

const inflationRate = 5.0

var investmentAmount float64
var years, expectedReturn = 10.0, 5.5

func main() {
	outputText("Enter your investment amount ")
	fmt.Scan(&investmentAmount)
	outputText("Enter your years ")
	fmt.Scan(&years)

	futureValue, futureRealvalue := calculateFutureValue(investmentAmount, expectedReturn, years)

	fmt.Println("Future Value: ", futureValue, "Real Future Value: ", futureRealvalue)

}

func outputText(text string) {
	fmt.Println(text)
}

func calculateFutureValue(investmentAmount, expectedReturn, years float64) (float64, float64) {
	futureValue := (investmentAmount * math.Pow((1+expectedReturn/100), years))
	futureRealvalue := futureValue / math.Pow((1+inflationRate/100), years)
	return futureValue, futureRealvalue
}
