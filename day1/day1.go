package day1

import (
	"aoc2021/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	content, _ := os.ReadFile("day1/input")
	count := 0
	prevLine := 0

	for i, line := range strings.Split(string(content), "\n") {
		lineVal, _ := strconv.Atoi(line)
		if i > 0 && lineVal > prevLine {
			count += 1
		}
		prevLine = lineVal
	}

	fmt.Println(count)
}

func Part2() {
	count := 0
	q3 := []int{}
	lastSum := 0
	curSum := 0

	for _, line := range util.GetLines("day1/input") {
		lineVal, _ := strconv.Atoi(line)
		q3 = append(q3, lineVal)
		lastSum = curSum
		curSum += lineVal

		if len(q3) > 3 {
			curSum -= q3[0]
			if curSum > lastSum {
				count += 1
			}
			q3 = q3[1:]
		}
	}

	fmt.Println(count)
}
