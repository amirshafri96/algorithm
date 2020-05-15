package main

import (
	"fmt"
	"time"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {

	test, err := time.Parse("15:04:05PM", s)
	if err != nil {
		panic(err)
	}
	fmt.Println(test.Format("15:04:05"))
	return s
}

func main() {

	result := timeConversion("07:05:45PM")

	fmt.Println(result)
}
