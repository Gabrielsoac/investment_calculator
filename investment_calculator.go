package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func calculateInvestment (
	investmentAmount float64,
	expectedReturnRate float64,
	periodInYears float64) (float64, error) {
		
		if investmentAmount <= 0 {
			return 0.0, errors.New("investment amount cannot be zero");
		}
		
		if expectedReturnRate <= 0 {
			return 0.0, errors.New("expected return rate cannot be zero")
		}
		
		if periodInYears <= 0 {
			return 0.0, errors.New("period in years cannot be zero")
		}

		const inflationRate float64 = 2.5;
		
		futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, periodInYears)
		futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, periodInYears);
		
		fmt.Printf("investment amount: R$%.2f\n", investmentAmount);
		fmt.Printf("Future value: R$%.2f\n", futureValue);
		fmt.Printf("Future value with inflationRate: R$%.2f\n", futureRealValue);
		println("####### END #########")

		return 0, nil;
}

func profitCalculator(
	revenue float64,
	expenses float64,
	taxRate float64) {

	var ebt float64 = revenue - expenses;
	var profit = ebt * (1 - taxRate / 100 );
	var ratio = ebt / profit;

	fmt.Printf("\n####### PROFIT RESULT #########\nYour ebt: R$%.2f\nYour profit: R$%.2f\nYour ratio: R$%.2f\n", ebt, profit, ratio,
	);
}

func getFloatValue(fieldName string) (value float64, err error) {
	fmt.Printf("Enter %s value: ", fieldName);
	fmt.Scan(&value);
	if value <= 0 {
		err = errors.New(fieldName + " cannot be zero or negative value ❌");
	}
	return;
}

func main() {
	var choice = 0;

	for (choice != 3) {
		fmt.Println("\n####### MENU #######");
		fmt.Println("(1) Profit Calculator");
		fmt.Println("(2) Investment Calculator");
		fmt.Println("(3) Exit");
		
		fmt.Print("Enter your choice: ");
		fmt.Scan(&choice);

		if(choice == 1){
			var err error;
			renenue, err := getFloatValue("reneue");
			if err != nil {
				log.Print(err.Error());
				continue;
			}
			expenses, err := getFloatValue("expenses");
			if err != nil {
				log.Print(err.Error());
				continue;
			}
			taxRate, err := getFloatValue("tax rate");
			if err != nil {
				log.Print(err.Error());
				continue;
			}
			
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
		
			var value, err = calculateInvestment(investmentAmount, expectedReturnRate, periodInYears);
			
			if err != nil {
				log.Fatal(err.Error())
			}

			if value == 0 {
				println("ok")
			}

			fmt.Println(string(checkMark));
			choice = 0;
		}

		if(choice == 3){
			goodbyeMessage := fmt.Sprintln("Goodbye :)");
			fmt.Print(goodbyeMessage);
		}
	}
}
