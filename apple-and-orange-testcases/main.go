package main

import (
	"fmt"
)

// Complete the countApplesAndOranges function below.
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	applesCoord := getFruitCoord(apples, a)
	orrangesCoord := getFruitCoord(oranges, a)
	fmt.Println(calcFruits(applesCoord, s, t))
	fmt.Println(calcFruits(orrangesCoord, s, t))
}

func calcFruits(input []int32, s int32, t int32) int32 {
	var total int32
	for _, v := range input {
		if s >= v && v <= t {
			total = total + 1
		}
	}
	return total
}

func getFruitCoord(fruits []int32, loc int32) []int32 {
	var output []int32
	for _, v := range fruits {
		output = append(output, v+loc)
	}
	return output
}

func main() {

	countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
}
