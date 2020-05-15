package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	newArr := convertInt32ToInt64(arr)
	var result []int64
	for i, _ := range newArr {
		result = append(result, sum(i, newArr))
	}
	min, max := getMinMax(result)
	fmt.Println(min, max)
}

func convertInt32ToInt64(input []int32) []int64 {
	var arrInt64 []int64
	for _, v := range input {
		arrInt64 = append(arrInt64, int64(v))
	}
	return arrInt64
}

func sum(index int, input []int64) int64 {
	var sum int64
	for i, v := range input {
		if i == index {
			continue
		}
		sum = sum + v
	}
	return sum
}

func getMinMax(input []int64) (min int64, max int64) {
	min = input[0]
	for _, v := range input {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return min, max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
