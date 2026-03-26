package fileops

import (
	"fmt"
	"os"
	"time"
)

func WriteCalculateHistoryDateOnFile(fileName string){
	os.WriteFile(fileName, []byte("Last calc — date: " + fmt.Sprint(time.Now().Format(time.RFC3339))), 0644);
}

func ReadLastCalcDateOnFile(fileName string) string {
	data, _ := os.ReadFile(fileName);
	var lastCalcDate = string(data);
	return lastCalcDate;
}