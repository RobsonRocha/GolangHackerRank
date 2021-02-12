package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"runtime"
)

func mergeSort(src []int32) {
	// We subtract 1 goroutine which is the one we are already running in.
	extraGoroutines := runtime.NumCPU() - 1
	semChan := make(chan struct{}, extraGoroutines)
	defer close(semChan)
	mergesort(src, semChan)
}

func mergesort(src []int32, semChan chan struct{}) {
	if len(src) <= 1 {
		return
	}

	mid := len(src) / 2

	left := make([]int32, mid)
	right := make([]int32, len(src)-mid)
	copy(left, src[:mid])
	copy(right, src[mid:])

	wg := sync.WaitGroup{}

	select {
	case semChan <- struct{}{}:
		wg.Add(1)
		go func() {
			mergesort(left, semChan)
			<-semChan
			wg.Done()
		}()
	default:
		// Can't create a new goroutine, let's do the job ourselves.
		mergesort(left, semChan)
	}

	mergesort(right, semChan)

	wg.Wait()

	merge(src, left, right)
}

func merge(result, left, right []int32) {
	var l, r, i int

	for l < len(left) || r < len(right) {
		if l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				result[i] = left[l]
				l++
			} else {
				result[i] = right[r]
				r++
			}
		} else if l < len(left) {
			result[i] = left[l]
			l++
		} else if r < len(right) {
			result[i] = right[r]
			r++
		}
		i++
	}
}

// Complete the maximumToys function below.
func maximumToys(prices []int32, k int32) int32 {
	mergeSort(prices)

	sum := int32(0)
	i := 0
	for _, v := range prices{
		if (sum + v) <= k {
			sum += v
			i++
		}else{
			break
		}
	}

	return int32(i)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	pricesTemp := strings.Split(readLine(reader), " ")

	var prices []int32

	for i := 0; i < int(n); i++ {
		pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
		checkError(err)
		pricesItem := int32(pricesItemTemp)
		prices = append(prices, pricesItem)
	}

	result := maximumToys(prices, k)

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
