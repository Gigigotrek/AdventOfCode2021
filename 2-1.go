package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	var horizontal = 0
	var depth = 0

	file, _ := os.Open("./2-1-input")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		full_order := strings.Fields(scanner.Text())

		var order = full_order[0]
		order_value, _ := strconv.Atoi(full_order[1])

		switch order {
			case "forward":
				horizontal = horizontal + order_value
			case "down":
				depth = depth + order_value
			case "up":
				depth = depth - order_value
			}
	}

	fmt.Println(horizontal)
	fmt.Println(depth)
	fmt.Println(horizontal*depth)
}