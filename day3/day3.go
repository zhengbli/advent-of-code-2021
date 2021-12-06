package day3

import (
	"aoc2021/util"
	"fmt"
	"strconv"
)

func Part1() {
	counts := []int{}

	for _, line := range util.GetLines("day3/input") {
		for i, ch := range line {
			var num int
			if ch == '1' {
				num = 1
			} else {
				num = -1
			}

			if len(counts) <= i {
				counts = append(counts, num)
			} else {
				counts[i] += num
			}

		}
	}

	str := ""
	comp := ""
	for _, c := range counts {
		if c > 0 {
			str += "1"
			comp += "0"
		} else {
			str += "0"
			comp += "1"
		}
	}

	strVal, _ := strconv.ParseInt(str, 2, 32)
	compVal, _ := strconv.ParseInt(comp, 2, 32)
	fmt.Println(strVal * compVal)
}

func Part2() {
	lines := util.GetLines("day3/input")

	posGetDigit := func(count int) byte {
		if count >= 0 {
			return '1'
		}
		return '0'
	}
	pos := filterLine(lines, posGetDigit)

	negGetDigit := func(count int) byte {
		if count < 0 {
			return '1'
		}
		return '0'
	}
	neg := filterLine(lines, negGetDigit)

	posVal, _ := strconv.ParseInt(pos, 2, 32)
	negVal, _ := strconv.ParseInt(neg, 2, 32)
	fmt.Println(posVal * negVal)
}

func filterLine(lines []string, getDigit func(int) byte) string {
	bads := make(map[string]bool)
	digitCount := len(lines[0])
	restCount := len(lines)

	for i := 0; i < digitCount; i++ {
		ones := make(map[string]bool)
		zeros := make(map[string]bool)
		count := 0
		for _, line := range lines {
			if _, ok := bads[line]; ok {
				continue
			}

			if line[i] == '1' {
				ones[line] = true
				count += 1
			} else {
				zeros[line] = true
				count -= 1
			}
		}

		matchDigit := getDigit(count)
		if matchDigit == '1' {
			restCount -= len(zeros)
			addRange(bads, zeros)
		} else {
			restCount -= len(ones)
			addRange(bads, ones)
		}

		if restCount == 1 {
			break
		}
	}

	for _, line := range lines {
		if _, ok := bads[line]; !ok {
			return line
		}
	}

	return ""
}

func addRange(map1, map2 map[string]bool) {
	for k, v := range map2 {
		map1[k] = v
	}
}
