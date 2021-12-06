package day2

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	x := 0
	depth := 0

	for _, line := range util.GetLines("day2/input") {
		parts := strings.Split(line, " ")
		delta, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			x += delta
		case "down":
			depth += delta
		case "up":
			depth -= delta
		}
	}

	fmt.Println(x * depth)
}

func Part2() {
	x := 0
	depth := 0
	aim := 0

	for _, line := range util.GetLines("day2/input") {
		parts := strings.Split(line, " ")
		delta, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			x += delta
			depth += delta * aim
		case "down":
			aim += delta
		case "up":
			aim -= delta
		}
	}

	fmt.Println(x * depth)
}
