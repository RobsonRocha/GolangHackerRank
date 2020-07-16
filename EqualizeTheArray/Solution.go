package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"math"
)

// Complete the equalizeArray function below.
func equalizeArray(arr []int32) int32 {

	m := make(map[int32]int32)

	var max int32 = 0

	m[arr[0]] = 1

	for i:=1; i < len(arr); i++ {
		value, found := m[arr[i]]
		if !found{
			m[arr[i]] = 1
		}else{
			m[arr[i]] = value+1
			if max < value+1{
				max = value +1
			}
		}
	}

	var result int32 = int32(len(arr)) -1
	if max != 0{
		result = int32(len(arr)) - max
	}

	return int32(math.Abs(float64(result)))

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := equalizeArray(arr)

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
