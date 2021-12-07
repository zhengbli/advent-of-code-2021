package day6

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	lines := util.GetLines("day6/input")

	bucket := parseLine(lines[0])
	for i := 0; i < 256; i++ {
		bucket = advanceOne(bucket)
	}

	fmt.Println(sum(bucket))
}

func sum(nums [9]int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func advanceOne(buckets [9]int) [9]int {
	newBuckets := [9]int{}

	for i := 0; i < 9; i++ {
		if i == 0 {
			newBuckets[8] += buckets[i]
			newBuckets[6] += buckets[i]
		} else {
			newBuckets[i-1] += buckets[i]
		}
	}

	return newBuckets
}

func parseLine(line string) [9]int {
	res := [9]int{}
	for _, str := range strings.Split(line, ",") {
		strVal, _ := strconv.Atoi(str)
		res[strVal] += 1
	}

	return res
}
