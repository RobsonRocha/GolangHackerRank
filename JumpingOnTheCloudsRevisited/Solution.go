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

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32, k int32) int32 {
    i := 0
    energy := 100

    for {
        index := math.Mod(float64(i)+float64(k), float64(len(c)))
        cloud := int(c[int(index)])

        energy = energy - 1 - cloud*2

        if index == 0 {
            break
        }
        i= int(index)
    }
    return int32(energy)
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

    cTemp := strings.Split(readLine(reader), " ")

    var c []int32

    for i := 0; i < int(n); i++ {
        cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
        checkError(err)
        cItem := int32(cItemTemp)
        c = append(c, cItem)
    }

    result := jumpingOnClouds(c, k)

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
