package day8

import (
	"aoc2021/util"
	"fmt"
	"sort"
	"strings"
)

var validPatterns = map[string]int{
	"36":      1,
	"13457":   2,
	"13467":   3,
	"2346":    4,
	"12467":   5,
	"124567":  6,
	"136":     7,
	"1234567": 8,
	"123467":  9,
}

func parseLine(line string) (left []string, right []string) {
	parts := strings.Split(line, "|")
	left = strings.Split(strings.TrimSpace(parts[0]), " ")
	right = strings.Split(strings.TrimSpace(parts[1]), " ")
	return
}

func count(strs []string) int {
	sum := 0
	for _, str := range strs {
		switch len(str) {
		case 2, 3, 4, 7:
			sum += 1
		}
	}

	return sum
}

func Part1() {
	lines := util.GetLines("day8/input")
	sum := 0
	for _, line := range lines {
		_, right := parseLine(line)
		sum += count(right)
	}
	fmt.Println(sum)
}

func Part2() {
	lines := util.GetLines("day8/input")
	sum := 0
	for _, line := range lines {
		left, right := parseLine(line)
		decoded := processLine(left)
		inverted := invertMap(decoded)

		num := 0
		for _, rightItem := range right {
			num = num*10 + inverted[sortWord(rightItem)]
		}

		sum += num
	}
	fmt.Println(sum)
}

func invertMap(m map[int]string) (res map[string]int) {
	res = make(map[string]int)
	for k, v := range m {
		res[v] = k
	}
	return
}

func processLine(words []string) (decoded map[int]string) {
	sortByLen(words)
	decoded = make(map[int]string)
	for _, word := range words {
		getDigit(word, decoded)
	}
	return
}

func sortByLen(words []string) {
	order := map[int]int{
		2: 1,
		3: 1,
		4: 1,
		7: 1,
		6: 2,
		5: 3,
	}

	sort.Slice(words, func(i, j int) bool {
		return order[len(words[i])] < order[len(words[j])]
	})
}

func sortWord(word string) string {
	sorted := strings.Split(word, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

func getDigit(word string, known map[int]string) int {
	mark := func(k int) int {
		known[k] = sortWord(word)
		return k
	}

	switch len(word) {
	case 2:
		return mark(1)
	case 3:
		return mark(7)
	case 4:
		return mark(4)
	case 7:
		return mark(8)
	case 6:
		if contains(word, known[4]) {
			return mark(9)
		}
		if contains(word, known[7]) {
			return mark(0)
		}
		return mark(6)
	case 5:
		if contains(word, known[7]) {
			return mark(3)
		}
		if contains(known[6], word) {
			return mark(5)
		}
		return mark(2)
	}

	return -1
}

func contains(str1, str2 string) bool {
	if len(str1) < len(str2) {
		return false
	}

	dict1 := make(map[rune]bool)
	for _, ch := range str1 {
		dict1[ch] = true
	}

	for _, ch := range str2 {
		_, ok := dict1[ch]
		if !ok {
			return false
		}
	}

	return true
}
