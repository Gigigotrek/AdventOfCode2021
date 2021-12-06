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

func bin_to_Dec(binary_value string) int {

	var dec_value, _= strconv.ParseInt(binary_value, 2, 64)
	return int(dec_value)
}

func get_unique_value(values []string, column int, most_values bool) string {
	if len(values) == 1 { return values[0] }
	
	var binary_1_count int = 0
	var binary_0_count int = 0
	var v_0 []string
	var v_1 []string
	for row := 0; row < len(values); row++ {
		if string(values[row][column]) == "0" {
			v_0 = append(v_0, values[row])
			binary_0_count++
		} else {
			v_1 = append(v_1, values[row])
			binary_1_count++
		}
	}

	if (binary_0_count > binary_1_count && most_values == true) {
		values = v_0
	} else if (binary_0_count <= binary_1_count && most_values == true) {
		values = v_1
	} else if (binary_0_count > binary_1_count && most_values == false) {
		values = v_1
	} else if (binary_0_count <= binary_1_count && most_values == false) {
		values = v_0
	}

	return get_unique_value(values, column+1, most_values)
}

func main() {

	file, _ := os.ReadFile("./3-2-input")
	lines, _ := ReadStr(strings.NewReader(string(file)))

	var most_values, least_values []string = lines, lines

	var oxygen_generator_rate int = bin_to_Dec(get_unique_value(most_values, 0, true))
	var co2_scrubber_rate int  = bin_to_Dec(get_unique_value(least_values, 0, false))

	var life_support_rate int = oxygen_generator_rate * co2_scrubber_rate

	fmt.Println(life_support_rate)
}