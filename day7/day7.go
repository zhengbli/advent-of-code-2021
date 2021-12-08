package day7

import (
	"aoc2021/util"
	"fmt"
	"math"
)

func calcFuel(target int, crabs []int) int {
	sum := 0

	for _, crab := range crabs {
		k := util.Abs(crab - target)
		delta := (1 + k) * k / 2
		sum += delta
	}

	return sum
}

func Part1() {
	lines := util.GetLines("day7/input")
	crabs := util.ParseNumLine(lines[0])

	res := biSearch(crabs)
	fmt.Println(res)
}

func biSearch(crabs []int) int {
	min := math.MaxInt
	max := 0
	for _, c := range crabs {
		if c < min {
			min = c
		}

		if c > max {
			max = c
		}
	}

	left, right := min, max
	for left <= right {
		mid := left + (right-left)/2
		midFuel := calcFuel(mid, crabs)

		if left == right {
			return midFuel
		}

		midLeftFuel := calcFuel(mid-1, crabs)
		midRightFuel := calcFuel(mid+1, crabs)

		if midFuel <= midLeftFuel && midFuel <= midRightFuel {
			return midFuel
		}

		if midFuel >= midRightFuel {
			left = mid + 1
			continue
		}

		if midFuel >= midLeftFuel {
			right = mid - 1
			continue
		}
	}

	return 0
}
