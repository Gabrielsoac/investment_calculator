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
		println("####### END #########")
}

func profitCalculator(
	revenue float64,
	expenses float64,
	taxRate float64) {

	var ebt float64 = revenue - expenses;
	var profit = ebt * (1 - taxRate / 100 );
	var ratio = ebt / profit;

	fmt.Printf("Your ebt: %.2f\n", ebt);
	fmt.Printf("Your profit: %.2f\n", profit);
	fmt.Printf("Your ratio: %.2f\n", ratio);
	println("####### END #########")
}

func main() {
	var choice = 0;

	for (choice != 3) {
		fmt.Println("(1) Profit Calculator");
		fmt.Println("(2) Investment Calculator");
		fmt.Println("(3) Exit");
		
		fmt.Print("Enter your choice: ");
		fmt.Scan(&choice);

		if(choice == 1){
			var renenue, expenses, taxRate float64;
			
			fmt.Print("Enter your revenue: ")
			fmt.Scan(&renenue);
			
			fmt.Print("Enter your expenses: ");
			fmt.Scan(&expenses);
			
			fmt.Print("Enter the tax rate: ");
			fmt.Scan(&taxRate);
			
			profitCalculator(renenue, expenses, taxRate);
		}

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
