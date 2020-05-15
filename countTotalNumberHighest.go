package main

import (
	"fmt"
)

// Complete the birthdayCakeCandles function below.
func countPersonHigh(ar []int32) int32 {
	highest := getHighest(ar)
	return countHeighest(ar, highest)
}

func getHighest(input []int32) int32 {
	var highest int32
	for _, v := range input {
		if highest < v {
			highest = v
		}
	}
	return highest
}

func countHeighest(input []int32, highest int32) int32 {
	var count int32
	for _, v := range input {
		if highest == v {
			count = count + 1
		}
	}
	return count
}

func main() {
	ar := []int32{2, 3, 5, 5, 5}
	result := countPersonHigh(ar)
	fmt.Println(result)
}
