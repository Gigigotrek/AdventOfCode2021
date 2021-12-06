package main

import (
	"os"
	"fmt"
	"strconv"
	"bufio"
	"strings"
	"io"
)

func ReadStr(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []string
	for scanner.Scan() {
			x := scanner.Text()
			result = append(result, x)
	}
	return result, scanner.Err()
}

func count_bits(lines []string) (string, string) {
	var gamma_rate string
	var epsilon_rate string

	for vectors_size := 0; vectors_size < len(lines[0]); vectors_size++ {
		var binary_1_count int = 0
		var binary_0_count int = 0
		for vector := 0; vector < len(lines); vector++ {
			if string(lines[vector][vectors_size]) == "0" {
				binary_0_count++
			} else {
				binary_1_count++
			}
		}
		if binary_0_count > binary_1_count {
			gamma_rate += "0"
			epsilon_rate += "1"
		} else {
			gamma_rate += "1"
			epsilon_rate += "0"
		}
	}

	return gamma_rate, epsilon_rate
}

func bin_to_Dec(binary_value string) int {

	var dec_value, _= strconv.ParseInt(binary_value, 2, 64)
	return int(dec_value)
}

func main() {

	file, _ := os.ReadFile("./3-1-input")
	lines, _ := ReadStr(strings.NewReader(string(file)))

	var gamma_rate, epsilon_rate string = count_bits(lines)

	var gamma_rate_dec int = bin_to_Dec(gamma_rate)
	var epsilon_rate_dec int  = bin_to_Dec(epsilon_rate)

	var power int = gamma_rate_dec * epsilon_rate_dec

	fmt.Println(power)
}