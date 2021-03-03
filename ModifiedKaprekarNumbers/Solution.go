package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func splitNumbers(num string)[]int32{
	result := make([]int32, 0)
	digits := 1
	length := len(num)

	if length > 1 {
		digits = len(num) / 2
	}

	value, _ := strconv.Atoi(substr(num, 0, digits))
	result = append(result, int32(value))
	rest, _ := strconv.Atoi(substr(num, digits, len(num)))
	result = append(result, int32(rest))


	return result

}

// Complete the kaprekarNumbers function below.
func kaprekarNumbers(p int32, q int32) {

	found := false

	for i := p; i<=q; i++{

		digits := len(strconv.Itoa(int(i)))

		sq := int64(i)*int64(i)

		str := strconv.Itoa(int(sq))

		if (2*digits) != len(str) && ((2*digits)-1)!=len(str){
			continue
		}

		arr := splitNumbers(str)

		count:=int32(0)

		for _, v := range arr{
			count += v
		}

		if count == i{
			fmt.Printf("%d ", i)
			found = true
		}

	}

	if !found {
		fmt.Print("INVALID RANGE")
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	kaprekarNumbers(p, q)
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
