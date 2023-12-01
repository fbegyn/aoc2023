package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fbegyn/aoc2023/go/helpers"
)

func main() {
	file := os.Args[1]
	input := helpers.InputToLines(file)
	sum := 0

	stringMap := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
		"zero":  '0',
	}
	for _, inp := range input {
		numbers := []rune{}
		for i, r := range inp {
			if '0' <= r && r <= '9' {
				numbers = append(numbers, r)
			}
			for k, v := range stringMap {
			        if strings.Index(inp[i:], k) == 0 {
					numbers = append(numbers, rune(v))
				}
			}
		}
		fmt.Printf("The numbers in %s are: %v\n", inp, numbers)
		calibrationNumber := fmt.Sprintf("%c%c", numbers[0], numbers[len(numbers)-1])
		sum += helpers.Atoi(calibrationNumber)
	}

	fmt.Printf("The sum of calibration values is: %d\n", sum)
}
