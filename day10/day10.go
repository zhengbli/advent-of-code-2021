package day10

import (
	"aoc2021/util"

	"github.com/montanaflynn/stats"
)

var points map[rune]int = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var part2Points map[rune]int = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

var match map[rune]rune = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

type Stack []rune

func (s Stack) push(val rune) Stack {
	return append(s, val)
}

func (s Stack) pop() (Stack, rune) {
	last := s[len(s)-1]
	return s[:len(s)-1], last
}

func (s Stack) peek() rune {
	return s[len(s)-1]
}

func Part2() {
	lines := util.GetLines("sampleinput")
	scores := []float64{}

	for _, line := range lines {
		score, incomplete := processLine(line)
		if !incomplete {
			continue
		}
		scores = append(scores, float64(score))
	}

	res, _ := stats.Median(scores)
	println(int(res))
}

func processLine(line string) (score int, incomplete bool) {
	stack := Stack{}
	for _, ch := range line {
		switch ch {
		case '(', '[', '{', '<':
			stack = stack.push(ch)
		case ')', ']', '}', '>':
			if stack.peek() != match[ch] {
				return 0, false
			} else {
				stack, _ = stack.pop()
			}
		}
	}

	sum := 0
	for i := len(stack) - 1; i >= 0; i-- {
		var last rune
		stack, last = stack.pop()
		sum = sum*5 + part2Points[last]
	}

	return sum, true
}
