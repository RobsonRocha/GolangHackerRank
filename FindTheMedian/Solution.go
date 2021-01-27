package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// MergeSort performs the merge sort algorithm taking advantage of multiple processors.
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

// Complete the findMedian function below.
func findMedian(arr []int32) int32 {

	mergeSort(arr)

	if len(arr) %2 == 0{
		i := int(len(arr)/2)
		j := int(len(arr)/2)+1
		return int32((arr[i]+arr[j])/2)
	}
	i := int(len(arr)/2)

	return arr[i]

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

	result := findMedian(arr)

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
