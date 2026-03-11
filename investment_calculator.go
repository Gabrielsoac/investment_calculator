package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5;
	investmentAmount, expectedReturnRate, years := 0.0, 0.0, 0.0;
	checkMark := rune(0x2705);

	fmt.Print("Enter your invesmentAmount: ")
	fmt.Scan(&investmentAmount);

	fmt.Print("Enter yout expected return rate, example: '5.5': ");
	fmt.Scan(&expectedReturnRate);

	fmt.Print("Enter the period in years: ");
	fmt.Scan(&years);

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years);
	
	fmt.Printf("investment amount: %.2f\n", investmentAmount);
	fmt.Printf("Future value: %.2f\n", futureValue);
	fmt.Printf("Future value with inflationRate: %.2f\n", futureRealValue);
	fmt.Println(string(checkMark));
}
