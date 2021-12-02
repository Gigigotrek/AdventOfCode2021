package main

import (
    "os"
    "bufio"
    "fmt"
    "io"
    "strconv"
    "strings"
)

func ReadInts(r io.Reader) ([]int, error) {
    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)
    var result []int
    for scanner.Scan() {
        x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, scanner.Err()
}

func main() {
    tf, _ := os.ReadFile("./1-1-input")
    ints, _ := ReadInts(strings.NewReader(string(tf)))

	increased_times := 0
    for i := 1; i < len(ints); i++ {
        if ints[i] > ints[i-1] {
            increased_times++
        }
    }
    fmt.Println(increased_times)
}