package main

import (
	"fmt"
	"os"
	"time"
)

const historyFileName = "history.txt";

func writeCalculateHistoryDate(){
	os.WriteFile(historyFileName, []byte("Last calc — date: " + fmt.Sprint(time.Now().Format(time.RFC3339))), 0644);
}

func getLastCalcDate() string {
	data, _ := os.ReadFile(historyFileName);
	var lastCalcDate = string(data);
	return lastCalcDate;
}