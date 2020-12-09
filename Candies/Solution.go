package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the candies function below.
func candies(n int32, arr []int32) int64 {

	prev := arr[0]
	candies := make([]int,0)
	result := 0

	for i :=0 ; i< len(arr); i++{
		candies = append(candies,1)
	}

	for i:=1; i < len(arr); i++{
		if prev < arr[i] {
			candies[i] = candies[i-1]+1
		}
		prev = arr[i]
	}

	prev = arr[len(arr)-1]
	for i:=len(arr)-2; i >= 0 ; i--{
		if prev < arr[i] && candies[i] < candies[i+1]+1{
			candies[i] = candies[i+1]+1
		}
		prev = arr[i]
	}

	for _, v:=range(candies){
		result += v
	}
	fmt.Printf("%v", candies)
	return int64(result)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout := os.Stdin
	//checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := candies(n, arr)

	fmt.Fprintf(writer, "%d\n", result)

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
