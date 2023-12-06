package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/fbegyn/aoc2023/go/helpers"
)

func main() {
	file := os.Args[1]
	input := helpers.InputToLines(file)

	timing := []float64{}
	distance := []float64{}

	for _, e := range strings.Split(input[0], " ")[1:] {
		if e == "" {
			continue
		}
		t := float64(helpers.Atoi(strings.TrimSpace(e)))
		timing = append(timing, t)
	}
	for _, e := range strings.Split(input[1], " ")[1:] {
		if e == "" {
			continue
		}
		d := float64(helpers.Atoi(strings.TrimSpace(e)))
		distance = append(distance, d)
	}

	fmt.Println(timing)
	fmt.Println(distance)

	margin := []int{}

	for i := range timing {
		time := timing[i]
		minSpeed, maxSpeed := float64(0), float64(time)
		minDistance, maxDistance := math.Ceil(minSpeed*(time-minSpeed)), math.Floor(maxSpeed*(time-maxSpeed))
		for minDistance <= distance[i] || maxDistance <= distance[i] {
			if minDistance <= distance[i] {
				minSpeed++
			}
			if maxDistance <= distance[i] {
				maxSpeed--
			}
			minDistance, maxDistance = math.Ceil(minSpeed*(time-minSpeed)), math.Floor(maxSpeed*(time-maxSpeed))
		}

		margin = append(margin, int(maxSpeed-minSpeed+1))
	}

	fmt.Println(margin)
	fmt.Println(helpers.Mult(margin))
}

//func calculateMinSpeed(d, t int) {
//	minSpeed := math.Ceil(float64(d) / float64(t))
//	minTime := math.Ceil(float64(d) / float64(minSpeed))
//
//	fmt.Println(minSpeed, minTime)
//}
