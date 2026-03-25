package main

import (
	"fmt"
	"math"
)

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