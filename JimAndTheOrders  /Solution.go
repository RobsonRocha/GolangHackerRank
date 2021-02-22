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

type order struct {
	Num int32
	Time int32
}

type orderLine []order

func (o orderLine) Len() int           { return len(o) }
func (o orderLine) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o orderLine) Less(i, j int) bool { return o[i].Time < o[j].Time }

// Complete the jimOrders function below.
func jimOrders(orders [][]int32) []int32 {

	list := make([]order,0)

	result := make([]int32, 0)

	for i, v := range orders{
		o := order{
			Num: int32(i+1),
			Time: v[0] + v[1],
		}
		list = append(list, o)
	}

	sort.Sort(orderLine(list))

	for _, v := range list{
		r := v.Num

		result = append(result, r)
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout := os.Stdin//Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var orders [][]int32
	for i := 0; i < int(n); i++ {
		ordersRowTemp := strings.Split(readLine(reader), " ")

		var ordersRow []int32
		for _, ordersRowItem := range ordersRowTemp {
			ordersItemTemp, err := strconv.ParseInt(ordersRowItem, 10, 64)
			checkError(err)
			ordersItem := int32(ordersItemTemp)
			ordersRow = append(ordersRow, ordersItem)
		}

		if len(ordersRow) != int(2) {
			panic("Bad input")
		}

		orders = append(orders, ordersRow)
	}

	result := jimOrders(orders)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result) - 1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
