package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	emptySpace string = " "
	hashTag    string = "#"
)

// Complete the staircase function below.
func staircase(n int32) {
	var i int32
	for i = 1; i <= n; i++ {
		emptySpaceNumber := n - i
		hashTagNumbr := i
		resultEmptySpace := formedString(emptySpace, emptySpaceNumber)
		resultHashTag := formedString(hashTag, hashTagNumbr)
		staircaseResult := fmt.Sprintf("%s%s", resultEmptySpace, resultHashTag)
		fmt.Println(staircaseResult)
	}

}

func formedString(val string, n int32) string {
	var text bytes.Buffer
	var i int32
	for i = 1; i <= n; i++ {
		text.WriteString(val)
	}
	return text.String()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
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
