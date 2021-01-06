package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"math"
)

// Complete the encryption function below.
func encryption(s string) string {
	root := math.Sqrt(float64(len(s)))

	a := math.Floor(root)
	b := math.Ceil(root)

	row := int(a)
	column := int(b)

	if row * column < len(s){
		row ++
	}

	firstStep := make([]string, 0)
	j := 0
	for i:=0; i < row; i++{
		if column + j < len(s) {
			firstStep = append(firstStep, s[j:column+j])
		}else{
			firstStep = append(firstStep, s[j:])
		}
		j += column
	}

	secondStep := ""
	for i:=0; i < column; i++{
		for j :=0; j < row; j++{
			if i < len(firstStep[j]) {
				secondStep = secondStep + string(firstStep[j][i])
			}
		}
		secondStep = secondStep + " "
	}

	return secondStep
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	s := readLine(reader)

	result := encryption(s)

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
