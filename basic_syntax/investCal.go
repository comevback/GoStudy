package main

import (
	"fmt"
	"math"
)

func investCal() {
	const inflationRate = 2.5

	fmt.Println("Investment Calculator:")
	var nowAmount float64
	fmt.Print("Enter the amount you want to invest: ")
	fmt.Scan(&nowAmount)

	var years float64
	fmt.Print("Enter the number of years you want to invest: ")
	fmt.Scanln(&years)

	var interestRate float64
	fmt.Print("Enter the interest rate: ")
	fmt.Scanln(&interestRate)

	var futureValue = nowAmount * math.Pow((1+(interestRate/100)), years)
	var inflation = futureValue / math.Pow((1+(inflationRate/100)), years)
	fmt.Println("The future value of your investment is: ", futureValue)
	fmt.Print("The future value of your investment after inflation is: ", inflation)
}
