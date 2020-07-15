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

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	if len(s) == 1 && strings.Contains(strings.ToLower(s), "a"){
		return n
	}

	if !strings.Contains(strings.ToLower(s), "a"){
		return 0
	}

	var first int64 = n/int64(len(s))

	var second int64 = first*int64(strings.Count(strings.ToLower(s),"a"))

	var third float64 = math.Mod(float64(n), float64(len(s)))

	if third == 0{
		return second
	}else{
		return second+int64(strings.Count(strings.ToLower(s[:int(third)]),"a"))
	}


}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
