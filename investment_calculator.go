package main

import (
	"fmt"
	"math"
)

func calculateInvestment (
	investmentAmount float64,
	expectedReturnRate float64,
	periodInYears float64) {

	const inflationRate float64 = 2.5;

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, periodInYears)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, periodInYears);
	
	fmt.Printf("investment amount: %.2f\n", investmentAmount);
	fmt.Printf("Future value: %.2f\n", futureValue);
	fmt.Printf("Future value with inflationRate: %.2f\n", futureRealValue);
}

func main() {
	var choice = 0;

	for (choice != 3) {
		fmt.Println("(1) Profit Calculator");
		fmt.Println("(2) Investment Calculator");
		fmt.Println("(3) Exit");

		fmt.Print("Enter your choice: ");
		fmt.Scan(&choice);

		if(choice == 2){
			checkMark := rune(0x2705);
			investmentAmount, expectedReturnRate, periodInYears := 0.0, 0.0, 0.0;
			
			fmt.Print("Enter your invesmentAmount: ")
			fmt.Scan(&investmentAmount);
			
			fmt.Print("Enter your expected return rate, example: '5.5': ");
			fmt.Scan(&expectedReturnRate);
			
			fmt.Print("Enter the period in periodInYears: ");
			fmt.Scan(&periodInYears);
		
			calculateInvestment(investmentAmount, expectedReturnRate, periodInYears);
		
			fmt.Println(string(checkMark));
			choice = 0;
		}

		if(choice == 3){
			println("Goodbye :)");
		}
	}
	

	
}
