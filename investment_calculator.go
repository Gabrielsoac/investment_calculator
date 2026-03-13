package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

func writeCalculateHistoryDate(){
	os.WriteFile("history.txt", []byte("Last calc — date: " + fmt.Sprint(time.Now().Format(time.RFC3339))), 0644);
}

func calculateFutureValue (investmentAmount float64, expectedReturnRate float64, periodInYears float64) float64 {
	 return investmentAmount * math.Pow(1 + expectedReturnRate / 100, periodInYears);
}

func calculateRealFutureValue (futureValue float64, periodInYears float64) float64 {
	const inflationRate float64 = 2.5;
	return futureValue / math.Pow(1 + inflationRate / 100, periodInYears);

}

func calculateInvestment (
	investmentAmount float64,
	expectedReturnRate float64,
	periodInYears float64) {
		
	futureValue := calculateFutureValue(investmentAmount, expectedReturnRate, periodInYears);
	futureRealValue := calculateRealFutureValue(futureValue, periodInYears);
	
	fmt.Printf("investment amount: R$%.2f\n", investmentAmount);
	fmt.Printf("Future value: R$%.2f\n", futureValue);
	fmt.Printf("Future value with inflationRate: R$%.2f\n", futureRealValue);
}

func profitCalculator(
	revenue float64,
	expenses float64,
	taxRate float64) {

	var ebt float64 = calculateEbt(revenue, expenses); 
	var profit = calculateProfit(ebt, taxRate);
	var ratio = calculateRatio(ebt, profit);


	writeCalculateHistoryDate()
	fmt.Printf("\n####### PROFIT RESULT #########\nYour ebt: R$%.2f\nYour profit: R$%.2f\nYour ratio: R$%.2f\n", ebt, profit, ratio,
	);
}

func calculateEbt (revenue float64, expenses float64) (ebt float64) {
	return revenue - expenses;
}

func calculateProfit (ebt float64, taxRate float64) (profit float64) {
	return ebt * (1 - taxRate / 100);
}

func calculateRatio (ebt float64, profit float64) (ratio float64) {
	return ebt / profit;
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
		} else if(choice == 2){
			checkMark := rune(0x2705);

			var err error;
			investmentAmount, err := getFloatValue("investment amount");
			if err != nil {
				log.Print(err.Error());
				continue;
			}
			expectedReturnRate, err := getFloatValue("expected return rate (example: 5.5)");
			if err != nil {
				log.Print(err.Error());
				continue;
			}
			periodInYears, err := getFloatValue("period in years");
			if err != nil {
				log.Print(err.Error());
				continue;
			}
		
			calculateInvestment(investmentAmount, expectedReturnRate, periodInYears);

			fmt.Println(string(checkMark));
			choice = 0;
		} else if(choice == 3){
			goodbyeMessage := fmt.Sprintln("Goodbye :)");
			fmt.Print(goodbyeMessage);
		} else {
			fmt.Println("Please, enter the valid option! :) ")
		}
	}
}
