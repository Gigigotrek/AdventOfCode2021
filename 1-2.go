package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	tf, _ := os.ReadFile("./1-2-input")
	ints, _ := ReadInts(strings.NewReader(string(tf)))

	increased_times := 0
	for i := 2; i < len(ints)-1; i++ {
		first_window := ints[i] + ints[i-1] + ints[i-2]
		second_window := ints[i] + ints[i-1] + ints[i+1]
		if second_window > first_window {
			increased_times++
		}
	}
	fmt.Println(increased_times)
}
