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

	stringMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}
	for _, inp := range input {
		fmt.Printf("Starting word is: %-64s\n", inp)
		limit := len(inp)
		for i := 0; i < limit; i++ {
			stringLength := len(inp)
			end := i + 5
			if stringLength < end {
				end = stringLength
			}
			prefix := inp[:i]
			selection := inp[i:end]
			suffix := inp[end:]
			for k, v := range stringMap {
				selection = strings.Replace(selection, k, v, -1)
			}
			inp = prefix + selection + suffix
			limit = len(inp)
		}
		fmt.Printf("transformed into: %-64s\n\n", inp)

		numbers := []rune{}
		for _, r := range inp {
			if '0' <= r && r <= '9' {
				numbers = append(numbers, r)
			}
		}
		calibrationNumber := fmt.Sprintf("%c%c", numbers[0], numbers[len(numbers)-1])
		sum += helpers.Atoi(calibrationNumber)

	}

	fmt.Printf("The sum of calibration values is: %d\n", sum)
}
