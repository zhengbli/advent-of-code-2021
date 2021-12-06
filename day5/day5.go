package day5

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type Line struct {
	Start, End *Point
}

func Part1() {
	lines := util.GetLines("day5/input")
	markers := make(map[Point]int)
	pointCount := 0
	for _, line := range lines {
		ps, ok := parseLine(line)
		if !ok {
			continue
		}

		for _, p := range ps {
			curCount, ok := markers[*p]
			if ok {
				markers[*p] += 1
				if curCount == 1 {
					pointCount += 1
				}
			} else {
				markers[*p] = 1
			}
		}
	}

	fmt.Println(pointCount)
}

func parseLine(line string) (ps []*Point, ok bool) {
	var p1 *Point
	var p2 *Point
	for _, part := range strings.Split(line, "->") {
		trimmed := strings.TrimSpace(part)
		nums := strings.Split(trimmed, ",")
		num0, _ := strconv.Atoi(nums[0])
		num1, _ := strconv.Atoi(nums[1])
		p := &Point{X: num0, Y: num1}
		if p1 == nil {
			p1 = p
		} else {
			p2 = p
		}
	}

	return getMidPoints(p1, p2)
}

func getMidPoints(p1, p2 *Point) (ps []*Point, ok bool) {
	if util.Abs(p1.X-p2.X) != util.Abs(p1.Y-p2.Y) && p1.X != p2.X && p1.Y != p2.Y {
		return nil, false
	}

	deltaX := util.Ternary(p1.X < p2.X, 1, util.Ternary(p1.X == p2.X, 0, -1))
	deltaY := util.Ternary(p1.Y < p2.Y, 1, util.Ternary(p1.Y == p2.Y, 0, -1))

	condFunc := func(cur int, delta int, target int) bool {
		if delta == 0 {
			return cur == target
		}

		if delta < 0 {
			return cur >= target
		}
		return cur <= target
	}

	res := []*Point{}
	for x, y := p1.X, p1.Y; condFunc(x, deltaX, p2.X) && condFunc(y, deltaY, p2.Y); x, y = x+deltaX, y+deltaY {
		res = append(res, &Point{X: x, Y: y})
	}

	return res, true
}
