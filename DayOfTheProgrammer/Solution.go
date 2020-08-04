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

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
	if year == 1918 {
		return "26.09.1918"
	}
	leap4 := math.Mod(float64(year),float64(4))
	leap400 := math.Mod(float64(year),float64(400))
	february := 28
	if year >= 1700 && year <= 1917 && leap4 == 0{
		february = 29
	} else if leap400 ==0 {
		february = 29
	}else if leap4 == 0 && math.Mod(float64(year),float64(100)) > 0 {
		february = 29
	}
	sum := 215+february

	day := 256 - sum

	result := fmt.Sprintf("%d.09.%d", day, year)

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

	fmt.Fprintf(writer, "%s\n", result)

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
