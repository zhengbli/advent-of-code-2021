package util

import (
	"os"
	"strings"
)

func GetLines(file string) []string {
	content, _ := os.ReadFile(file)
	return strings.Split(string(content), "\n")
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Ternary(cond bool, pos int, neg int) int {
	if cond {
		return pos
	}
	return neg
}

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
