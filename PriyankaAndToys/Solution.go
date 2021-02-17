package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sort"
)

func unique(arr []int) []int {
	occured := map[int]bool{}
	result := []int{}
	for e := range arr {

		// check if already the mapped
		// variable is set to true or not
		if occured[int(arr[e])] != true {
			occured[int(arr[e])] = true

			// Append to result slice.
			result = append(result, arr[e])
		}
	}

	return result
}

// Complete the toys function below.
func toys(w []int) int {
	sort.Ints(w)
	wl := unique(w)
	min := wl[0]
	count := 1
	for _, v := range wl{
		if !(v <= (min+4)) {
			min = v
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	wTemp := strings.Split(readLine(reader), " ")

	var w []int

	for i := 0; i < int(n); i++ {
		wItemTemp, err := strconv.ParseInt(wTemp[i], 10, 64)
		checkError(err)
		wItem := int(wItemTemp)
		w = append(w, wItem)
	}

	result := toys(w)

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
