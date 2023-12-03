package main

import (
	"fmt"
	"os"

	"github.com/fbegyn/aoc2023/go/helpers"
)

var (
	 schematic map[helpers.Point]int
	 t map[helpers.Point][]helpers.Point
)

func main() {

	file := os.Args[1]

	runeCh := make(chan rune)
	var x, y int64
	x, y = 0, 0
	//schematicStr := map[helpers.Point]rune{}
	schematic = map[helpers.Point]int{}
	symbolIndex := []helpers.Point{}
	numberTracker, numberTrackerInd := []rune{}, []helpers.Point{}
	t = map[helpers.Point][]helpers.Point{}

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
				t[p] = numberTrackerInd
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

	fmt.Println(schematic)
	s := 0
	for _, p := range symbolIndex {
		s += LookPartNumber(p)
	}
	fmt.Println(schematic)
	fmt.Printf("Sum of part numbers is %d\n", s)
}

func LookPartNumber(p helpers.Point) int {
	fmt.Println(p)
	sum := 0
	// check center of the symbol
	cuk, cdk := false, false
	cuc := helpers.Point{p.X, p.Y-1}
	if cu, ok := schematic[cuc]; ok {
		cuk = ok
		fmt.Println(cu)
		sum += cu
		for _, v := range t[cuc] {
			delete(schematic, v)
		}
	}
	cdc := helpers.Point{p.X, p.Y+1}
	if cd, ok := schematic[cdc]; ok {
		cdk = ok
		fmt.Println(cd)
		sum += cd
		for _, v := range t[cdc] {
			delete(schematic, v)
		}
	}
	// check left of the symbol
	luc := helpers.Point{p.X-1, p.Y-1}
	if lu, ok := schematic[luc]; ok && ok != cuk {
		fmt.Println(lu)
		sum += lu
		for _, v := range t[luc] {
			delete(schematic, v)
		}
	}
	lc := helpers.Point{p.X-1, p.Y}
	if l, ok := schematic[lc]; ok {
		fmt.Println(l)
		sum += l
		for _, v := range t[lc] {
			delete(schematic, v)
		}
	}
	ldc := helpers.Point{p.X-1, p.Y+1}
	if ld, ok := schematic[ldc]; ok && ok != cdk {
		fmt.Println(ld)
		sum += ld
		for _, v := range t[ldc] {
			delete(schematic, v)
		}
	}
	// check right of the symbol
	ruc := helpers.Point{p.X+1, p.Y-1}
	if ru, ok := schematic[ruc]; ok && ok != cuk {
		fmt.Println(ru)
		sum += ru
		for _, v := range t[ruc] {
			delete(schematic, v)
		}
	}
	rc := helpers.Point{p.X+1, p.Y}
	if r, ok := schematic[rc]; ok {
		fmt.Println(r)
		sum += r
		for _, v := range t[rc] {
			delete(schematic, v)
		}
	}
	rdc := helpers.Point{p.X+1, p.Y+1}
	if rd, ok := schematic[rdc]; ok && ok != cdk {
		fmt.Println(rd)
		sum += rd
		for _, v := range t[rdc] {
			delete(schematic, v)
		}
	}
	return sum
}
