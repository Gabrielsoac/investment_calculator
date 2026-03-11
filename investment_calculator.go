package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5;
	investmentAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0;
	checkMark := rune(0x2705);

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years);
	
	fmt.Printf("%.2f\n", futureValue);
	fmt.Println(string(checkMark));
	fmt.Printf("%.2f\n", futureRealValue);
}
