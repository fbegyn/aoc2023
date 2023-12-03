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
	ids := []int{}
	s := 0

	// iterate over each game
	for _, inp := range input {
		words := strings.Split(inp, ":")
		id := strings.Split(words[0], " ")[1]
		showings := strings.Split(words[1], ";")
		possible := true
		fmt.Printf("Game ID: %s\n", id)
		gameMax := map[string]int{}
		for _, show := range showings {
			showCount := map[string]int{}
			sections := strings.Split(strings.TrimSpace(show), ", ")
			// iteratre over each showing of the dice
			for _, section := range sections {
				temp := strings.Split(section, " ")
				number, color := temp[0], temp[1]
				n := helpers.Atoi(number)
				showCount[color] += n

				if t, ok := gameMax[color]; ok && (t < n) {
					gameMax[color] = n
				} else if !ok {
					gameMax[color] = n
				}
			}
			if showCount["red"] > 12 {
				possible = false
			}
			if showCount["green"] > 13 {
				possible = false
			}
			if showCount["blue"] > 14 {
				possible = false
			}
			//if !possible {
			//	break
			//}
		}
		fmt.Println(gameMax)
		fmt.Println(helpers.MultMap(gameMax))
		s += helpers.MultMap(gameMax)
		if possible {
			ids = append(ids, helpers.Atoi(id))
		}
	}

	fmt.Println(ids)
	fmt.Println(helpers.Sum(ids))
	fmt.Println(s)
}
