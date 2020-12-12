package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

)

/*
 * Complete the 'stockmax' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY prices as parameter.
 */

func Max(prices []int32) int32{
	max := int32(-1)

	for _, v := range prices{
		if max <= v {
			max = v
		}
	}

	return max
}

func stockmax(prices []int32) int64 {
	// Write your code here
	result := 0

	max := Max(prices)

	for i:=0; i < len(prices); i++{
		if i + 1 < len(prices) && prices[i] < max {
			result += int(max - prices[i])
		}else{
			max = Max(prices[i+1:])
		}
	}

	if result < 0{
		return 0
	}

	return int64(result)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		pricesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var prices []int32

		for i := 0; i < int(n); i++ {
			pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
			checkError(err)
			pricesItem := int32(pricesItemTemp)
			prices = append(prices, pricesItem)
		}

		result := stockmax(prices)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
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
