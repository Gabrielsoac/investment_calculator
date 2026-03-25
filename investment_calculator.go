package main

import (
	"errors"
	"fmt"
	"log"
)

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

		showMenu();
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
		} else if(choice == 4){
			lastCalcDateMessage := getLastCalcDate();
			fmt.Println(lastCalcDateMessage);
		} else {
			fmt.Println("Please, enter the valid option! :) ")
		}
	}
}
