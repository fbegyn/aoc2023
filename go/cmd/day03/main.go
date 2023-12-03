package main

import (
	"fmt"
	"os"

	"github.com/fbegyn/aoc2023/go/helpers"
)

func main() {

	file := os.Args[1]

	runeCh := make(chan rune)
	var x, y int64
	x, y = 0, 0
	//schematicStr := map[helpers.Point]rune{}
	schematic := map[helpers.Point]int{}
	symbolIndex := []helpers.Point{}
	numberTracker, numberTrackerInd := []rune{}, []helpers.Point{}

	// initialise stream of runes from the schematic
	go helpers.StreamRunes(file, runeCh)
	for r := range runeCh {
		coord := helpers.Point{x, y}
		if ( r < 'a' || 'Z' > r ) && ( r < '0' || '9' < r ) && r != '.' && r != 10 {
			symbolIndex = append(symbolIndex, coord)
			schematic[coord] = -1
		}
		//schematicStr[coord] = r

		if '0' <= r && r <= '9' {
			numberTracker = append(numberTracker, r)
			numberTrackerInd = append(numberTrackerInd, helpers.Point{x, y})

		} else if (r == '.' || r == 10) {
			for _, p := range numberTrackerInd {
				schematic[p] = helpers.Atoi(string(numberTracker))
			}
			numberTracker = []rune{}
			numberTrackerInd = []helpers.Point{}
		}

		x++
		if r == 10 {
			y++
			x = 0
		}

	}
	fmt.Println(symbolIndex)
	fmt.Println(schematic)

	s := 0
	for _, p := range symbolIndex {
		s += LookPartNumber(schematic, p)
	}
	fmt.Println(s)
}

func LookPartNumber(schematic map[helpers.Point]int, p helpers.Point) int {
	fmt.Println(p)
	sum := 0
	// check center of the symbol
	cuk, cdk := false, false
	if cu, ok := schematic[helpers.Point{p.X, p.Y-1}]; ok {
		fmt.Println("center up")
		cuk = ok
		fmt.Println(cu)
		sum += cu
	}
	if cd, ok := schematic[helpers.Point{p.X, p.Y+1}]; ok {
		fmt.Println("center down")
		cdk = ok
		fmt.Println(cd)
		sum += cd
	}
	// check left of the symbol
	if lu, ok := schematic[helpers.Point{p.X-1, p.Y-1}]; ok && ok != cuk {
		fmt.Println("left up")
		fmt.Println(lu)
		sum += lu
	}
	if l, ok := schematic[helpers.Point{p.X-1, p.Y}]; ok {
		fmt.Println("left")
		fmt.Println(l)
		sum += l
	}
	if ld, ok := schematic[helpers.Point{p.X-1, p.Y+1}]; ok && ok != cdk {
		fmt.Println("left down")
		fmt.Println(ld)
		sum += ld
	}
	// check right of the symbol
	if ru, ok := schematic[helpers.Point{p.X+1, p.Y-1}]; ok && ok != cuk {
		fmt.Println("right up")
		fmt.Println(ru)
		sum += ru
	}
	if r, ok := schematic[helpers.Point{p.X+1, p.Y}]; ok {
		fmt.Println("right")
		fmt.Println(r)
		sum += r
	}
	if rd, ok := schematic[helpers.Point{p.X+1, p.Y+1}]; ok && ok != cdk {
		fmt.Println("right down")
		fmt.Println(rd)
		sum += rd
	}
	return sum
}
