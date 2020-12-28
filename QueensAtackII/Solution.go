package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func min(a, b int32) int32{
	if a <= b{
		return a
	}

	return b
}

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	if n <= 1{
		return int32(0)
	}

	result := int32(0)

	if k == 0{
		result += (n - r_q) + (r_q -1) + (n - c_q) + (c_q -1)
		diagInfRigth := min(n - c_q, r_q -1)
		diagInfLeft := min(c_q -1, r_q-1)
		diagSupRigth := min(n - r_q, n - c_q)
		diagSupLeft := min(n - r_q, c_q -1)
		result += diagInfRigth + diagInfLeft + diagSupRigth + diagSupLeft
		return result
	}else{
		diagInfRigth := min(n - c_q, r_q -1)
		diagInfLeft := min(c_q -1, r_q-1)
		diagSupRigth := min(n - r_q, n - c_q)
		diagSupLeft := min(n - r_q, c_q -1)
		diffRowLeft := c_q -1
		diffRowRight := n - c_q
		diffColTop := n - r_q
		diffColBottom := r_q - 1

		for _, v := range obstacles{
			if v[0] == r_q {
				if v[1] < c_q {
					mini := (c_q -1) - v[1]
					diffRowLeft = min(diffRowLeft, mini)
				}else if v[1] > c_q {
					mini := n - c_q - (n - v[1]) -1
					diffRowRight = min(mini, diffRowRight)
				}
			}else if v[1] == c_q{
				if v[0] < r_q{
					mini := (r_q -1) - v[0]
					diffColBottom = min(mini, diffColBottom)
				}else{
					mini := n - r_q - (n - v[0])-1
					diffColTop = min(mini, diffColTop)
				}
			}else if math.Abs(float64(v[0]) - float64(r_q)) ==
				math.Abs(float64(v[1]) - float64(c_q)){

				if v[0] < r_q && v[1] < c_q {
					mini := min(int32(math.Abs(float64(c_q - v[1]))), int32(math.Abs(float64(r_q - v[0])))) -1
					diagInfLeft = min(mini,diagInfLeft)
				}else if v[0] > r_q && v[1] > c_q {
					mini := min(int32(math.Abs(float64(c_q - v[1]))), int32(math.Abs(float64(r_q - v[0])))) -1
					diagSupRigth = min(mini, diagSupRigth)
				}else if v[0] < r_q && v[1] > c_q {
					mini := min(int32(math.Abs(float64(c_q - v[1]))), int32(math.Abs(float64(r_q - v[0])))) -1
					diagInfRigth = min(mini, diagInfRigth)
				} else if v[0] > r_q && v[1] < c_q {
					mini := min(int32(math.Abs(float64(c_q - v[1]))), int32(math.Abs(float64(r_q - v[0])))) -1
					diagSupLeft = min(mini, diagSupLeft)
				}
			}
		}

		result = diffRowLeft+diffColBottom+diffColTop+diffRowRight+diagInfLeft+diagSupRigth+diagInfRigth+diagSupLeft
	}


	return result

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

	r_qC_q := strings.Split(readLine(reader), " ")

	r_qTemp, err := strconv.ParseInt(r_qC_q[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(r_qC_q[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(readLine(reader), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != int(2) {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

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
