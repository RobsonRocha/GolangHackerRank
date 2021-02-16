package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)


// Complete the twoArrays function below.
func twoArrays(k int, A []int, B []int) string {

	result := "YES"

	sort.Ints(A)
	sort.Sort(sort.Reverse(sort.IntSlice(B)))

	for i, vi := range A{
		if vi + B[i] < k {
			return "NO"
		}
	}


	return result

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nk := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nk[0], 10, 64)
		checkError(err)
		n := int(nTemp)

		kTemp, err := strconv.ParseInt(nk[1], 10, 64)
		checkError(err)
		k := int(kTemp)

		ATemp := strings.Split(readLine(reader), " ")

		var A []int

		for i := 0; i < int(n); i++ {
			AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
			checkError(err)
			AItem := int(AItemTemp)
			A = append(A, AItem)
		}

		BTemp := strings.Split(readLine(reader), " ")

		var B []int

		for i := 0; i < int(n); i++ {
			BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
			checkError(err)
			BItem := int(BItemTemp)
			B = append(B, BItem)
		}

		result := twoArrays(k, A, B)

		fmt.Fprintf(writer, "%s\n", result)
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
